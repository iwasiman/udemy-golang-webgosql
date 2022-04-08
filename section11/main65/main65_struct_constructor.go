package main //section11
import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

// 戻り値のところで*を使いポインタ型を返す。
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	fmt.Println("セクション11: 構造体")
	fmt.Println("struct 型コンストラクタ")

	user1 := NewUser("user1", 10)
	fmt.Println("user1", user1)  // &{user1 10}
	fmt.Println("user1", *user1) // {user1 10} 変数user1が指している先を見る
	fmt.Println("user1", &user1) // 0xcうんたら
}
