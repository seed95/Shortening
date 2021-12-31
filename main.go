package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/seed95/shortening/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("cannot run the app, why? %v\n", aurora.Red(err))
	}
}
