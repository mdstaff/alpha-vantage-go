package main

import (
	"fmt"

	"github.com/mdstaff/alpha-vantage-go/client"
)

func main(){
	fmt.Println("av.go")
	nc := client.NewClient()
	nc.GetQuote("T")
}