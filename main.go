package main

import (
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head>
			<title>AOE3 Asian Dynasties</title>
		</head>
		<body>
			<h1>Hello, world!</h1>
			<button>TESTE</button>
		</body>
	</html>
	`), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
