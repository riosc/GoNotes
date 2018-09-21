package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	scheme = "http://"
)

func main() {
	for _, url := range os.Args[1:] {

		if strings.HasPrefix(url, scheme) == false {
			url = scheme + url
		}

		resp, error := http.Get(url)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", error)
			os.Exit(1)
		}

		// b, error := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()

		copyResp, error := io.Copy(os.Stdout, resp.Body)

		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, error)
			os.Exit(1)
		}

		fmt.Printf("%s -- StatusCode: %s\n", string(copyResp), resp.Status)
	}
}
