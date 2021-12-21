package basic

import "fmt"

func typeConversion() {
	var x int8
	var y int16
	y = 129
	x = int8(y)
	fmt.Print(x, y)

	//128 -> -128
	//129 ->
	//256 -> 0
	//127 -> 127

}
