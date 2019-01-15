package main

import "math"

func calcPsnr(a []float32, b []float32) float64 {
	errSum := 0.0
	mx := 0.0
	for i := range a {
		errSum += float64(a[i]-b[i])*float64(a[i]-b[i])
		mx = math.Max(mx, float64(a[i]))
	}
	mse := errSum / float64(len(a))
	return 20*math.Log10(mx/math.Sqrt(mse))
}

func calcRmse(a []float32, b []float32) float64 {
	errSum := 0.0
	for i := range a {
		errSum += float64(a[i]-b[i])*float64(a[i]-b[i])
	}
	mse := errSum / float64(len(a))
	return math.Sqrt(mse)
}