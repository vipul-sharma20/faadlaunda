/*
Usage:
    -p="8100": port to serve on
    -d=".":    the directory of static files to host
*/
package main

import (
    "flag"
    "log"
    "net/http"
)

type Faad struct {
    port *string
    directory *string
}

func main() {
    port := flag.String("p", "8100", "port to serve on")
    directory := flag.String("d", ".", "the directory of static file to host")

    flag.Parse()
    faad := Faad{port: port, directory: directory}
    faad.url_handler()

    log.Printf("Serving faadness on HTTP port: %s\n", *port)
    log.Fatal(http.ListenAndServe(":"+*port, nil))
}
