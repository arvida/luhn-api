package main

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
  "encoding/json"
)

func TestIndex(t *testing.T){
  server := httptest.NewServer(http.HandlerFunc(indexHandler))
  defer server.Close()

  response, err := http.Get(server.URL + "/")
  if err != nil {
    t.Fatal(err)
  }

  content_type := response.Header.Get("Content-Type")
  if content_type != "application/json" {
    t.Errorf("expected content_type to be ”application/json” but was ”%s”", content_type)
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    t.Fatal(err)
  }

  var about AboutResponse
  json.Unmarshal(body, &about)

  if about.Name != "luhn_api" {
    t.Errorf("expected body to include ”name: luhn_api”, %v", about) 
  }
  if about.Version != "v1" {
     t.Errorf("expected body to include ”version: v1”, %v", about) 
  }
}

func TestValidationIsValid(t *testing.T){
  server := httptest.NewServer(http.HandlerFunc(validationHandler))
  defer server.Close()

  response, err := http.Get(server.URL + "/validate?luhn=1111111116")
  if err != nil {
    t.Fatal(err)
  }

  content_type := response.Header.Get("Content-Type")
  if content_type != "application/json" {
    t.Errorf("expected content_type to be ”application/json” but was ”%s”", content_type)
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    t.Fatal(err)
  }

  var validation ValidationResponse
  json.Unmarshal(body, &validation)

  if validation.Valid != true {
    t.Errorf("expected validation to be true, was %v", validation) 
  }
}

func TestValidationIsNotValid(t *testing.T){
  server := httptest.NewServer(http.HandlerFunc(validationHandler))
  defer server.Close()

  response, err := http.Get(server.URL + "/validate?luhn=1111111111")
  if err != nil {
    t.Fatal(err)
  }

  content_type := response.Header.Get("Content-Type")
  if content_type != "application/json" {
    t.Errorf("expected content_type to be ”application/json” but was ”%s”", content_type)
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    t.Fatal(err)
  }

  var validation ValidationResponse
  json.Unmarshal(body, &validation)

  if validation.Valid == true {
    t.Errorf("expected validation to be false, was %v", validation) 
  }
}


func TestGenerate(t *testing.T){
  server := httptest.NewServer(http.HandlerFunc(generationHandler))
  defer server.Close()

  response, err := http.Get(server.URL + "/generate?size=8")
  if err != nil {
    t.Fatal(err)
  }

  content_type := response.Header.Get("Content-Type")
  if content_type != "application/json" {
    t.Errorf("expected content_type to be ”application/json” but was ”%s”", content_type)
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    t.Fatal(err)
  }

  var generation GenerationResponse
  json.Unmarshal(body, &generation)

  if len(generation.Luhn) != 8 {
    t.Errorf("expected generated lugn to be 8 in length, was %v", len(generation.Luhn)) 
  }
}
