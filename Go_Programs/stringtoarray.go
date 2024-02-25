package main

func main() {
	str := "Abhijeet"
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		println(char)
	}
}
