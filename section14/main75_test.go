package main //section14
import (
	"testing"
	"udemy-golang-webgosql/section14/alib"
)

var Debug bool = false

// 関数名の先頭がTest。引数は固定。
func TestIsONe1(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップさせるよ")
	}
	v := IsOne(1)
	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}

func TestIsONe0(t *testing.T) {
	i := 0
	if Debug {
		t.Skip("スキップさせるよ")
	}
	v := IsOne(1)
	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}

func TestAverage1(t *testing.T) {
	if Debug {
		t.Skip("スキップさせるよ")
	}
	v := alib.Average([]int{2, 4, 6})
	if v != 4 {
		t.Errorf("%v != %v", v, 4)
	}
}

/*
パッケージのディレクトリで単に実行してもうまくいかない。
>go test
go: cannot find main module, but found .git/config in D:\xxx\udemy-golang-webgosql
        to create a module there, run:
        cd .. && go mod init

環境変数{GOPATH}/src/udemy-golang-webgosql としてクローンしなおす。
> cd その場所
> go mod init
go: creating new go.mod: module udemy-golang-webgosql
go: to add module requirements and sums:
        go mod tidy

ルートに go.mod ファイルができる。
---
module udemy-golang-webgosql

go 1.18
---
> cd section14
> go test
PASS
ok      udemy-golang-webgosql/section14 7.223s

任意の場所で実行したい場合
/somewhere/udemy-golang-webgosql としてクローン済みの場合
cd  /somewhere/udemy-golang-webgosql
> go mod init udemy-golang-webgosql //ここでは引数にモジュール名がいるのに注意！
go: creating new go.mod: module udemy-golang-webgosql
go: to add module requirements and sums:
        go mod tidy
ルートに go.mod ファイルができる。中身は同じ。
---
module udemy-golang-webgosql

go 1.18
---
> cd section14
> go test
PASS
ok      udemy-golang-webgosql/section14 5.580s
*/

/*
-v で詳細が分かる。
> pwd
> /somewhereudemy-golang-webgosql/section14
> go test -v
=== RUN   TestIsONe1
--- PASS: TestIsONe1 (0.00s)
=== RUN   TestIsONe0
--- PASS: TestIsONe0 (0.00s)
=== RUN   TestAverage1
    main75_test.go:38: 2 != 4
--- FAIL: TestAverage1 (0.00s)
FAIL
exit status 1
FAIL    udemy-golang-webgosql/section14 8.579s

=== RUN   TestIsONe1
--- PASS: TestIsONe1 (0.00s)
=== RUN   TestIsONe0
--- PASS: TestIsONe0 (0.00s)
=== RUN   TestAverage1
--- PASS: TestAverage1 (0.00s)
PASS
ok      udemy-golang-webgosql/section14 7.348s

パッケージに限ったテストもできる。パッケージは相対パス。
> go test ./alib
?       udemy-golang-webgosql/section14/alib    [no test files] //_test.goがない時

=== RUN   TestAverage1
--- PASS: TestAverage1 (0.00s)
PASS
ok      udemy-golang-webgosql/section14/alib    6.825s

間違った指定の場合
> go test -v alib
package alib is not in GOROOT (D:\xxx\go1.18\src\alib)

全パッケージの全テストファイルを実行する場合
> go test ./...

=== RUN   TestIsONe1
--- PASS: TestIsONe1 (0.00s)
=== RUN   TestIsONe0
--- PASS: TestIsONe0 (0.00s)
=== RUN   TestAverage1
--- PASS: TestAverage1 (0.00s)
PASS
ok      udemy-golang-webgosql/section14 9.838s
=== RUN   TestAverage1
--- PASS: TestAverage1 (0.00s)
PASS
ok      udemy-golang-webgosql/section14/alib    (cached)

カバレッジも出るらしい。(出なかった)

*/
