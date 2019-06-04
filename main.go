package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	addrFlag = flag.String("addr", ":5555", "server address:port")
)

func main() {
	flag.Parse()
	log.Printf("please visit %q", *addrFlag)

	http.HandleFunc("/", root)
	err := http.ListenAndServe(*addrFlag, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, page)
}

const page = `
<html>
	<head>
		<title>Sirhus Viewer</title>
		<script>
function count(str) {
	var n1 = document.getElementById("top");
	n1.innerHTML = "Number of characters: "+ str.value.length;
	var n2 = document.getElementById("bottom");
	n2.innerHTML = "Number of characters: "+ str.value.length;
};
		</script>
	</head>

	<body>
		<div id="top">Number of characters: 0</div>
		<div id="content">
			<textarea contenteditable="true" id="text" rows="50" cols="80" onkeyup="count(this)" onpaste="count(this)"></textarea>
		</div>
		<div id="bottom">Number of characters: 0</div>
	</body>
</html>
`
