package main

import (
	//"fmt"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	//"html/template"
	"log"
)

type Product struct {
	id             int
	fio            string
	date           int
	school         string
	dateschoolend  int
	addres         string
	specialization string
	red            string
}

var database *sql.DB

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		fio := r.FormValue("fio")
		date := r.FormValue("date")
		school := r.FormValue("school")
		dateschoolend := r.FormValue("dateschoolend")
		addres := r.FormValue("addres")
		specialization := r.FormValue("specialization")
		red := r.FormValue("red")
		_, err = database.Exec("INSERT INTO students (fio, date, school, dateschoolend, addres, specialization, red) values (?, ?, ?, ?, ?, ?, ?)",
			fio, date, school, dateschoolend, addres, specialization, red)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "templates/student.html")
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	db, err := sql.Open("mysql", "ssdyubrf_project:*TbB5SF8@tcp(ssdyubrf.beget.tech:3306)/ssdyubrf_project")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", Index)
	http.HandleFunc("/create", CreateHandler)
	//http.HandleFunc("/create", CreateHandler)
	http.ListenAndServe(":8080", nil)
}
