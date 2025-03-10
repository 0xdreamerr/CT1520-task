package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"

	"github.com/0xdreamerr/url-shortener/config"

	"github.com/go-chi/chi/v5"
)

var urls = make(map[string]string)

func getShortURL(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		h := sha256.New()
		h.Write([]byte(body))
		hash := "/" + hex.EncodeToString(h.Sum(nil))
		addr := config.Config.ResultAddr

                result := addr[:len(addr)-5] + hash[:8]

		res.Header().Set("content-type", "text/plain")
		res.WriteHeader(http.StatusCreated)

		res.Write([]byte(result))
		urls[hash[:8]] = string(body)
	} else {
		http.Error(res, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
}

func redirectTo(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		shortURL := req.URL.String()

		result := urls[shortURL]
		if result == "" {
			http.Error(res, "ShortURL not found", http.StatusNotFound)
			return
		}

		res.Header().Set("Location", result)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Error(res, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	config.SetConfig()

	r := chi.NewRouter()

	r.Post("/", getShortURL)
	r.Get("/{id}", redirectTo)

	fmt.Printf("Server started at: %s", config.Config.ServerAddr)
	err := http.ListenAndServe(config.Config.ServerAddr, r)
	if err != nil {
		panic(err)
	}
}
