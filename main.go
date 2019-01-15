package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

const BUFSIZE = 1024

func float32FromBytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return  float
}

func importImage(p string) ([]float32,error) {
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf := make([]byte, BUFSIZE)
	var r []float32
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return nil, err
		}
		for i := 0; i < n; i += 4 {
			end := i + 4
			if n < end{
				end = n
			}
			r = append(r, float32FromBytes(buf[i:end]))
		}
	}
	return r, nil
}

func main() {
	var(
		o = flag.String("original", "original.img","the path of original image")
		c = flag.String("compare", "compare.img", "the path of image which compare to original one")
	)
	flag.Parse()
	fmt.Printf("[file path] original => %s, compare => %s\n",*o ,*c)

	originalImg, err := importImage(*o)
	if err != nil {
		log.Fatal(err)
	}

	compareImg, err := importImage(*c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[file size] original => %d, compare => %d\n", len(originalImg), len(compareImg))

}
