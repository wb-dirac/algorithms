package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
)

func main() {
	Tn1 := qudric(-2)
	x := -0.618
	for n := 0; n < 100; n++ {
		x = Tn1(x)
		if n%4 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%f,", x)
	}
	value := 10.8
	//mul10000 := decimal.NewFromFloat(value).Mul(decimal.NewFromInt(100))
	mul10000 := decimal.NewFromFloat(value).Mul(decimal.NewFromFloat(math.Pow10(2)))
	fmt.Println(mul10000.Floor().Equal(mul10000))
}

func qudric(c float64) func(float64) float64 {
	return func(x float64) float64 {
		return x*x + c
	}
}

func threeNPlusOne() {
	var i, n int
	var step int
	var stemSum int
	var count = 2 << 32
	//var rateSum float64
	for i = 2; i < count; i = i<<1 + 1 {
		n = i
		step = 0
		for n != 1 {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = n*3 + 1
			}
			step++
		}
		stemSum += step
		//rateSum +=
		//fmt.Printf("%f, ", rateSum/float64(i))
		fmt.Printf("%f, ", float64(step)/math.Log(float64(i)))
	}
}
