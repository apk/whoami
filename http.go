package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	if ((r.Method == http.MethodGet ||
	    r.Method == http.MethodHead) && r.URL.Path == "/") {
	    en := "REDIR_" + r.Host;
	    ev := os.Getenv(en);
	    if (ev != "") {
		w.Header().Add("Location", ev);
		w.WriteHeader (http.StatusMovedPermanently);
	    }
	}
        fmt.Fprintf(os.Stdout, "%s: %s %s\n", r.Host, r.Method, r.URL.Path)
        fmt.Fprintf(w, "I'm %s\n", hostname)

        fmt.Fprintf(w, "\n")
        fmt.Fprintf(w, "Method: %s\n", r.Method)
        fmt.Fprintf(w, "Proto: %s\n", r.Proto)
        fmt.Fprintf(w, "Host: %s\n", r.Host)
        fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
        fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
        fmt.Fprintf(w, "\n")
        for k, v := range r.Header {
            for _, a := range v {
                 fmt.Fprintf(w, "%s: %s\n", k, a)
            }
        }
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))
}

