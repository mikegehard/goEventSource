package main

import (
	"net/http"
	"github.com/msgehard/eventsource"
	"time"
	"html/template"
)

const homepageHtml = `
<html>
  <head>
    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
  	<script>
		var source = new EventSource("/events");
		source.addEventListener('message', function(e) {
		  console.log(e);
		  $("#messages").append("<p>" + e.data + "</p>");
		}, false);
  	</script>
  </head>
  <body>
    Look at all of the messages from the server:
    <div id="messages"/>
  </body>
</html>
`

var homepageTemplate = template.Must(template.New("home").Parse(homepageHtml))

func eventHandler(es *eventsource.Conn) {
	for {
		select {
		case <-time.After(2*time.Second):
			es.Write("Hello world")
		}
	}
}

func homePage(w http.ResponseWriter, req *http.Request) {
	homepageTemplate.Execute(w, nil)
}

func main() {
	http.Handle("/events", eventsource.Handler(eventHandler))
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":9001", nil)
}
