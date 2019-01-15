package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func importImage(p string) ([]float32,error) {
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return nil, nil
}

func main() {
	var(
		o = flag.String("original", "original.img","the path of original image")
		c = flag.String("compare", "compare.img", "the path of image which compare to original one")
	)
	flag.Parse()
	fmt.Printf("original => %s, compare => %s",*o ,*c)


	if _,err := importImage(*o); err != nil {
		log.Fatal(err)
	}

	if _,err := importImage(*c); err != nil {
		log.Fatal(err)
	}

}
