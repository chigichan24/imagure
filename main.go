package main

import (
	"flag"
	"fmt"
)

func main() {
	var(
		o = flag.String("original", "original.img","the path of original image")
		c = flag.String("compare", "compare.img", "the path of image which compare to original one")
	)
	flag.Parse()
	fmt.Println(*o, *c)
}
