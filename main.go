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
	id    int
	fio   string
	place string
	//document string
	//sex string
}

var database *sql.DB

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		fio := r.FormValue("fio")
		place := r.FormValue("place")
		//sex := r.FormValue("sex")
		//specialization := r.FormValue("specialization")
		//red :- r.FormValue("red")
		document := r.FormValue("document")
		database.Exec("INSERT INTO students (fio, place, document) values (?, ?, ?)",
			fio, place, document)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "templates/student.html")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	db, err := sql.Open("mysql", "ssdyubrf_project:*TbB5SF8@tcp(ssdyubrf.beget.tech:3306)/ssdyubrf_project")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", handler)
	http.HandleFunc("/create", CreateHandler)
	http.ListenAndServe(":8080", nil)
}
