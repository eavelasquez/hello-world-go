package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

// Page represents a wiki page
type Page struct {
	Title string
	Body  []byte
}

// template.Must is a helper function for loading templates
var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))

// validPath is a regexp that matches paths that are valid for the wiki
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// save saves the Page's Body to a text file
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile("data/"+filename, p.Body, 0600)
}

// loadPage reads a text file and stores the body in the Page struct
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile("data/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// renderTemplate is a helper function for rendering templates
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// viewHandler handles requests to /view/{title}
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// saveHandler handles requests to /save/{title} for saving a page
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// editHandler handles requests to /edit/{title} for editing a page
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// makeHandler is a helper function for creating handlers
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2]) // The title is the second subexpression
	}
}

// main is the entry point for the application
func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
