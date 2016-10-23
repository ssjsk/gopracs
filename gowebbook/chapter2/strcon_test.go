package main

import(
	"fmt"
	"github.com/ssjsk/gowebbook/strcon"
)

func main(){
	s := strcon.SwapCase("GoPher")
	fmt.Println("Converted string is : ", s)
}