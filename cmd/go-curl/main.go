package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const buildVersion = "dev"
const buildTime = "<missing build time>"
const expectedStatusCode = 200

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s: (version %s, built on %s)\n", os.Args[0], buildVersion, buildTime)
		flag.PrintDefaults()
	}
	var fail bool
	flag.BoolVar(&fail, "f", false, "Fail on non-200 response")
	flag.Parse()

	url := flag.Arg(0)
	if url == "" {
		log.Fatalln("fatal: no url specified")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("fatal: GET %s: %+v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatusCode && fail {
		log.Fatalf("fatal: GET %s: expected %d, got %s", url, expectedStatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("fatal: GET %s: unable to read body: %+v", url, err)
	}

	if len(body) > 0 {
		fmt.Print(string(body))
	}
}
