package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	templates := populateTemplates()
	//http.ListenAndServe(":3006", http.FileServer(http.Dir("/Users/Admin/Desktop/go-workspace/Web-application-with-golang/public/")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("/Users/Admin/Desktop/go-workspace/Web-application-with-golang/public/")))
	http.Handle("/css/", http.FileServer(http.Dir("/Users/Admin/Desktop/go-workspace/Web-application-with-golang/public/")))
	http.ListenAndServe(":8080", nil)
}

func populateTemplates() map[string]*template.Template {
	/*const templates = "/Users/Admin/Desktop/go-workspace/Web-application-with-golang/templates/"
	result := template.New("template")
	const basePath = templates
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result */

	result := make(map[string]*template.Template)
	const basePath = "templetes"
	layout := template.Must(template.ParseFiles(basePath + "/_layout,html"))

}
