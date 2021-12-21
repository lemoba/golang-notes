package main

var a = "G"

func main() {
	n()
	m()
	n()
}

// GOG
func n() {
	print(a)
}

func m() {
	a := "O" //local
	print(a)
}
