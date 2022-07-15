package main

import (
	"log"
	"net/http"
)

// main - set up a function handler to send some HTML and CSS to a browser
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
        <html>
        <head>
		<title>GOLANG TEST</title>
		</head>
		<style type="text/css">
        .cell-green {
            background: rgba(0,255,0,0.1);
        }
		</style>
        <body>
		<p class="cell-green">Hello World</p>
		</body>
        </html>`))
	})

	//start the webserver and listen on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("LisenAndServer:", err)
	}
}
