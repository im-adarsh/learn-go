package tempconv

import (


)


type Celsius float64
type Fahrenhiet float64


func CelsiusToFahrenhiet(c Celsius) Fahrenhiet {
	return Fahrenhiet(c*9/5+32)
}

func FahrenhietToCelsius(f Fahrenhiet) Celsius {
	return Celsius((f-32)*5/9)
}