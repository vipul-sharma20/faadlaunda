package main

import (
    "net/http"
)

func (options *Faad) url_handler() {
    http.HandleFunc("/faad", buildTemplate)
    http.Handle("/", http.FileServer(http.Dir(*options.directory)))
}
