package belajargolangserver

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var myTemplates = template.Must(template.ParseFiles("templates/post.gohtml"))

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Apa itu XSS?",
		"Body":  "<p>XSS (Cross-Site Scripting) adalah serangan keamanan web dimana penyerang menyisipkan kode berbahaya (biasanya JavaScript) ke dalam website yang dapat dieksekusi di browser pengguna lain. XSS dapat digunakan untuk mencuri data sensitif, mengubah tampilan website, atau mengambil alih sesi pengguna.</p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Disable Auto Escape adalah fitur yang
// memungkinkan kita untuk menonaktifkan
// fitur auto escape pada template.
// dan ketika kita ingin menampilkan HTML biasa
// tetapi featur ini akan rentan terhadap serangan XSS
// jika featur auto escape dinonaktifkan.

func TemplateAutoEscapeDisable(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Apa itu XSS?",
		"Body":  template.HTML("<p>XSS (Cross-Site Scripting) adalah serangan keamanan web dimana penyerang menyisipkan kode berbahaya (biasanya JavaScript) ke dalam website yang dapat dieksekusi di browser pengguna lain. XSS dapat digunakan untuk mencuri data sensitif, mengubah tampilan website, atau mengambil alih sesi pengguna.</p>"),
	})
}

func TestTemplateAutoEscapeDisable(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisable(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisableServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisable),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// uji coba xss
func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Apa itu XSS?",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateAutoEscapeXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?body=<script>alert('XSS')</script>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
