package main

import (
    "flag"
    "github.com/pomack/webmachine.go/webmachine"
    "io"
    "log"
    "net/http"
    "strconv"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
}

func main() {
    directory := "."
    urlPathPrefix := "/"
    allowWrite := false
    allowDirectoryListing := false
    port := 12345
    flag.StringVar(&directory, "dir", ".", "Directory to serve")
    flag.StringVar(&urlPathPrefix, "path", "/", "URL Path Prefix")
    flag.BoolVar(&allowWrite, "write", false, "Allow and process PUT, POST, DELETE without any authorization")
    flag.BoolVar(&allowDirectoryListing, "listing", false, "Allow Directory Listing on GET and HEAD")
    flag.IntVar(&port, "port", 12345, "Port to serve files")
    flag.Parse()
    wm := webmachine.NewWebMachine()
    wm.AddRouteHandler(webmachine.NewFileResource(directory, urlPathPrefix, allowWrite, allowDirectoryListing))
    err := http.ListenAndServe(":"+strconv.Itoa(port), wm)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
