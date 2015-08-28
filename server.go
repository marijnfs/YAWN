package main

import (
	"fmt"
	"flag"
	"net/http"
	"os"
	"io/ioutil"
	"text/template"
	"html"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/gorilla/schema"
)

var Wd string
var body_template, edit_template *template.Template

type Page struct {
	Sub string
	Body string
}

func init() {
	var err error
	
	flag.Parse()
	Wd, _ = os.Getwd()

	body_template, err = template.ParseFiles("tpl/body.tpl")
	check(err)
	edit_template, err = template.ParseFiles("tpl/edit.tpl")
	check(err)
}

func GetBody(name []byte) []byte {
	return MDToHTML(GetBodyRaw(name))
}

func GetBodyRaw(name []byte) []byte {
	return []byte(html.EscapeString(string(Get(name))))
}

func getSubFromRequest(r *http.Request) string {
	vars := mux.Vars(r)
	sub := vars["sub"]
	if sub == "" || sub == "index.html" { sub = "root" }
	return sub
}

func EditGetHandler(w http.ResponseWriter, r *http.Request) {
	sub := getSubFromRequest(r)
	body := GetBodyRaw([]byte(sub))
	page := Page{sub, string(body)}
	
	edit_template.Execute(w, page)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	sub := getSubFromRequest(r)
	
	err := r.ParseForm()
	if err != nil {
		check(err)
	}

	//create page
	var page Page
	decoder := schema.NewDecoder()
	decoder.Decode(&page, r.PostForm)
	page.Sub = sub

	Put([]byte(page.Sub), []byte(page.Body))
	
	edit_template.Execute(w, page)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	sub := getSubFromRequest(r)
	fmt.Println("page: ", sub)
	body := GetBody([]byte(sub))
	page := Page{sub, string(body)}

	body_template.Execute(w, page)
}

func RootRedirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("root")
	http.Redirect(w, r, "/", http.StatusOK)
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	header := []byte("HTTP/1.1 200 OK\r\nContent-Type: image/x-icon\r\n\r\n")

	data, err := ioutil.ReadFile("favicon.ico")
	check(err)
	data = append(header, data...)
	w.Write(data)
}


func main() {
	fmt.Println("Server Started in ", Wd)
	r := mux.NewRouter()


	r.HandleFunc("/favicon.ico", FaviconHandler)

	//root handlers
	r.HandleFunc("/edit/", EditGetHandler).Methods("GET")
	r.HandleFunc("/edit/", EditPostHandler).Methods("POST")
	r.HandleFunc("/", PageHandler)

	//regular handlers
	r.HandleFunc("/edit/{sub}", EditGetHandler).Methods("GET")
	r.HandleFunc("/edit/{sub}", EditPostHandler).Methods("POST")
	r.HandleFunc("/{sub}", PageHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(Wd + "static/"))))
	
	
	http.ListenAndServe("127.0.0.1:1234", handlers.CompressHandler(r))
}
