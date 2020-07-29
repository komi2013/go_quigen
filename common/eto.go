package common

import (
// 	"math/rand"
  "fmt"
  // "strconv"
)

func Eto (usr_id int) [2]string {
  if usr_id < 0 {
    usr_id = usr_id * -1
  } else if usr_id == 0 {
    usr_id = 87654321;
  }
  color_id := usr_id * 87654321  // it must be more than 6 digit (16 hex)
  color := fmt.Sprintf("%x", color_id) //16 hexadecimal
  color = color[len(color) - 6:len(color)]
  eto_num := ( usr_id % 12 ) +1
  var eto_img string
  switch eto_num {
    case 1:
      eto_img = "/img/icon/01_rat.png";
      break;
    case 2:
      eto_img = "/img/icon/02_buffalo.png";
      break;
    case 3:
      eto_img = "/img/icon/03_tiger.png";
      break;
    case 4:
      eto_img = "/img/icon/04_rabbit.png";
      break;
    case 5:
      eto_img = "/img/icon/05_dragon.png";
      break;
    case 6:
      eto_img = "/img/icon/06_snake.png";
      break;
    case 7:
      eto_img = "/img/icon/07_horse.png";
      break;
    case 8:
      eto_img = "/img/icon/08_sheep.png";
      break;
    case 9:
      eto_img = "/img/icon/09_monkey.png";
      break;
    case 10:
      eto_img = "/img/icon/10_hen.png";
      break;
    case 11:
      eto_img = "/img/icon/11_dog.png";
      break;
    case 12:
      eto_img = "/img/icon/12_pig.png";
      break;
  }
  return [2]string {eto_img, color}
}