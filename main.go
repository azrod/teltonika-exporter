package main

import (
	"fmt"

	log "github.com/mgutz/logxi/v1"

	"github.com/common-nighthawk/go-figure"
)

var (
	logger log.Logger
)

func main() {

	myFigure := figure.NewFigure("Ttk exporter", "", true)
	myFigure.Print()
	fmt.Println("Teltonika-exporter")

	
}
