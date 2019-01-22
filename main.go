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
	return float
}

func float64FromBytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	double := math.Float64frombits(bits)
	return double
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

func importImage64(p string) ([]float64,error) {
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf := make([]byte, BUFSIZE)
	var r []float64
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return nil, err
		}
		for i := 0; i < n; i += 8 {
			end := i + 8
			if n < end{
				end = n
			}
			r = append(r, float64FromBytes(buf[i:end]))
		}
	}
	return r, nil
}

func float32ToFloat64(i []float32) []float64 {
	f := make([]float64, len(i))
	for n := range i {
		f[n] = float64(i[n])
	}
	return f
}

func main() {
	var(
		o = flag.String("original", "original.img","the path of original image")
		c = flag.String("compare", "compare.img", "the path of image which compare to original one")
		t = flag.Bool("double", false, "image type. if image is double set true")
	)
	flag.Parse()
	fmt.Printf("[file path] original => %s, compare => %s\n",*o ,*c)
	if *t {
		originalImg, err := importImage64(*o)
		if err != nil {
			log.Fatal(err)
		}

		compareImg, err := importImage64(*c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[file size] original => %d, compare => %d\n", len(originalImg), len(compareImg))

		fmt.Printf("PSNR => %f\n", calcPsnr(originalImg, compareImg))
		fmt.Printf("RMSE => %f\n", calcRmse(originalImg, compareImg))
		fmt.Printf("SSIM => %f\n", calcSsim(originalImg, compareImg))

	} else {
		originalImg, err := importImage(*o)
		if err != nil {
			log.Fatal(err)
		}

		compareImg, err := importImage(*c)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[file size] original => %d, compare => %d\n", len(originalImg), len(compareImg))

		fmt.Printf("PSNR => %f\n", calcPsnr(float32ToFloat64(originalImg), float32ToFloat64(compareImg)))
		fmt.Printf("RMSE => %f\n", calcRmse(float32ToFloat64(originalImg), float32ToFloat64(compareImg)))
		fmt.Printf("SSIM => %f\n", calcSsim(float32ToFloat64(originalImg), float32ToFloat64(compareImg)))
	}

}
