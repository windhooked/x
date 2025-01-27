package presets

import (
	"context"
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets/actions"
	. "github.com/goplaid/x/vuetify"
	"github.com/iancoleman/strcase"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

type FieldContext struct {
	Name      string
	Label     string
	Errors    []string
	ModelInfo *ModelInfo
	Context   context.Context
}

func (fc *FieldContext) StringValue(obj interface{}) (r string) {
	val := fc.Value(obj)
	switch vt := val.(type) {
	case []rune:
		return string(vt)
	case []byte:
		return string(vt)
	}
	return fmt.Sprint(val)
}

func (fc *FieldContext) Value(obj interface{}) (r interface{}) {
	fieldName := fc.Name
	return reflectutils.MustGet(obj, fieldName)
}

func (fc *FieldContext) ContextValue(key interface{}) (r interface{}) {
	if fc.Context == nil {
		return
	}
	return fc.Context.Value(key)
}

type FieldDefaultBuilder struct {
	valType    reflect.Type
	compFunc   FieldComponentFunc
	setterFunc FieldSetterFunc
}

type FieldMode int

const (
	WRITE FieldMode = iota
	LIST
	DETAIL
)

func NewFieldDefault(t reflect.Type) (r *FieldDefaultBuilder) {
	r = &FieldDefaultBuilder{valType: t}
	return
}

func (b *FieldDefaultBuilder) ComponentFunc(v FieldComponentFunc) (r *FieldDefaultBuilder) {
	b.compFunc = v
	return b
}

func (b *FieldDefaultBuilder) SetterFunc(v FieldSetterFunc) (r *FieldDefaultBuilder) {
	b.setterFunc = v
	return b
}

var numberVals = []interface{}{
	int(0), int8(0), int16(0), int32(0), int64(0),
	uint(0), uint(8), uint16(0), uint32(0), uint64(0),
	float32(0.0), float64(0.0),
}

var stringVals = []interface{}{
	string(""),
	[]rune(""),
	[]byte(""),
}

type FieldDefaults struct {
	mode             FieldMode
	fieldTypes       []*FieldDefaultBuilder
	excludesPatterns []string
}

func NewFieldDefaults(t FieldMode) (r *FieldDefaults) {
	r = &FieldDefaults{
		mode: t,
	}
	r.builtInFieldTypes()
	return
}

func (b *FieldDefaults) FieldType(v interface{}) (r *FieldDefaultBuilder) {
	return b.fieldTypeByTypeOrCreate(reflect.TypeOf(v))
}

func (b *FieldDefaults) Exclude(patterns ...string) (r *FieldDefaults) {
	b.excludesPatterns = patterns
	return b
}

func (b *FieldDefaults) InspectFields(val interface{}) (r *FieldBuilders) {
	r, _ = b.inspectFieldsAndCollectName(val, nil)
	return
}

func (b *FieldDefaults) inspectFieldsAndCollectName(val interface{}, collectType reflect.Type) (r *FieldBuilders, names []string) {
	v := reflect.ValueOf(val)

	for v.Elem().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v = v.Elem()

	t := v.Type()

	r = &FieldBuilders{
		defaults: b,
		obj:      val,
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		ft := b.fieldTypeByType(f.Type)

		if !hasMatched(b.excludesPatterns, f.Name) && ft != nil {
			r.Field(f.Name).
				ComponentFunc(ft.compFunc).
				SetterFunc(ft.setterFunc)
		}

		if collectType != nil && f.Type == collectType {
			names = append(names, strcase.ToSnake(f.Name))
		}
	}

	return
}

func hasMatched(patterns []string, name string) bool {
	for _, p := range patterns {
		ok, err := filepath.Match(p, name)
		if err != nil {
			panic(err)
		}
		if ok {
			return true
		}
	}
	return false
}

func (b *FieldDefaults) fieldTypeByType(tv reflect.Type) (r *FieldDefaultBuilder) {
	for _, ft := range b.fieldTypes {
		if ft.valType == tv {
			return ft
		}
	}
	return nil
}

func (b *FieldDefaults) fieldTypeByTypeOrCreate(tv reflect.Type) (r *FieldDefaultBuilder) {
	if r = b.fieldTypeByType(tv); r != nil {
		return
	}

	r = NewFieldDefault(tv)

	if b.mode == LIST {
		r.ComponentFunc(cfTextTd)
	} else {
		r.ComponentFunc(cfTextField)
	}
	b.fieldTypes = append(b.fieldTypes, r)
	return
}

func cfTextTd(obj interface{}, field *FieldContext, ctx *web.EventContext) h.HTMLComponent {
	if field.Name == "ID" {
		id := field.StringValue(obj)
		if len(id) > 0 {
			mi := field.ModelInfo
			if mi == nil {
				return h.Td().Text(id)
			}

			var a h.HTMLComponent
			if mi.HasDetailing() {
				a = h.A().Text(id).Attr("@click", web.Plaid().
					PushStateURL(mi.DetailingHref(id)).
					Go(),
				)
			} else {
				if field.ModelInfo.Verifier().Do(PermUpdate).ObjectOn(obj).WithReq(ctx.R).IsAllowed() == nil {
					a = h.A().Text(id).Attr("@click", web.Plaid().EventFunc(actions.DrawerEdit, id).
						Go())
				} else {
					a = h.Text(id)
				}
			}
			return h.Td(a)
		}
	}
	return h.Td(h.Text(field.StringValue(obj)))
}

func cfCheckbox(obj interface{}, field *FieldContext, ctx *web.EventContext) h.HTMLComponent {
	return VCheckbox().
		FieldName(field.Name).
		Label(field.Label).
		InputValue(reflectutils.MustGet(obj, field.Name).(bool)).
		ErrorMessages(field.Errors...)
}

func cfNumber(obj interface{}, field *FieldContext, ctx *web.EventContext) h.HTMLComponent {
	return VTextField().
		Type("number").
		FieldName(field.Name).
		Label(field.Label).
		Value(fmt.Sprint(reflectutils.MustGet(obj, field.Name))).
		ErrorMessages(field.Errors...)
}

func cfTextField(obj interface{}, field *FieldContext, ctx *web.EventContext) h.HTMLComponent {
	return VTextField().
		Type("text").
		FieldName(field.Name).
		Label(field.Label).
		Value(reflectutils.MustGet(obj, field.Name).(string)).
		ErrorMessages(field.Errors...)
}

func (b *FieldDefaults) builtInFieldTypes() {

	if b.mode == LIST {
		b.FieldType(true).
			ComponentFunc(cfTextTd)

		for _, v := range numberVals {
			b.FieldType(v).
				ComponentFunc(cfTextTd)
		}

		for _, v := range stringVals {
			b.FieldType(v).
				ComponentFunc(cfTextTd)
		}
		return
	}

	b.FieldType(true).
		ComponentFunc(cfCheckbox)

	for _, v := range numberVals {
		b.FieldType(v).
			ComponentFunc(cfNumber)
	}

	for _, v := range stringVals {
		b.FieldType(v).
			ComponentFunc(cfTextField)
	}

	b.Exclude("ID")
	return
}
