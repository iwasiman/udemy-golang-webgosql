package main // section16

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// ファイル操作
	// Create
	f, _ := os.Create("foo.txt")
	f.Write([]byte("Heloo\n"))
	f.WriteAt([]byte("Go-lang\n"), 10)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Happy Friday!")

	// udemy-golang-webgosql\section16\77os_create\foo.txt が生成される。
	// 既に存在していると再作成になるので注意。

	// Read
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs := make([]byte, 128)

	n, err := f.Read(bs)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)          // 読み込んだバイト数
	fmt.Println(string(bs)) //読み込んだ内容

	bs2 := make([]byte, 128)
	n2, err := f.ReadAt(bs2, 10) // 10バイト目から読み込んでいる
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n2)
	fmt.Println(string(bs2))

	// 詳細に指定したファイルの読み込み
	f3, err := os.OpenFile("foo.text", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f3.Close()

	bs3 := make([]byte, 128)
	n3, err := f3.Read(bs3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n3)
	fmt.Println(string(bs3))

}
