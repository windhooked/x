package e13_vuetify_list

// @snippet_begin(VuetifyListSample)
import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func HelloVuetifyList(ctx *web.EventContext) (pr web.PageResponse, err error) {
	wrapper := func(children ...h.HTMLComponent) h.HTMLComponent {
		return VContainer(
			VLayout(
				VFlex(
					VCard(children...),
				).Col(Xs, 6).Offset(Sm, 3),
			).Row(true),
		).GridList(Md).TextAlign(Xs, Center)
	}

	pr.Body = wrapper(
		VToolbar(
			//VToolbarSideIcon(),
			VToolbarTitle("Inbox"),
			VSpacer(),
			VBtn("").Icon(true).Children(
				VIcon("search"),
			),
		).Color("cyan").Dark(true),
		VList(
			VSubheader(h.Text("Today")),
			VListItem(
				VListItemAvatar(
					h.Img("https://cdn.vuetifyxjs.com/images/lists/1.jpg"),
				),
				VListItemContent(
					VListItemTitle(h.Text("Brunch this weekend?")),
					VListItemSubtitle(
						h.Span("Ali Connors").Class("text--primary"),
						h.Text("&mdash; I'll be in your neighborhood doing errands this weekend. Do you want to hang out?"),
					),
				),
			),
			VDivider().Inset(true),
			VListItem(
				VListItemAvatar(
					h.Img("https://cdn.vuetifyxjs.com/images/lists/2.jpg"),
				),
				VListItemContent(
					VListItemTitle(h.RawHTML(`Summer BBQ <span class="grey--text text--lighten-1">4</span>`)),
					VListItemSubtitle(h.RawHTML(`<span class='text--primary'>to Alex, Scott, Jennifer</span> &mdash; Wish I could come, but I'm out of town this weekend.`)),
				),
			),
			VDivider().Inset(true),
			VListItem(
				VListItemAvatar(
					h.Img("https://cdn.vuetifyxjs.com/images/lists/3.jpg"),
				),
				VListItemContent(
					VListItemTitle(h.Text(`Oui oui`)),
					VListItemSubtitle(h.RawHTML(`<span class='text--primary'>Sandra Adams</span> &mdash; Do you have Paris recommendations? Have you ever been?`)),
				),
			),
		).TwoLine(true),
	)

	return
}

// @snippet_end

const HelloVuetifyListPath = "/samples/hello-vuetify-list"
