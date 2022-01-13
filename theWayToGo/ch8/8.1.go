package main

import "fmt"

func main() {
	date := map[string]string{
		"Sun": "星期天",
		"Mon": "星期一",
		"Tue": "星期二",
		"Wed": "星期三",
		"Thu": "星期四",
		"Fri": "星期五",
		"Sta": "星期六",
	}
	v, d := date["Tue"]

	if d {
		fmt.Println(v)
	} else {
		fmt.Println("not exists Tue")
	}

	v, d = date["Hol"]
	if d {
		fmt.Println(v)
	} else {
		fmt.Println("not exists Hol")
	}

	fmt.Println(date)
}
