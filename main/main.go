package main

import (
	"net/http"
	"strings"
	"github.com/zserge/webview"
	"fmt"
	"net"
	"log"
)

const PORT = 8765
const WIDTH = 800
const HEIGHT = 600
const TITLE = "Webview Application"

type writeHtml func(w http.ResponseWriter, r *http.Request)

func main() {
	url := startServer()
	w := webview.New(webview.Settings{
		Width:     WIDTH,
		Height:    HEIGHT,
		Title:     TITLE,
		Resizable: false,
		URL:       url,
	})
	w.SetColor(255, 255, 255, 255)
	defer w.Exit()
	w.Run()
}

func startServer() string {
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", PORT))
	if err != nil {
		log.Fatal(err)
	}

	go handle("/", index, ln)
	go handle("/leo", leo, ln)

	return "http://" + ln.Addr().String()
}

func handle(pattern string, function writeHtml, ln net.Listener) {
	defer ln.Close()
	http.HandleFunc(pattern, function)
	log.Fatal(http.Serve(ln, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	message = message + "<a href='/leo'>Go To Leo</a>"
	w.Write([]byte(message))
}

func leo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	message := "Leo "
	message = message + "<button onclick='history.back();' style='color: black'>Go Back</button>"
	w.Write([]byte(message))
}