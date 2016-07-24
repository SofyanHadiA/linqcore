package linqcore

import (
	"net/http"
	"text/template"

	"github.com/SofyanHadiA/linqcore/utils"
)

// View holds view objet data
type View struct {
	BaseURL      string
	PageTitle    string
	PageDesc     string
	Template     map[string]string
	Data         map[string]interface{}
	ViewLocation string
}

// NewView create new view data
func NewView(viewLocation string, configs Configs) View {
	return View{
		BaseURL:      configs.GetStrConfig("app.baseUrl"),
		PageTitle:    configs.GetStrConfig("app.pageTitle"),
		ViewLocation: viewLocation,
		Template: map[string]string{
			"mainTemplate":    viewLocation + "template.html",
			"headerTemplate":  viewLocation + "header.html",
			"footerTemplate":  viewLocation + "footer.html",
			"sidebarTemplate": viewLocation + "sidebar.html",
			"menubarTemplate": viewLocation + "menubar.html",
		},
	}
}

// ParseHTML parse html without header and footer
func (view View) ParseHTML(templateLoc string, w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	utils.Log.Debug("Parsing view(s): ", templateLoc)
	t := template.Must(template.ParseFiles(view.ViewLocation + templateLoc))

	err := t.ExecuteTemplate(w, "main", data)
	if err != nil {
		utils.Log.Fatal("executing template: ", err)
	}
}

// ParseHTMLTemplate with html with footer and header
func (view View) ParseHTMLTemplate(templateLoc string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	utils.Log.Debug("Parsing view(s): ", view.Template["mainTemplate"], templateLoc)
	t := template.Must(template.ParseFiles(
		view.Template["mainTemplate"],
		view.Template["headerTemplate"],
		view.Template["footerTemplate"],
		view.Template["sidebarTemplate"],
		view.Template["menubarTemplate"],
		templateLoc,
	))

	err := t.ExecuteTemplate(w, "main", view)
	if err != nil {
		utils.Log.Fatal("executing template: ", err)
	}
}
