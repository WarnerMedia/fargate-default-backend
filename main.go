// Copyright (c) Warner Media, LLC. All rights reserved. Licensed under the MIT license.
// See the LICENSE file for license information.

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/health", defaultHandler)
	http.Handle("/", r)

	var port = getenv("PORT", "8001")
	fmt.Println("Listening on: " + port)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	if os.Getenv("ENABLE_LOGGING") == "true" {
		fmt.Println("Logging enabled")
		http.ListenAndServe(":"+port, loggedRouter)
	} else {

		http.ListenAndServe(":"+port, nil)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the Fargate default backend."))
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
