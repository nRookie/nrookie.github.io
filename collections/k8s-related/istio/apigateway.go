package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	log.Println("Service api-gateway")
	http.HandleFunc("/helloworld", AppRouter)
	http.ListenAndServe(":12344", nil)
}

func AppRouter(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, false)
	log.Printf("%q\n", dump)
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: &transport,
	}

	var p http.Request = *r
	var err error
	p.RequestURI = ""
	p.Host = "helloworld.prj-uhost-ussg.svc.a3.uae:12345"
	upstream := "http://" + p.Host + r.RequestURI
	p.URL, err = url.Parse(upstream)
	resp, err := client.Do(&p)

	if err != nil {
		msg := fmt.Sprintf("Call  hello world failed:%v\n", err)
		io.WriteString(w, msg)
		return
	} else if resp.StatusCode != 200 {
		msg := fmt.Sprintf("Call  hello world failed, status:%s\n", err, resp.Status)
		io.WriteString(w, msg)
		return
	} else {
		dump, _ = httputil.DumpResponse(resp, true)
		log.Printf("%q\n", dump)
	}
	body, err := ioutil.ReadAll(resp.Body)
	w.Write(body)
}
