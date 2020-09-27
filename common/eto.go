package common

import (
	// 	"math/rand"
	"fmt"
	// "strconv"
)

// Eto select img and background color based on number
func Eto(usrID int) [2]string {
	if usrID < 0 {
		usrID = usrID * -1
	} else if usrID == 0 {
		usrID = 87654321
	}
	colorID := usrID * 87654321         // it must be more than 6 digit (16 hex)
	color := fmt.Sprintf("%x", colorID) //16 hexadecimal
	color = color[len(color)-6 : len(color)]
	etoNum := (usrID % 12) + 1
	var etoImg string
	switch etoNum {
	case 1:
		etoImg = "/img/icon/01_rat.png"
		break
	case 2:
		etoImg = "/img/icon/02_buffalo.png"
		break
	case 3:
		etoImg = "/img/icon/03_tiger.png"
		break
	case 4:
		etoImg = "/img/icon/04_rabbit.png"
		break
	case 5:
		etoImg = "/img/icon/05_dragon.png"
		break
	case 6:
		etoImg = "/img/icon/06_snake.png"
		break
	case 7:
		etoImg = "/img/icon/07_horse.png"
		break
	case 8:
		etoImg = "/img/icon/08_sheep.png"
		break
	case 9:
		etoImg = "/img/icon/09_monkey.png"
		break
	case 10:
		etoImg = "/img/icon/10_hen.png"
		break
	case 11:
		etoImg = "/img/icon/11_dog.png"
		break
	case 12:
		etoImg = "/img/icon/12_pig.png"
		break
	}
	return [2]string{etoImg, color}
}
