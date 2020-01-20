package main

import (
	"encoding/json"
	"github.com/Nieless/tpro/tpro"
	"github.com/gorilla/mux"
	"github.com/grokify/html-strip-tags-go"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/crawling", GetWordCounter).Methods("GET")
	myRouter.HandleFunc("/crawling", PostWordCounter).Methods("POST")

	err := tpro.ListenAndServe(tpro.MustGetEnv("WORD_COUNTER_SERVICE_HTTP_PORT"), myRouter)
	if err != nil {
		panic(err)
	}
}

// GetWordCounter serves the html file
func GetWordCounter(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/crawler.html")
}

// PostWordCounter counts the word and return map
func PostWordCounter(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := r.FormValue("url")
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	readAll, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Body.Close()

	stripped := strip.StripTags(string(readAll))

	type data struct {
		WordCounts map[string]int `json:"word_counts"`
		URL        string         `json:"url"`
	}

	d := data{WordCounts: wordCount(stripped), URL: url}
	outputParsed, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(outputParsed)
}

func wordCount(str string) map[string]int {
	wordList := strings.Fields(str)

	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	return counts
}
