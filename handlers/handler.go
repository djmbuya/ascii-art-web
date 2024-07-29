package handlers

import (
	"html/template"
	"log"
	"net/http"

	a "ascii/ascii_art"
)

type ExecOutput struct {
	In  string
	Out string
}

func ValidAscii(s string) bool {
	for _, i := range []byte(s) {
		if i > 127 {
			return false
		}
	}
	return true
}

// Handler handles the HTTP requests.
// Handler handles the HTTP requests.
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			t, err := template.ParseFiles("error/404.html")
			if err != nil {
				internalServerError(w)
				return
			}
			t.Execute(w, nil)
		}
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("index.html")
			if err != nil {
				internalServerError(w)
				return
			}
			t.Execute(w, nil)
		case "POST":
			if err := r.ParseForm(); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Printf("Error parsing form: %v", err)
				return
			}
			input := r.Form.Get("input")
			font := r.Form.Get("font")
			if !ValidAscii(input) {
				w.WriteHeader(http.StatusBadRequest)
				t, err := template.ParseFiles("error/400.html")
				if err != nil {
					internalServerError(w)
					return
				}
				t.Execute(w, nil)
			} else {
				file, status := a.FindFile(input, font)
				if status == 500 {
					w.WriteHeader(http.StatusInternalServerError)
					t, err := template.ParseFiles("error/500.html")
					if err != nil {
						internalServerError(w)
						return
					}
					t.Execute(w, nil)
					internalServerError(w)
					return
				}
				contents, err := a.GetFile(file)
				if err != nil {
					internalServerError(w)
					return
				}
				output := a.ProcessInput(contents, input)
				log.Printf("method: %v / font: %v / input: %v / statuscode: %v\n", r.Method, font, input, status)
				ex := ExecOutput{
					In:  input,
					Out: output,
				}
				t, err := template.ParseFiles("index.html")
				if err != nil {
					internalServerError(w)
					return
				}
				t.Execute(w, ex)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		t, err := template.ParseFiles("error/404.html")
		if err != nil {
			internalServerError(w)
			return
		}
		t.Execute(w, nil)
	}
}

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	t, err := template.ParseFiles("error/500.html")
	if err != nil {
		log.Printf("Error parsing 500.html template: %v", err)
		return
	}
	if err := t.Execute(w, nil); err != nil {
		log.Printf("Error executing 500.html template: %v", err)
	}
}
