package main //section11
import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("セクション11: 構造体")
	fmt.Println("struct マップ")

	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}
	fmt.Println("キーがintのマップ", m) // map[1:{user1 10} 2:{user2 20}]
	m[3] = User{Name: "user3", Age: 30}
	fmt.Println("キーがintのマップ", m) // map[1:{user1 10} 2:{user2 20} 3:{user3 30}]

	m2 := map[User]string{
		{Name: "key user1", Age: 10}: "Tokyo",
		{Name: "key user2", Age: 20}: "Kanagawa",
	}
	fmt.Println("キーが構造体のマップ", m2) // map[{key user1 10}:Tokyo {key user2 20}:Kanagawa]

	for k, v := range m2 {
		fmt.Println(k, v)
		// {key user1 10} Tokyo
		// {key user2 20} Kanagawa
	}

}
