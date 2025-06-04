// package main
// import (
// 	"fmt"
// 	"image"
// )
// func main(){

// }

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	// "os"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// fmt.Println(m.Bounds())
	// fmt.Println(m.At(0, 0).RGBA())

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			m.Set(x, y, color.RGBA{73,24 , 108, 255})
		}
	}
	fmt.Println(m)

	file, err := os.Create("output.png")
	if err != nil {
		panic("kuch gadbad hogyi")
	}

	defer file.Close()

	fmt.Println(file)
	png.Encode(file, m)

	// os.Create("abc.png")
}
