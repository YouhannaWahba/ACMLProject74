package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init (){
	tpl = 	template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/maps",maps)
	http.HandleFunc("/process",processor)
	http.ListenAndServe(":3000",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"index.html",nil)
}

func maps(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"maps.html",nil)
}

func processor(w http.ResponseWriter , r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		First string 
		Last string 
	}{
		First: fname,
		Last:lname,
	}
	tpl.ExecuteTemplate(w,"processor.html",d)
}
