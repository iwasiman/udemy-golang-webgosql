package main //section11
import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

// 関数名の前にレシーバーで (変数名+型) を設定。
func (u User) SayName() {
	fmt.Println("君の名は。", u.Name)
}
func (u User) SetName(name string) {
	u.Name = name
}

// 基本的にメソッドではポインター型で受けるべき
func (u *User) SetNameByPointer(name string) {
	u.Name = name // ここで変数名はu.でよい
}

func main() {
	fmt.Println("セクション11: 構造体")
	fmt.Println("struct メソッド")

	user1 := User{Name: "だれか"}
	user1.SayName() // ドットを打つと紐づいたメソッドも補完候補に出る。素晴らしい!
	user1.SetName("変えるよ")
	user1.SayName()                        // 変わらない
	user1.SetNameByPointer("ポインターで名前変え太郎") // ここでuser1はポインター型でないが使える。
	user1.SayName()                        // これで変わる。

	user2 := &User{Name: "User2"}
	user2.SetNameByPointer("ポインタで宣言した侍")
	user2.SayName() //これで変わる。

}
