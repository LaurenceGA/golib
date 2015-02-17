/*
Package colours Provides a simple way to manipulate colours
*/
package colours

//Colour defines a single rgba colours
type Colour struct {
	R, G, B, A float32
}

var (
	//Black colour (0, 0, 0, 1)
	Black = Colour{0, 0, 0, 1}
	//White colour (1, 1, 1, 1)
	White = Colour{1, 1, 1, 1}
	//Grey colour (0.5, 0.5, 0.5, 1)
	Grey = Colour{0.5, 0.5, 0.5, 1}

	//Red colour (1, 0, 0, 1)
	Red = Colour{1, 0, 0, 1}
	//Green colour (0, 1, 0, 1)
	Green = Colour{0, 1, 0, 1}
	//Blue colour (0, 0, 1, 1)
	Blue = Colour{0, 0, 1, 1}
)
