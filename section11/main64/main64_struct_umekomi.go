package main //section11
import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type T struct {
	User User // User型を埋め込む。
}
type T2 struct {
	User // 型名=フィールド名は型名を省略できる
}

// 基本的にメソッドではポインター型で受けるべき
func (u *User) SetNameByPointer(name string) {
	u.Name = name // ここで変数名はu.でよい
}

func main() {
	fmt.Println("セクション11: 構造体")
	fmt.Println("struct 埋め込み")

	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println("t:", t)              // {{user1 10}}
	fmt.Println("tの中へ:", t.User)      // {user1 10}
	fmt.Println("tの中へ:", t.User.Name) // user1

	t2 := T2{User: User{Name: "user2", Age: 10}}
	fmt.Println("t2の中へ:", t2.Name) // user2 フィールド名を省略して内側へアクセスできる。補完でもちゃんと出る!
	t.User.SetNameByPointer("t1の新しい名前")
	t2.SetNameByPointer("t2の新しい名前")
	fmt.Println("setName後 t:", t)   // {{t1の新しい名前 10}}
	fmt.Println("setName後 t2:", t2) // {{t2の新しい名前 10}}

}
