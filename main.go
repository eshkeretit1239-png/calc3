package main

import (
	"fmt"
	"strings"
)
		
type maps = map[string]float64

func main(){
	valueValut := maps{
		"eur" : 0.85,
		"rub" : 75.63,
		"usd" : 75.63/0.85,
	}
var firstV string
var quantity float64
var secondV string
for{
firstV, quantity, secondV = input()
result := quantity/ valueValut[firstV] * valueValut[secondV]
fmt.Printf("%.2f , %s\n", result, secondV)

fmt.Println("Хотите ли продолжит?")
var answer bool = cicle()
if answer == false{
	break
}


}
}


func input () (string, float64, string){
fmt.Println("Введите исходную валюту")
var from, to string
var amount float64
for{
	fmt.Scanln(&from)
	if strings.ToLower(from) == "usd" || strings.ToLower(from) == "eur" || strings.ToLower(from) == "rub"{
		break
		}else{
fmt.Println("ошипка")
}
}
fmt.Println("Количество")
for { 
	fmt.Scanln(&amount)
if amount >= 0{
break
}else{
fmt.Println("Ведите число больше ноля")
}
}
fmt.Println("целевую валюту")
for{ 
	fmt.Scanln(&to)
if strings.ToLower(to) == "usd" || strings.ToLower(to) == "rub" || strings.ToLower(to) == "eur"{
break
}else{
fmt.Printf("ошипка")
}
}
return from, amount, to
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