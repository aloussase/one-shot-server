package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Opts struct {
	Port     int
	Path     string
	Status   int
	Body     string
	Help     bool
	Fallback string
}

func main() {
	var opts Opts

	flag.StringVar(&opts.Path, "path", "/", "The path from which to serve the resource")
	flag.StringVar(&opts.Fallback, "fallback", "/", "The URL to fallback to for unmapped paths")
	flag.IntVar(&opts.Port, "port", 3000, "The port from which to start the server")
	flag.IntVar(&opts.Status, "status", 200, "The status code with which to respond")
	flag.StringVar(&opts.Body, "body", "", "A path to a file containing the resource to serve as JSON. May be omitted.")
	flag.BoolVar(&opts.Help, "help", false, "Print usage information")

	flag.Parse()

	if opts.Help {
		fmt.Println("Usage: one-shot-server [OPTIONS]\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	http.HandleFunc(opts.Path, func(w http.ResponseWriter, req *http.Request) {
		if opts.Body == "" {
			w.WriteHeader(opts.Status)
			return
		}

		data, err := os.ReadFile(opts.Body)
		if err != nil {
			log.Fatal("Failed to read body from path: " + opts.Body)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(opts.Status)
		w.Write(data)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if opts.Fallback == "" {
			w.WriteHeader(404)
			return
		}

		redirectTo := opts.Fallback + req.URL.Path
		log.Printf("Redirecting request to: %s", redirectTo)

		w.Header().Set("Location", redirectTo)
		w.WriteHeader(302)
	})

	fmt.Printf(
		"One-shot server running on port %d\n\n"+
			"== SERVER CONFIGURATION ==\n\n"+
			"> Serving resource on path: %s\n"+
			"> Will reply with status: %d\n",
		opts.Port, opts.Path, opts.Status)

	if opts.Body != "" {
		fmt.Printf("> Will server payload from: %s\n\n", opts.Body)
	} else {
		fmt.Printf("> No payload configured\n\n")
	}

	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(opts.Port), nil))
}
