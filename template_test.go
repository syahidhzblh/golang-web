package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	// t, err := template.ParseFiles("./templates/simple.gohtml")
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.ParseFiles("./template/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateHTMLFile(t *testing.T) {
	reequest := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, reequest)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplatesEmbed(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template via embed")
}

func TestTemplateEmbed(t *testing.T) {
	reequest := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplatesEmbed(recorder, reequest)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
