package std

import "log"

func CheckErr(err error) {
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PrintErr(err error) {
	printErr(err)
}

func printErr(err error) {
	if err != nil {
		log.Printf("Err:%v", err)
	}
}
