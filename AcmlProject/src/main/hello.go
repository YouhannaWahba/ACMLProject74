package main

import (
	"net/http"
	"html/template"
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/maps", maps)
	http.HandleFunc("/process", processor)
	http.HandleFunc("/calculateGPA", calculatGPA)
	http.HandleFunc("/dropCourse", dropCourse)
	http.ListenAndServe(":3000", nil)

	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	  db, err := sql.Open("postgres", psqlInfo)
	  if err != nil {
		panic(err)
	  }
	  defer db.Close()
	
	  err = db.Ping()
	  if err != nil {
		panic(err)
	  }
	
	  fmt.Println("Successfully connected!")
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func maps(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "maps.html", nil)
}

func dropCourse(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dropCourse.html", nil)
}

func calculatGPA(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "calculateGPA.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}
	tpl.ExecuteTemplate(w, "processor.html", d)
}
