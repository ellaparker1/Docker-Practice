// @author Ella Parker
// @author Lance Queen

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
		<style>
		body {
			background-color: #FFFDD0;
			}
		p {text-align: center;}
        .cell-green {
            background: rgba(0,255,0,0.1);
        }
		</style>
        <body>
		<p class="cell-green"><h1>ICE SHOPPE</h1></p>
		<p>An RTI Intern-run ice cream shoppe located in Charlotte, NC</p>
		<hr />
		<p><h3>Flavors</h3></p>
		<ul>
  			<li>Chocolate</li>
  			<li>Vanilla</li>
  			<li>That's it. There are no more flavors.</li>
		</ul>
		<style>
		.dot {
  			height: 25px;
  			width: 25px;
  			background-color: #99badd;
  			border-radius: 50%;
  			display: inline-block;
		}
		</style>
		<div style="text-align:center">
  			<span class="dot"></span>
  			<span class="dot"></span>
  			<span class="dot"></span>
		</div>
		<p><em><span style="text-decoration: underline;">This place is amazing!</span></em></p>
		<p>Quote by someone we certainly did not pay to say that ;)</p>
		</body>
		<style>
		.dot {
  			height: 25px;
  			width: 25px;
  			background-color: #99badd;
  			border-radius: 50%;
  			display: inline-block;
		}
		</style>
		<div style="text-align:center">
  			<span class="dot"></span>
  			<span class="dot"></span>
  			<span class="dot"></span>
		</div>
		<p><q>Success is a journey not a destination.</q> -A. Succesful Person</p>
		<style>
		.dot {
  			height: 25px;
  			width: 25px;
  			background-color: #99badd;
  			border-radius: 50%;
  			display: inline-block;
		}
		</style>
		<div style="text-align:center">
  			<span class="dot"></span>
  			<span class="dot"></span>
  			<span class="dot"></span>
		</div>
        </html>`))
	})

	//start the webserver and listen on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("LisenAndServer:", err)
	}
}
