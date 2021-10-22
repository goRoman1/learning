package main

/*
import (
	"fmt"
	"strings"
)

func envelopeCompare(){
	var x1 float64
	fmt.Println("Введите ширину первого конверта")
	_, err := fmt.Scanf("%f\n", &x1)
	if err != nil {
		panic(err)
	}
	var y1 float64
	fmt.Println("Введите высоту первого конверта")
	_, err = fmt.Scanf("%f\n", &y1)
	if err != nil {
		panic(err)
	}
	var x2 float64
	fmt.Println("Введите ширину второго конверта")
	_, err = fmt.Scanf("%f\n", &x2)
	if err != nil {
		panic(err)
	}
	var y2 float64
	fmt.Println("Введите высоту второго конверта")
	_, err = fmt.Scanf("%f\n", &y2)
	if err != nil {
		panic(err)
	}
	if x1 > x2{
		if y1 > y2{
			fmt.Printf("Второй конверт со сторонами , %f, %f можно вложить в первый конверт со сторонами %f, %f \n", x2, y2, x1, y1 )
		}else {
			if x1 > y2 {
				if y1 > x2 {
					fmt.Printf("Второй конверт со сторонами , %f, %f можно вложить в первый конверт со сторонами %f, %f \n", x2, y2, x1, y1 )
				}
			} else {
				fmt.Println("Конверты нельзя вложить друг в друга")
			}
		}
	}else {
		if x1 > y2{
			if y1 > x2 {
				fmt.Printf("Пеорвый конверт со сторонами , %f, %f можно вложить в второй конверт со сторонами %f, %f \n", x1, y1, x2, y2)
			}else {
				fmt.Println("Конверты нельзя вложить друг в друга")
			}
		}else{
			fmt.Printf("Пеорвый конверт со сторонами , %f, %f можно вложить в второй конверт со сторонами %f, %f \n", x1, y1, x2, y2)
		}
	}
}

func main() {
	var count int = 0
	for {
		if count == 0 {
			envelopeCompare()
			count++
		}else{
			var ask string
			fmt.Println("Для повторного подсчета введите 'y' или 'yes'")
			fmt.Scanf("%s\n", &ask)
			if strings.ToLower(ask) == "y" || strings.ToLower(ask) == "yes"{
				envelopeCompare()
			}else{
				fmt.Println("Конец выполенния")
				break
			}
		}
	}
}


 */