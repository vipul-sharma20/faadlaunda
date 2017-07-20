package main

import (
    "io/ioutil"
    "net/http"
    "html/template"
    "encoding/xml"
)

type Context struct {
    Url string `xml:"data>images>image>url"`
}

func catRequest() []byte {
    request, _ := http.NewRequest("GET", "http://thecatapi.com/api/images/get", nil)

    query_string := request.URL.Query()
    query_string.Add("type", "jpg,png")
    query_string.Add("format", "html")
    query_string.Add("size", "med,full")
    query_string.Add("format", "xml")

    request.URL.RawQuery = query_string.Encode()
    resp, _ := http.Get(request.URL.String())
    body, _ := ioutil.ReadAll(resp.Body)

    return body
}

func buildTemplate(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("cats.html")
    body := catRequest()
    var context Context

    xml.Unmarshal(body, &context)
    t.Execute(w, context)
}

