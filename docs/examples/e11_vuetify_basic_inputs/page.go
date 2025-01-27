package e11_vuetify_basic_inputs

// @snippet_begin(VuetifyBasicInputsSample)
import (
	"mime/multipart"

	"github.com/goplaid/web"
	"github.com/goplaid/x/docs/utils"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

type myFormValue struct {
	MyValue          string
	TextareaValue    string
	Gender           string
	Agreed           bool
	Feature1         bool
	Slider1          int
	PortalAddedValue string
	Files1           []*multipart.FileHeader
	Files2           []*multipart.FileHeader
	Files3           []*multipart.FileHeader
}

var s = &myFormValue{
	MyValue:       "123",
	TextareaValue: "This is textarea value",
	Gender:        "M",
	Agreed:        false,
	Feature1:      true,
	Slider1:       60,
}

func VuetifyBasicInputs(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("update", update)
	ctx.Hub.RegisterEventFunc("addPortal", addPortal)

	var verr web.ValidationErrors
	if ve, ok := ctx.Flash.(web.ValidationErrors); ok {
		verr = ve
	}

	pr.Body = VContainer(
		utils.PrettyFormAsJSON(ctx),
		VTextField().
			Label("Form ValueIs").
			Solo(true).
			Clearable(true).
			FieldName("MyValue").
			ErrorMessages(verr.GetFieldErrors("MyValue")...).
			Value(s.MyValue),
		VTextarea().FieldName("TextareaValue").
			ErrorMessages(verr.GetFieldErrors("TextareaValue")...).
			Solo(true).Value(s.TextareaValue),
		VRadioGroup(
			VRadio().Value("F").Label("Female"),
			VRadio().Value("M").Label("Male"),
		).FieldName("Gender").Value(s.Gender),
		VCheckbox().FieldName("Agreed").
			ErrorMessages(verr.GetFieldErrors("Agreed")...).
			Label("Agree").InputValue(s.Agreed),
		VSwitch().FieldName("Feature1").InputValue(s.Feature1),

		VSlider().FieldName("Slider1").
			ErrorMessages(verr.GetFieldErrors("Slider1")...).
			Value(s.Slider1),
		web.Portal().Name("Portal1"),

		VFileInput().FieldName("Files1"),

		VFileInput().Label("Auto post to server after select file").Multiple(true).
			Attr("@change", web.Plaid().
				EventFunc("update").
				FieldValue("Files2", web.Var("$event")).
				Go()),

		h.Div(
			h.Input("Files3").Type("file").
				Attr("@input", web.Plaid().
					EventFunc("update").
					FieldValue("Files3", web.Var("$event")).
					Go()),
		).Class("mb-4"),

		VBtn("Update").OnClick("update").Color("primary"),
		h.P().Text("The following button will update a portal with a hidden field, if you click this button, and then click the above update button, you will find additional value posted to server"),
		VBtn("Add Portal Hidden Value").OnClick("addPortal"),
	)

	return
}

func addPortal(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
		Name: "Portal1",
		Body: h.Input("").Type("hidden").Value("this is my portal added hidden value").Attr(web.VFieldName("PortalAddedValue")...),
	})
	return
}

func update(ctx *web.EventContext) (r web.EventResponse, err error) {
	s = &myFormValue{}
	ctx.MustUnmarshalForm(s)
	verr := web.ValidationErrors{}
	if len(s.MyValue) < 10 {
		verr.FieldError("MyValue", "my value is too small")
	}

	if len(s.TextareaValue) > 5 {
		verr.FieldError("TextareaValue", "textarea value is too large")
	}

	if !s.Agreed {
		verr.FieldError("Agreed", "You must agree the terms")
	}

	if s.Slider1 > 50 {
		verr.FieldError("Slider1", "You slide too much")
	}

	ctx.Flash = verr
	r.Reload = true

	return
}

// @snippet_end

const VuetifyBasicInputsPath = "/samples/vuetify-basic-inputs"
