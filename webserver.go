// forms.go
package main

import (
    "html/template"
    "net/http"
	"fmt"
)

type ContactDetails struct {
	Success bool
    Email   string
    Subject string
    Message string
}

type city1 struct {
    name   string
    lat string
    long string
    temp int
}

type city2 struct {
    name   string
    lat string
    long string
    temp int
}


func main() {
    tmpl := template.Must(template.ParseFiles("./static/index.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
			Success: true,
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // do something with details
        _ = details
		fmt.Println(details.Email)

        tmpl.Execute(w, details)
    })

    http.ListenAndServe(":8080", nil)
}