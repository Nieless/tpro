package main

import (
	"github.com/Nieless/tpro/tpro"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/ecf", GetExcelColumnFinder).Methods("GET")

	err := tpro.ListenAndServe(tpro.MustGetEnv("EXCEL_COLUMN_FINDER_SERVICE_HTTP_PORT"), myRouter)
	if err != nil {
		panic(err)
	}
}

// GetExcelColumnFinder serves the html file
func GetExcelColumnFinder(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/excelColumnFinder.html")
}
