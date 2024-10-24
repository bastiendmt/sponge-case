// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a simple hello, world demonstration web server.
//
// It serves version information on /version and answers
// any other request like /name by saying "Hello, name!".
//
// See golang.org/x/example/outyet for a more sophisticated server.
package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"unicode"

	"github.com/gorilla/mux"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: helloserver [options]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (

	addr     = flag.String("addr", "localhost:8080", "address to serve")
)

func main() {
	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	// Parse and validate arguments (none).
	args := flag.Args()
	if len(args) != 0 {
		usage()
	}

	router := mux.NewRouter()

	router.HandleFunc("/version", version).Methods("GET")
	router.HandleFunc("/api/{input}", spongeCase).Methods("GET")
	http.Handle("/", router)

	log.Printf("serving http://%s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func version(w http.ResponseWriter, r *http.Request) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "no build information available", 500)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
}

// main api that reads the query params
func spongeCase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	input := vars["input"]
	if input == "" {
		input = "Hello world"
	}

	transformedInput := alternateCase(input)

	fmt.Fprintf(w, "<!DOCTYPE html>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(transformedInput))
}

// Alternates case given as an input. Note that the first carater is lowercase.
func alternateCase(input string) string {
	var result strings.Builder
	for i, char := range input {
		if i%2 == 1 {
			result.WriteRune(unicode.ToUpper(char))
		} else {
			result.WriteRune(unicode.ToLower(char))
		}
	}
	return result.String()
}

// get method or post ?
