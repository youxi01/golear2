package main

var a,b int = 1,2

var(
	c,d string = "ccccc", "dddddd"
	e int = 3
	f bool = true
)
var g int
var h,i = 123,"striiiii"
var l string
var m bool

func main(){
	j,k := 123,false
	println(a, b, c, d, e, f, g, h, i, j, k, l, m)
	&a
}