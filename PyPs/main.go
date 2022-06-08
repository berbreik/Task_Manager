package main

import (
	"PyPs/driver"
	"PyPs/handler"
	"PyPs/service"
	"PyPs/store"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
// main is the entry point of the program
func main() {

	sqlconfif := driver.Mysqlconfig{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "root",
		Dbname:   "task_manager",
	}

	db,err := driver.ConnectToMySQL(sqlconfif)
	if err != nil {
		log.Fatal(err)
	}

	defer func (db *sql.DB){
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}

	}(db) 

	s := store.NewDataStore(db)
	serv := service.NewService(s)
	hand := handler.NewHandler(serv)

	route := mux.NewRouter()
	route.HandleFunc("/Admin/Projects", hand.PostAdminHandler).Methods("POST")
	route.HandleFunc("/Manager/Task", hand.PostTaskHandler).Methods("POST")
	route.HandleFunc("/Manager/Task", hand.UpdateTaskHandler).Methods("PUT")
	route.HandleFunc("/Manager/Task/{id}", hand.GetTaskByIdHandler).Methods("GET")
	route.HandleFunc("/Manager/Task", hand.GetAllHandler).Methods("GET")
	

	log.Fatal(http.ListenAndServe(":8000", route))
	
}
