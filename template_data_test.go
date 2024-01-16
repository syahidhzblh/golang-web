package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Syahid",
		"Address": map[string]interface{}{
			"Street": "Jalan Kenangan",
		},
	})
}

func TestTemplateData(t *testing.T) {
	reequest := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, reequest)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Name    string
	Title   string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Name:  "Syahid",
		Title: "Template Data Struct",
		Address: Address{
			Street: "Jalan Kenangan",
		},
	})
}

func TestTemplateDataStrurct(t *testing.T) {
	reequest := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, reequest)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
