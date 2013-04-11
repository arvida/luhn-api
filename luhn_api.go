package main

import (
	"encoding/json"
	"github.com/joeljunstrom/go-luhn"
	"log"
	"net/http"
	"strconv"
  "os"
)

const (
	ApiVersion = "v1"
)

type AboutResponse struct {
	Name    string
	Version string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := AboutResponse{
		Name:    "luhn_api",
		Version: ApiVersion,
	}

	json_response, _ := json.Marshal(response)

	w.Write(json_response)
}

type ValidationResponse struct {
	Valid bool
}

func validationHandler(w http.ResponseWriter, r *http.Request) {
	subject := r.FormValue("luhn")

	w.Header().Set("Content-Type", "application/json")

	valid := luhn.Valid(subject)
	response := ValidationResponse{
		Valid: valid,
	}

	json_response, _ := json.Marshal(response)

	w.Write(json_response)
}

type GenerationResponse struct {
	Luhn string
}

func generationHandler(w http.ResponseWriter, r *http.Request) {
	size, _ := strconv.Atoi(r.FormValue("size"))

	w.Header().Set("Content-Type", "application/json")

	luhn := luhn.Generate(size)
	response := GenerationResponse{
		Luhn: luhn,
	}

	json_response, _ := json.Marshal(response)

	w.Write(json_response)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/validate", validationHandler)
	http.HandleFunc("/generate", generationHandler)

  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "8000"
  }

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
