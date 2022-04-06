package main //section11
import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

// type Users struct でフィールドに[]*Userを持つのもできるが、下の方が望ましい。

type Users []*User

func main() {
	fmt.Println("セクション11: 構造体")
	fmt.Println("struct スライス")

	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}
	fmt.Println(users) // []
	users = append(users, &user1)
	//users = append(users, user2) // Usersは各要素をポインタ型で定義しているので、cannot use user2 (variable of type User) as *User value in argument to append
	users = append(users, &user2, &user3, &user4)
	fmt.Println(users) // [アドレス アドレス アドレス アドレス]

	for _, u := range users {
		fmt.Println(u)        // &{user1 10}
		fmt.Println(" ", *u)  // {user1 10}
		fmt.Println("  ", &u) // アドレス
		// ... が4回
	}

	// 組み込み関数makeを使っても同じ。
	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)
	for _, u := range users2 {
		fmt.Println(u)        // &{user1 10}
		fmt.Println(" ", *u)  // {user1 10}
		fmt.Println("  ", &u) // アドレス
		// ... が2回
	}

}
