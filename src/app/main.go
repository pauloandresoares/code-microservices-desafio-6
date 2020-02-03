package main

import (
	"fmt"
	"math"
	"net/http"
)

func calculo(x float64) float64 {
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}

	return x
}

func handler(w http.ResponseWriter, r *http.Request) {
	calculo(0.0001)

	fmt.Fprintf(w, "Code.education Rocks!")
}

func main() {
	fmt.Println("Go looping...")

	http.HandleFunc("/", handler)

	http.ListenAndServe(":80", nil)
}
