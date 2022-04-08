package main //section3
import "fmt"

func main() {
	// byte(uint8)型
	byteA := []byte{72, 73}
	fmt.Println("byte(uint8)型", byteA)
	fmt.Println("アスキーコードなので文字列に変換", string(byteA))

	byteC := []byte("HI")
	fmt.Println("文字列からbyte型生成", byteC)
	fmt.Println("文字列からbyte型生成してまた文字列に戻す", string(byteC))

	// 18.配列型
	// 数を後から変えられない(!)のがGoの特徴。
	var arr1 [3]int
	fmt.Println("配列の基本的な宣言", arr1)
	fmt.Printf("%T\n", arr1) // 型は[3]int になる
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println("配列の初期値付き宣言", arr2) //[A B 空文字]
	arr3 := [3]int {1,2,3}
	fmt.Println("配列の初期値付き暗黙的宣言", arr3) //[1 2 3]
	arr4 := [...]string{"C", "D"}
	fmt.Println("配列の数を...で省略して宣言", arr4) //[C D]
	fmt.Printf("%T\n", arr4) // 型は [2]string になる

	// 配列の操作
	// 値の取り出し
	fmt.Println("配列arr2の中身をインデックス番号で取り出す", arr2[0]) // A
	fmt.Println("配列arr2の中身をインデックス番号で取り出す", arr2[1]) // B
	fmt.Println("配列arr2の中身をインデックス番号で取り出す", arr2[2]) // 空文字

	// 値の更新
	arr2[2] = "Cにかえるよ"
	fmt.Println("配列arr2の中身を変更", arr2) // [A B Cにかえるよ]
	//var arr5 [4]int // 要素数が違うと代入できない。別の型扱い。
	//arr5 = arr1 // cannot use arr1 (variable of type [3]int) as [4]int value in assignment
	fmt.Println("配列arr2の長さをlen関数で表示", len(arr2)) // 3

}