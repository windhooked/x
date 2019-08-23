package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VContainerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VContainer(children ...h.HTMLComponent) (r *VContainerBuilder) {
	r = &VContainerBuilder{
		tag: h.Tag("v-container").Children(children...),
	}
	return
}

func (b *VContainerBuilder) AlignBaseline(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":align-baseline", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) AlignCenter(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":align-center", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) AlignContentCenter(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":align-content-center", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) AlignContentEnd(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":align-content-end", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) AlignContentSpaceAround(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":align-content-space-around", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) AlignContentSpaceBetween(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":align-content-space-between", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) AlignContentStart(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":align-content-start", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) AlignEnd(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":align-end", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) AlignStart(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":align-start", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) Dtype(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":d-{type}", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) FillHeight(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":fill-height", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) Fluid(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":fluid", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) GridListxsThroughXl(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":grid-list-{xs-through-xl}", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) Id(v string) (r *VContainerBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VContainerBuilder) JustifyCenter(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":justify-center", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) JustifyEnd(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":justify-end", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) JustifySpaceAround(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":justify-space-around", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) JustifySpaceBetween(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":justify-space-between", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) JustifyStart(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":justify-start", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) Tag(v string) (r *VContainerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VContainerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VContainerBuilder) Attr(vs ...interface{}) (r *VContainerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VContainerBuilder) Children(children ...h.HTMLComponent) (r *VContainerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VContainerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VContainerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VContainerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VContainerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VContainerBuilder) Class(names ...string) (r *VContainerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VContainerBuilder) ClassIf(name string, add bool) (r *VContainerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VContainerBuilder) On(name string, value string) (r *VContainerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VContainerBuilder) Bind(name string, value string) (r *VContainerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VContainerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
