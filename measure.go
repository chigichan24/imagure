package main

import "math"

func calcPsnr(a []float32, b []float32) float64 {
	errSum := 0.0
	mx := 0.0
	for i := range a {
		errSum += float64(a[i]-b[i]) * float64(a[i]-b[i])
		mx = math.Max(mx, float64(a[i]))
	}
	mse := errSum / float64(len(a))
	return 20 * math.Log10(mx/math.Sqrt(mse))
}

func calcRmse(a []float32, b []float32) float64 {
	errSum := 0.0
	for i := range a {
		errSum += float64(a[i]-b[i]) * float64(a[i]-b[i])
	}
	mse := errSum / float64(len(a))
	return math.Sqrt(mse)
}

func mirror(x int, min int, max int) int {
	for x < min || x >= max {
		if x < min {
			x = min + (min - x - 1)
		}
		if x >= max {
			x = max + (max - x - 1)
		}
	}
	return x
}

func calcSsim(a []float32, b []float32) float64 {
	mx := 0.0
	for _, v := range a {
		mx = math.Max(mx, float64(v))
	}

	sum := 0.0

	c1 := (mx * 0.01) * (mx * 0.01)
	c2 := (mx * 0.03) * (mx * 0.03)

	m := 2

	h := int(math.Sqrt(float64(len(a))))
	w := h

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ave1 := 0.0
			ave2 := 0.0
			var1 := 0.0
			var2 := 0.0
			cov := 0.0

			for dy := -m; dy < m; dy++ {
				for dx := -m; dx < m; dx++ {
					mx := mirror(x+dx, 0, w)
					my := mirror(y+dy, 0, h)

					ave1 += float64(a[mx+my*w])
					ave2 += float64(b[mx+my*w])
				}
			}

			ave1 /= float64((m*2.0 + 1.0) * (m*2.0 + 1.0))
			ave2 /= float64((m*2.0 + 1.0) * (m*2.0 + 1.0))

			for dy := -m; dy < m; dy++ {
				for dx := -m; dx < m; dx++ {
					mx := mirror(x+dx, 0, w)
					my := mirror(y+dy, 0, h)

					var1 += (float64(a[mx+my*w]) - ave1) * (float64(a[mx+my*w]) - ave1)
					var2 += (float64(b[mx+my*w]) - ave2) * (float64(b[mx+my*w]) - ave2)
					cov += (float64(a[mx+my*w]) - ave1) * (float64(b[mx+my*w]) - ave2)
				}
			}

			var1 /= float64((m*2.0 + 1.0) * (m*2.0 + 1.0))
			var2 /= float64((m*2.0 + 1.0) * (m*2.0 + 1.0))
			cov /= float64((m*2.0 + 1.0) * (m*2.0 + 1.0))

			sum += ((2.0*ave1*ave2 + c1) * (2.0*cov + c2)) / ((ave1*ave1 + ave2*ave2 + c1) * (var1 + var2 + c2))

		}
	}
	return sum / float64(w*h)
}
