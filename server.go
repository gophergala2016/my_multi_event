package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var Tmpl map[string]*template.Template

func init() {
	InitTmpl()
}
func InitTmpl() {
	if Tmpl != nil && os.Getenv("DEV") == "" {
		return
	}
	Tmpl = map[string]*template.Template{}
	base := template.New("pages")
	assets, _ := AssetDir("assets/tmpl/base")
	for _, asset := range assets {
		src, _ := Asset("assets/tmpl/base/" + asset)
		fmt.Println(base.New(asset).Parse(string(src)))
	}
	pages, _ := AssetDir("assets/tmpl/pages")
	for _, page := range pages {
		src, _ := Asset("assets/tmpl/pages/" + page)
		tmpl, _ := base.Clone()
		fmt.Println(tmpl.New(page).Parse(string(src)))
		Tmpl[page] = tmpl
	}
}

func main() {
	http.HandleFunc("/assets/", AssetHandler)

	http.HandleFunc("/schedule", SchedulePage)
	http.HandleFunc("/create", CreatePage)
	http.HandleFunc("/results", ResultsPage)
	http.HandleFunc("/modify", ModifyPage)
	http.HandleFunc("/my_events", MyEventsPage)

	http.HandleFunc("/event", GetEvent)
	http.HandleFunc("/responses", GetResponses)
	http.HandleFunc("/create_event", SaveEvent)
	http.HandleFunc("/create_response", SaveResponse)

	http.HandleFunc("/", FrontPage)
	http.ListenAndServe(":8000", nil)
	fmt.Println("Beginnings")
}

func FrontPage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
	Tmpl["landing.html"].ExecuteTemplate(w, "landing.html", nil)
}
func SchedulePage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
	id := r.URL.Query().Get("event_id")
	event, ok := Store.GetEvent(id)
	if ok {
		fmt.Println(Tmpl["schedule.html"].ExecuteTemplate(w, "schedule.html", map[string]interface{}{"Event": event}))
	} else {
		Tmpl["schedule.html"].ExecuteTemplate(w, "schedule.html", nil)
	}
}
func CreatePage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
	m := map[string]string{
		"NGApp":  "CreateApp",
		"NGCtrl": "CreateEventCtrl",
	}
	fmt.Println(Tmpl["create.html"].ExecuteTemplate(w, "create.html", m))
}
func ResultsPage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
	id := r.URL.Query().Get("event_id")
	event, ok := Store.GetEvent(id)
	if ok {
		Tmpl["schedule.html"].ExecuteTemplate(w, "schedule.html", map[string]interface{}{"Event": event})
	} else {
		Tmpl["schedule.html"].ExecuteTemplate(w, "schedule.html", nil)
	}

	Tmpl["my_results.html"].ExecuteTemplate(w, "my_results.html", nil)
}
func ModifyPage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
	Tmpl["modify.html"].ExecuteTemplate(w, "modify.html", nil)
}
func MyEventsPage(w http.ResponseWriter, r *http.Request) {
	InitTmpl()
	Tmpl["my_events.html"].ExecuteTemplate(w, "my_events.html", nil)
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

func SaveEvent(w http.ResponseWriter, r *http.Request) {
	evt := Event{}
	json.NewDecoder(r.Body).Decode(&evt)
	key := Store.NewEvent(evt)
	json.NewEncoder(w).Encode(map[string]string{
		"response": "ok",
		"key":      key,
	})
}
func SaveResponse(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Created: time.Now(),
		Answers: []string{},
	}
	b := bytes.Buffer{}
	io.Copy(&b, r.Body)
	fmt.Println(b.String())
	json.NewDecoder(&b).Decode(&resp)
	Store.NewResponse(r.URL.Query().Get("event_id"), resp)
	json.NewEncoder(w).Encode(map[string]string{
		"response": "ok",
	})
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	event, ok := Store.GetEvent(r.URL.Query().Get("event_id"))
	enc := json.NewEncoder(w)
	if ok {
		enc.Encode(event)
	} else {
		w.WriteHeader(404)
	}
}
func GetResponses(w http.ResponseWriter, r *http.Request) {
	resp := Store.GetResponses(r.URL.Query().Get("event_id"))
	json.NewEncoder(w).Encode(resp)
}
