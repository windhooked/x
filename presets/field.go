package presets

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/goplaid/web"
	"github.com/goplaid/x/i18n"
	v "github.com/goplaid/x/vuetify"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

type FieldBuilder struct {
	NameLabel
	compFunc   FieldComponentFunc
	setterFunc FieldSetterFunc
	context    context.Context
}

func NewField(name string) (r *FieldBuilder) {
	r = &FieldBuilder{}
	r.name = name
	r.compFunc = emptyComponentFunc
	return
}

func emptyComponentFunc(obj interface{}, field *FieldContext, ctx *web.EventContext) (r h.HTMLComponent) {
	log.Printf("No ComponentFunc for field %v\n", field.Name)
	return
}

func (b *FieldBuilder) Label(v string) (r *FieldBuilder) {
	b.label = v
	return b
}

func (b *FieldBuilder) Clone() (r *FieldBuilder) {
	r = &FieldBuilder{}
	r.name = b.name
	r.label = b.label
	r.compFunc = b.compFunc
	r.setterFunc = b.setterFunc
	return r
}

func (b *FieldBuilder) ComponentFunc(v FieldComponentFunc) (r *FieldBuilder) {
	if v == nil {
		panic("value required")
	}
	b.compFunc = v
	return b
}

func (b *FieldBuilder) SetterFunc(v FieldSetterFunc) (r *FieldBuilder) {
	b.setterFunc = v
	return b
}

func (b *FieldBuilder) WithContextValue(key interface{}, val interface{}) (r *FieldBuilder) {
	if b.context == nil {
		b.context = context.Background()
	}
	b.context = context.WithValue(b.context, key, val)
	return b
}

type NameLabel struct {
	name  string
	label string
}

type FieldBuilders struct {
	obj         interface{}
	defaults    *FieldDefaults
	fieldLabels []string
	fields      []*FieldBuilder
}

func (b *FieldBuilders) Clone() (r *FieldBuilders) {
	r = &FieldBuilders{
		obj:         b.obj,
		defaults:    b.defaults,
		fieldLabels: b.fieldLabels,
	}
	return
}

func (b *FieldBuilders) Field(name string) (r *FieldBuilder) {
	r = b.GetField(name)
	if r != nil {
		return
	}

	r = NewField(name)
	b.fields = append(b.fields, r)
	return
}

func (b *FieldBuilders) Labels(vs ...string) (r *FieldBuilders) {
	b.fieldLabels = append(b.fieldLabels, vs...)
	return b
}

func (b *FieldBuilders) getLabel(field NameLabel) (r string) {
	if len(field.label) > 0 {
		return field.label
	}

	for i := 0; i < len(b.fieldLabels)-1; i = i + 2 {
		if b.fieldLabels[i] == field.name {
			return b.fieldLabels[i+1]
		}
	}

	return field.name
}

func (b *FieldBuilders) GetField(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	return
}

func (b *FieldBuilders) Only(names ...string) (r *FieldBuilders) {
	if len(names) == 0 {
		return b
	}

	r = b.Clone()

	for _, n := range names {
		f := b.GetField(n)
		if f == nil {
			fType := reflectutils.GetType(b.obj, n)
			if fType == nil {
				fType = reflect.TypeOf("")
			}

			ft := b.defaults.fieldTypeByTypeOrCreate(fType)
			r.Field(n).
				ComponentFunc(ft.compFunc).
				SetterFunc(ft.setterFunc)
		} else {
			r.fields = append(r.fields, f.Clone())
		}
	}

	return
}

func (b *FieldBuilders) Except(patterns ...string) (r *FieldBuilders) {
	if len(patterns) == 0 {
		return
	}

	r = &FieldBuilders{fieldLabels: b.fieldLabels}
	for _, f := range b.fields {
		if hasMatched(patterns, f.name) {
			continue
		}
		r.fields = append(r.fields, f.Clone())
	}
	return
}

func (b *FieldBuilders) String() (r string) {
	var names []string
	for _, f := range b.fields {
		names = append(names, f.name)
	}
	return fmt.Sprint(names)
}

func (b *FieldBuilders) ToComponent(mb *ModelBuilder, obj interface{}, verr *web.ValidationErrors, ctx *web.EventContext) h.HTMLComponent {

	var comps []h.HTMLComponent

	if verr == nil {
		verr = &web.ValidationErrors{}
	}

	gErr := verr.GetGlobalError()
	if len(gErr) > 0 {
		comps = append(
			comps,
			v.VAlert(h.Text(gErr)).
				Border("left").
				Type("error").
				Elevation(2).
				ColoredBorder(true),
		)
	}

	for _, f := range b.fields {
		if f.compFunc == nil {
			continue
		}

		if mb.Info().Verifier().Do(PermUpdate).ObjectOn(obj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil {
			continue
		}

		comps = append(comps, f.compFunc(obj, &FieldContext{
			ModelInfo: mb.Info(),
			Name:      f.name,
			Label:     i18n.PT(ctx.R, ModelsI18nModuleKey, mb.label, b.getLabel(f.NameLabel)),
			Errors:    verr.GetFieldErrors(f.name),
			Context:   f.context,
		}, ctx))
	}

	return h.Components(comps...)
}
