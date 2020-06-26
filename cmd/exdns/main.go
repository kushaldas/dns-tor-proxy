package main

import (
	"fmt"

	"github.com/kushaldas/exampledns/pkg/dserver"
)

func main(){
	fmt.Printf("Hello World!")
	dserver.Listen();
}
