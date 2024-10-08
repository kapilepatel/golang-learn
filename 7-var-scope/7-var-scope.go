package main

var x int = 40

func main() {
	//here we simply access the variable from package scope
	println("Print x from main func ", x)
	PrintExt()

	//here we update the variable which is in package scope
	x = 41
	println("Print x from main func ", x)
	PrintExt()

	//This is called variable "shadowing"
	x := 42
	println("Print x from main func ", x)
	//exteral print will still refer to x which is in main scope
	PrintExt()

}

func PrintExt() {
	println("Print x from external func ", x, "\n")
}
