package main // section16

import (
	"fmt"
	"os"
)

func main() {
	//Argsを使う
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])
	fmt.Println(os.Args[4])

	fmt.Println("要素数", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}

/*
// 引数の０番目はコマンド自体が入る。
> go run main77_args.go aa bb cc dd ee
D:\Users\xxx\AppData\Local\Temp\go-build23134668\b001\exe\main77_args.exe
aa
bb
cc
dd
要素数 6
0 D:\Users\xxxAppData\Local\Temp\go-build23134668\b001\exe\main77_args.exe
1 aa
2 bb
3 cc
4 dd
5 ee

*/
