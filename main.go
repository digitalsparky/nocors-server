package main

/**
 * This is a basic file server that sends CORS headers so you can access external resources in a local testing environment
 * Mostly useful for dev'ing with React, VueJS, or other client side libraries
 *
 * Written by DigitalSparky - github.com/digitalsparky
**/

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func addCORSOverride(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		h.ServeHTTP(w, r)
	}
}

func Serve(listenAddr string, listenPort int, filePath string) error {
	fmt.Printf("Server started at: http://%s:%d\n", listenAddr, listenPort)
	http.Handle("/", addCORSOverride(http.FileServer(http.Dir(filePath))))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", listenAddr, listenPort), nil))
	return nil
}

func main() {
	app := cli.NewApp()

	app.Name = "nocors-server"
	app.Usage = "Serve local files (like React, VueJS, etc) with valid CORS * exemption"

	if (os.Getenv("CI") == "true") && (os.Getenv("TRAVIS") == "true") {
		if os.Getenv("TRAVIS_TAG") != "" {
			app.Version = os.Getenv("TRAVIS_COMMIT")
		} else {
			app.Version = os.Getenv("TRAVIS_TAG")
		}
	} else {
		app.Version = "development"
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "listenAddr, a",
			Value:  "127.0.0.1",
			Usage:  "Listen Address",
			EnvVar: "LISTEN_ADDRESS",
		},
		cli.IntFlag{
			Name:   "listenPort, p",
			Value:  4000,
			Usage:  "Listen Port",
			EnvVar: "LISTEN_PORT",
		},
		cli.StringFlag{
			Name:   "servePath, s",
			Value:  ".",
			Usage:  "Static File Serve Path",
			EnvVar: "SERVE_PATH",
		},
	}

	app.Action = func(c *cli.Context) error {
		return Serve(c.String("listenAddr"), c.Int("listenPort"), c.String("servePath"))
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
