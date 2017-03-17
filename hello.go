package main

import (
	"encoding/json"
	"net/http"
	"time"
	morning "github.com/liuggio/morning-component"
	evening "github.com/liuggio/evening-component"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := time.Now()
	var wts map[string]string
	var err error
	switch {
	case t.Hour() < 12:
		wts, err = morning()
	default:
		wts, err = evening()
	}
	if err != nil {
		wts = map[string]string{"error": err.Error()}
	}

	jData, err := json.Marshal(wts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\":\"world is not a safe place.\"}"))
		return
	}

	w.Write(jData)
}
