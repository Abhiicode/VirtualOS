// //print primes
// package main

// import (
// 	"fmt"
// 	"math"
// )
// func linearSearch(dataList []int, key int)bool{
// 	for _,item:= range dataList{
// 		if item == key {
// 			return true
// 		}
// 	}
// 	return false
// }
// func printBar(arr []int){
// 	var max int = 0
// 	for i:=0 ; i<=len(arr) ; i++{
// 		if arr[i]>max{
// 			max = arr[i]
// 		}
// 	}

// 	for i:=max ; i>=0; i--{
// 		var str string = ""
// 		for j:=0 ; j<len(arr) ; j++{
// 			if(arr[j]<=max){
// 				str = str+" \t"
// 			}else{
// 				str = str+"*\t"
// 			}
// 		}
// 		fmt.Println(str)
// 	}
// }
// func printPrimes(num1, num2 int){
// 	if num1<2 || num2<2 {
// 		fmt.Println("Number must be greater than 2")
// 		return
// 	}
// 	for num1<=num2{
// 		isPrime:= true

// 		for i:=2 ; i<=int(math.Sqrt(float64(num1))) ; i++{
// 			if num1%i==0{
// 				isPrime = false;
// 				break;
// 			}
// 		}
// 		if isPrime{
// 			fmt.Printf("%d ",num1)
// 		}
// 		num1++
// 	}
// 	fmt.Println()

// }
// func main(){
// //  printPrimes(5,15)
// arr:= []int {5,4,3,1,0}
// fmt.Print(linearSearch(arr,3))
// }