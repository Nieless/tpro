package main

import (
	"encoding/json"
	"github.com/Nieless/tpro/tpro"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/datetime", GetDateTime).Methods("GET")
	myRouter.HandleFunc("/datetime", PostDateTime).Methods("POST")

	err := tpro.ListenAndServe(tpro.MustGetEnv("DATE_TIME_SERVICE_HTTP_PORT"), myRouter)
	if err != nil {
		panic(err)
	}
}

// GetDateTime serves the html file
func GetDateTime(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/dateTimePicker.html")
}

// PostDateTime 
func PostDateTime(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	datetime, err := time.Parse(time.RFC3339, r.FormValue("datetime"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedDateTime := datetime.Add(time.Hour * time.Duration(12))

	type data struct {
		OriginalDateTime time.Time `json:"original_date_time"`
		UpdatedDateTime  time.Time `json:"updated_date_time"`
	}

	d := data{OriginalDateTime: datetime, UpdatedDateTime: updatedDateTime}
	outputParsed, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(outputParsed)
}
