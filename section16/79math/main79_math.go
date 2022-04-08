package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("数学的な定数-----")
	fmt.Println("円周率", math.Pi)
	// 2の平方根
	fmt.Println("2の平方根", math.Sqrt2)
	// 数値型に関する定数
	fmt.Println("float32の最大値", math.MaxFloat32)
	fmt.Println("float32の0でない最小値", math.SmallestNonzeroFloat32)
	fmt.Println("float64の最大値", math.MaxFloat64)
	fmt.Println("float64の0でない最小値", math.SmallestNonzeroFloat64)
	fmt.Println("int64の最大値", math.MaxInt64)
	fmt.Println("int64の最小値", math.MinInt64)
	/*
	   円周率 3.141592653589793
	   2の平方根 1.4142135623730951
	   float32の最大値 3.4028234663852886e+38
	   float32の0でない最小値 1.401298464324817e-45
	   float64の最大値 1.7976931348623157e+308
	   float64の0でない最小値 5e-324
	   int64の最大値 9223372036854775807
	   int64の最小値 -9223372036854775808
	*/

}
