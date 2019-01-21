package main

import (
	"fmt"
	"os"

	infweb "github.com/tampopos/go.api-template/src/3-framework-and-drivers/web"
)

func main() {
	web, err := infweb.NewWeb()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	web.Start()
}
