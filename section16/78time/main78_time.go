package main // section16

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("78.time 時間の生成")
	// 現在の時間
	t := time.Now()
	fmt.Println(t)

	// 時刻を指定した時間
	t2 := time.Date(2022, 4, 9, 15, 00, 34, 0, time.Local)
	fmt.Println("t2", t2)
	fmt.Println("t2.Year()", t2.Year())
	fmt.Println("t2.yearDay()", t2.YearDay())
	fmt.Println("t2.Month()", t2.Month())
	fmt.Println("t2.Weekday()", t2.Weekday())
	fmt.Println("t2.Day()", t2.Day())
	fmt.Println("t2.Hour()", t2.Hour())
	fmt.Println("t2.Minute()", t2.Minute())
	fmt.Println("t2.Second()", t2.Second())
	fmt.Println("t2.nanosecond()", t2.Nanosecond())
	//なぜかミリ秒はない。int64(time.Millisecond)で割るとよいそうな。
	fmt.Println(t2.Zone())
	/*
	   2022-04-08 15:05:38.7058625 +0900 JST m=+0.032084901
	   t2 2022-04-09 15:00:34 +0900 JST
	   t2.Year() 2022
	   t2.yearDay() 99
	   t2.Month() April
	   t2.Weekday() Saturday
	   t2.Day() 9
	   t2.Hour() 15
	   t2.Minute() 0
	   t2.Second() 34
	   t2.nanosecond() 0
	   JST 32400
	*/

	fmt.Println("時間の間隔を出力する---- ")
	// 知time.Duration型を返す。
	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)
	/*
		1h0m0s
		time.Duration
		1m0s
		1s
		1ms
		1µs
		1ns
	*/

	fmt.Println("指定の時間先へ---- ")
	// time.Duration型を文字列から生成
	d, _ := time.ParseDuration("2h30m")
	fmt.Println("d", d)
	t3 := time.Now()
	fmt.Println("なう", t3)
	t3 = t3.Add(2*time.Minute + 15*time.Second) //
	fmt.Println("2分15秒後", t3)
	/*
	   d 2h30m0s
	   なう 2022-04-08 15:15:24.2936898 +0900 JST m=+0.670676301
	   2分30秒後 2022-04-08 15:17:39.2936898 +0900 JST m=+135.670676301
	*/

	fmt.Println("時刻の比較---- ")
	// 時刻の差分を取得
	tFuture := time.Date(2022, 7, 24, 0, 0, 0, 0, time.Local)
	tNow := time.Now() // tNowはtFutureより前
	fmt.Println(tNow)
	// t5- t6
	tDiff := tFuture.Sub(tNow)
	fmt.Println("tFuture - tNow:", tDiff) // 2552h34m20.5874541s
	// 時刻を比較する
	fmt.Println(tNow.Before(tFuture)) // true
	fmt.Println(tNow.After(tFuture))  // false
	fmt.Println(tFuture.Before(tNow)) // false
	fmt.Println(tFuture.After(tNow))  // true
	fmt.Println("指定時間のスリープ 5s---- ")
	time.Sleep(5 * time.Second)
	fmt.Println("5秒後におっきしたよ")

}
