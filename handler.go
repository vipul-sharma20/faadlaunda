package main

import (
    "net/http"
)

func (options *Faad) url_handler() {
    http.HandleFunc("/", buildTemplate)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
