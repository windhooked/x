package e00_basics

//@snippet_begin(ReloadWithFlashSample)
import (
	"fmt"
	"time"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

var count int

func ReloadWithFlash(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("update2", update2)

	var msg HTMLComponent

	if d, ok := ctx.Flash.(*Data1); ok {
		msg = Div().Text(d.Msg).Style("border: 5px solid orange;")
	} else {
		count = 0
	}

	pr.Body = Div(
		H1("Whole Page Reload With a Flash"),
		msg,
		Div().Text(time.Now().Format(time.RFC3339Nano)),
		Button("Do Something").
			Attr("@click", web.Plaid().EventFunc("update2").Go()),
	)
	return
}

type Data1 struct {
	Msg string
}

func update2(ctx *web.EventContext) (er web.EventResponse, err error) {
	count++
	ctx.Flash = &Data1{Msg: fmt.Sprintf("The page is reloaded: %d", count)}
	er.Reload = true
	return
}

//@snippet_end

const ReloadWithFlashPath = "/samples/reload_with_flash"
