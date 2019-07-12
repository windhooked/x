package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VListTileTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListTileTitle(children ...h.HTMLComponent) (r *VListTileTitleBuilder) {
	r = &VListTileTitleBuilder{
		tag: h.Tag("v-list-tile-title").Children(children...),
	}
	return
}

func (b *VListTileTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}