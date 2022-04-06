package main //section11
import (
	"fmt"
)

type User struct {
	Name string
	Age  int
	X, Y int //まとめて定義もできる
}

// 引数は構造型の実体
func UpdateUser(user User) {
	// コピーが作成されてその構造体が操作される。
	user.Name = "変えちゃうぞ"
	user.Age = 1000
}

// こちらは引数がポインタ、アドレスを渡す
func UpdateUserByAddress(user *User) {
	user.Name = "変えちゃうぞ～"
	user.Age = 1000
}

func main() {
	fmt.Println("セクション11: 構造体")
	fmt.Println("struct")
	// クラス的なもの。出力すると波カッコで囲われる。

	var user1 User
	fmt.Println(user1) // {(空文字) 0 0 0 } 全て初期値
	user1.Name = "user1"
	user1.Age = 100
	fmt.Println(user1) // {user1 100 0 0 }

	user2 := User{}             // 波カッコをつける
	fmt.Println("user2", user2) // {(空文字) 0 0 0 } 全て初期値
	user2.Name = "user2"
	user2.Age = 999
	fmt.Println("user2", user2)             // {user2 999 0 0 }
	user3 := User{Name: "user3だよ", Age: 20} // コンストラクタ的に入れられる
	fmt.Println("user3", user3)             //
	user4 := User{"user4だよ", 15, 1, 2}      // フィールドなしの場合は順番が一致しないとエラー
	fmt.Println("user4", user4)             //

	user6 := User{Name: "ユーザー6"} // フィールドの一部だけも入れられる
	fmt.Println("user6", user6)  //

	user7 := new(User)          // newを使うとアドレスが入る
	fmt.Println("user7", user7) // &{ 0 0 0}

	user8 := &User{}            // &を使って代入してもnewと同じ。
	fmt.Println("user8", user8) // &{ 0 0 0}
	// 関数の引数に構造値を渡すとき、参照渡しにするためにポインタをよく使う。

	UpdateUser(user6)
	fmt.Println("user6 関数通過後", user6) // 変わらない
	UpdateUserByAddress(user7)
	fmt.Println("user7 関数通過後", user7, &user7) // 変わる。&user7=番地

}
