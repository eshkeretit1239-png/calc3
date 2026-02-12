package main

import (
	"fmt"
	"strings"
)
		
type maps = map[string]float64

func main(){
	valueValut := maps{
		"convertUsdtToEur" : 0.85,
		"convertUsdtToRub" : 75.63,
		"convertRubinEur" : 75.63/0.85,
	}
var firstV string
var quantity int 
var secondV string
for{
firstV, quantity, secondV = input (firstV, quantity, secondV)
if firstV == "rub" && secondV == "usd"{
	var usdAfterConvertrub float64 = float64(quantity)*(1/valueValut["convertUsdtToRub"])
	fmt.Printf("%.2f usd\n", usdAfterConvertrub)
} else if firstV == "rub" && secondV == "eur"{
	var rubAfterConvertrub float64 = float64(quantity)*(1/valueValut["convertRubinEur"])
	fmt.Printf("%.2f eur\n", rubAfterConvertrub)
} else if firstV == "usd" && secondV == "rub"{
	var usdAfterConvertrub float64 = float64(quantity)*(valueValut["convertUsdtToRub"])
	fmt.Printf("%.2f rub\n", usdAfterConvertrub)
} else if firstV == "usd" && secondV == "eur"{
	var usdAfterConverteur float64 = float64(quantity)*(valueValut["convertUsdtToEur"])
	fmt.Printf("%.2f eur\n", usdAfterConverteur)
} else if firstV == "eur" && secondV == "rub"{
	var eurAfterConvertrub float64 = float64(quantity)*(valueValut["convertRubinEur"])
	fmt.Printf("%.2f rub\n", eurAfterConvertrub)
}  else if firstV == "eur" && secondV == "usd"{
	var eurAfterConvertusd float64 = float64(quantity)*(1/valueValut["convertUsdtToEur"])
	fmt.Printf("%.2f usd\n", eurAfterConvertusd)
}  

fmt.Println("Хотите ли продолжит?")
var answer bool = cicle()
if answer == false{
	break
}


}
}


func input (firstV string, quantity int, secondV string) (string, int, string){
fmt.Println("Введите исходную валюту")
for{
	fmt.Scanln(&firstV)
	if strings.ToLower(firstV) == "usd" || strings.ToLower(firstV) == "eur" || strings.ToLower(firstV) == "rub"{
		break
		}else{
fmt.Println("ошипка")
}
}
fmt.Println("Количество")
for { 
	fmt.Scanln(&quantity)
if quantity >= 0{
break
}else{
fmt.Println("Ведите число больше ноля")
}
}
fmt.Println("целевую валюту")
for{ 
	fmt.Scanln(&secondV)
if strings.ToLower(secondV) == "usd" || strings.ToLower(secondV) == "rub" || strings.ToLower(secondV) == "eur"{
break
}else{
fmt.Printf("ошипка")
}
}
return firstV, quantity, secondV
}


func cicle ()(bool){
var answer string	
	for{
fmt.Scanln(&answer) 
	if answer == "yes"{
	return true
} else if answer == "no"{
	return false
} else {
	fmt.Println("Введите yes или no")
	}
	}
}