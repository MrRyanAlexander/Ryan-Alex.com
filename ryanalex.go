package ryanalex

import (
    "html/template"
    "net/http"
)
type Data struct {
	Title string
}
func init() {
    http.HandleFunc("/", handler)
}

func renderTemplate(w http.ResponseWriter, tmpl string, d *Data) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, d)
}

func handler(w http.ResponseWriter, r *http.Request) {
	d := &Data{Title: "Ryan-Alex.com"}
    renderTemplate(w, "index", d)
}