package main //section13
import (
	"fmt"
)

// "foo"
//could not import foo (cannot find package "foo" in any of
//D:\xxx\go1.18\src\foo (from $GOROOT)
//D:\Users\xxx\go\src\foo (from $GOPATH))/
//"./foo"
// main74_public_private.go:7:2: "./foo" is relative, but relative import paths are not supported in module mode

// 単に”foo” だと、$GOROOTや$GOHOMEの下しか見ない模様。→ go modules を使えば解決する。

func appName() string {
	const AppName = "Iketenai-main"
	var Version string = "1.0"
	return AppName + " " + Version + "だよ"
}

func Do(s string) (b string) {
	//var b string = s 戻り値と同じ名前の変数は再定義できない。
	b = s
	// {}のブロックの中なら、bも新しい変数として再定義できる。{}を抜けると元の値に戻る。
	{
		b := "あたらしいよ"
		fmt.Println(b)
	}
	return b
}

func main() {
	fmt.Println("セクション13: パブリックとプライベートと分割")
	fmt.Println("パブリックとプライベートと分割")

	// 実際の演習では、VSCOdeは自動でインポート補完はされなかった。
	//fmt.Println(foo.Max)
	//fmt.Println(foo.min) // プライベートなのでダメ
	//fmt.Println(foo.ReturnMin()) // 関数経由だとOK。
	//fmt.Println(foo.Max)
	fmt.Println(appName()) //呼べる
	//AppName // 定数のスコープが関数内なので呼べない
	//Version // 変数も

	fmt.Println(Do("DoDoDO!"))
}
