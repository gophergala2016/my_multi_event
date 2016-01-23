package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var Tmpl *template.Template

func init() {
	InitTmpl()
}
func InitTmpl() {
	Tmpl = template.New("pages")
	assets, _ := AssetDir("assets/tmpl")
	for _, asset := range assets {
		src, _ := Asset("assets/tmpl/" + asset)
		fmt.Println(Tmpl.New(asset).Parse(string(src)))
	}
}

func main() {
	http.HandleFunc("/assets/", AssetHandler)

	http.HandleFunc("/schedule", SchedulePage)
	http.HandleFunc("/create", SchedulePage)
	http.HandleFunc("/results", SchedulePage)
	http.HandleFunc("/modify", SchedulePage)
	http.HandleFunc("/my_events", MyEventsPage)

	http.HandleFunc("/", FrontPage)
	http.ListenAndServe(":8000", nil)
	fmt.Println("Beginnings")
}

func FrontPage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
	Tmpl.ExecuteTemplate(w, "index.html", nil)
}
func SchedulePage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
}
func CreatePage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
}
func ResultsPage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
}
func ModifyPage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
}
func MyEventsPage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
}

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.Contains(r.URL.Path, "/js/"):
		w.Header().Add("Content-Type", "application/javascript")
	case strings.Contains(r.URL.Path, "/css/"):
		w.Header().Add("Content-Type", "text/css")
	}
	data, err := Asset(r.URL.Path[1:])
	if err != nil || strings.HasPrefix(r.URL.Path, "/assets/tmpl") {
		fmt.Println(r.URL.Path, err)
		w.WriteHeader(404)
	} else {
		w.Write(data)
	}
}
