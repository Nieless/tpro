package main

import (
	"encoding/json"
	"fmt"
	"github.com/Nieless/tpro/tpro"
	"github.com/aquilax/go-perlin"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"sync"
)

const (
	alpha       = 2.0
	beta        = 2.0
	n           = 3
	seed  int64 = 100
)

func main() {

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/noisecalc", GetPerlinImageNoiseCalcHandler).Methods("GET")
	myRouter.HandleFunc("/noisecalc", PerlinImageNoiseCalcHandler).Methods("POST")

	err := tpro.ListenAndServe(tpro.MustGetEnv("PERLIN_IMAGE_NOISE_CALC_SERVICE_HTTP_PORT"), myRouter)
	if err != nil {
		panic(err)
	}
}

// GetPerlinImageNoiseCalcHandler serves the html file
func GetPerlinImageNoiseCalcHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/perlinImageNoiseCalculator.html")
}

// PerlinImageNoiseCalcHandler returns noise of image coordinates
func PerlinImageNoiseCalcHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a, err := strconv.Atoi(r.FormValue("x"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(r.FormValue("y"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p := perlin.NewPerlin(alpha, beta, n, seed)

	var wg sync.WaitGroup
	noises := make([]float64, a*b)
	coordinates := make([]string, a*b)

	for x := 0; x < a; x++ {
		for y := 0; y < b; y++ {

			// maintain index to avoid appending discrepancy
			index := b*x + y

			// Increment the WaitGroup counter.
			wg.Add(1)
			go func(p *perlin.Perlin, x, y int) {
				// Decrement the counter when the goroutine completes.
				defer wg.Done()

				noise := CalculateNoise(p, x, y)
				noises[index] = noise
				coordinates[index] = fmt.Sprintf("(%v, %v)", strconv.Itoa(x), strconv.Itoa(y))
			}(p, x, y)
		}
	}

	// Wait for noise calculation of all coordinates complete
	wg.Wait()

	type data struct {
		Coordinates []string  `json:"coordinates"`
		Noises      []float64 `json:"noises"`
	}

	d := data{Noises: noises, Coordinates: coordinates}
	outputParsed, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(outputParsed)
}

func CalculateNoise(p *perlin.Perlin, x, y int) float64 {
	noise := p.Noise2D(float64(x)/10, float64(y)/10)
	return noise
}
