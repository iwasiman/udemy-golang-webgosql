package main // section16

import (
	"log"
	"os"
)

func main() {
	//log.fatal
	_, err := os.Open("naiyo.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

/*
2022/04/08 11:32:43 open naiyo.txt: The system cannot find the file specified.
exit status 1

*/
