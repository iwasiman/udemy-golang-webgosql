package main //section3
import "fmt"

func main() {
	fmt.Println("interface型")
	// 全ての型と互換性がある。PythonのNone
	var x interface{}
	fmt.Println("xは?", x) // <nil>
	x = 1
	fmt.Println("xに代入", x) // 変わっていく
	x = 3.14
	fmt.Println("xに代入", x) // 変わっていく
	x = "aaa"
	fmt.Println("xに代入", x) // 変わっていく
	x = [3]int{1,2,3}
	fmt.Println("xに代入", x) // 変わっていく
	x = 2 // しかしinterface型と他の型の演算はできない。
	//fmt.Println(x + 3) // invalid operation: x + 3 (mismatched types interface{} and int)
}