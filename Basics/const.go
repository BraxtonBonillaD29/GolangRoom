package main

import (
	"fmt"
	"math"
)

const firstConstantValue float32 = 10

func main() {
	const calculatedConst = firstConstantValue + (math.Pi)
	fmt.Println(calculatedConst)
}
