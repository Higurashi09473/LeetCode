package main

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func scoreOfString(s string) int {
    var num int
    mas := []byte(s)
    for i := 0; i < len(mas)-1; i++{
        num += Abs(int(mas[i])-int(mas[i+1]))
    }
    return num
}