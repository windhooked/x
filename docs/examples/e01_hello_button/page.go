package e01_hello_button

import (
	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	Message string
}

func HelloButton(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("reload", reload)

	var s = &mystate{}
	if ctx.Flash != nil {
		s = ctx.Flash.(*mystate)
	}

	pr.Body = Div(
		Button("Hello").Attr("@click", web.Plaid().EventFunc("reload").Go()),
		Tag("input").
			Attr("type", "text").
			Attr("value", s.Message).
			Attr("@input", web.Plaid().
				EventFunc("reload").
				FieldValue("Message", web.Var("$event.target.value")).
				Go()),
		Div().
			Style("font-family: monospace;").
			Text(s.Message),
	)
	return
}

func reload(ctx *web.EventContext) (r web.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)
	ctx.Flash = s

	r.Reload = true
	return
}
