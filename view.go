package main

import (
    "encoding/xml"
    "html/template"
    "io/ioutil"
    "math/rand"
    "net/http"
    "time"
)

type Context struct {
    Url string `xml:"data>images>image>url"`
    Trait string
    Person string
}

func catRequest() []byte {
    request, _ := http.NewRequest("GET", "http://thecatapi.com/api/images/get", nil)

    query_string := request.URL.Query()
    query_string.Add("type", "jpg,png")
    query_string.Add("format", "html")
    query_string.Add("size", "med")
    query_string.Add("format", "xml")

    request.URL.RawQuery = query_string.Encode()
    resp, _ := http.Get(request.URL.String())
    body, _ := ioutil.ReadAll(resp.Body)

    return body
}

func buildTemplate(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/cats.html")
    body := catRequest()
    var context Context

    rand.Seed(time.Now().Unix())
    context.Trait = traits[rand.Intn(len(traits))]
    context.Person = personalities[rand.Intn(len(personalities))]

    xml.Unmarshal(body, &context)
    t.Execute(w, context)
}

