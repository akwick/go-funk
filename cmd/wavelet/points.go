package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/soundcloud/wavelet/lib/sampler"
	"github.com/soundcloud/wavelet/lib/wav"
)

func points(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("fixtures/fixture.wav")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	samples, err := wav.Samples(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	points, err := sampler.AvgSample(samples, 1200)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(points)
}
