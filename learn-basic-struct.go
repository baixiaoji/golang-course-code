package main

import "fmt"
import "strconv"

type Product struct {
	id int
	name string
	price float64
}

func main()  {
	//fmt.Printf("hello")

	ageInput, err := strconv.ParseInt("23eee", 0, 64)
	if err != nil {
		fmt.Println("input error")
	} else {
		fmt.Println(ageInput)
	}
}

//func main()  {
//	// 确定长度的 array
//	prices := [2]float64{1.22, 3.44}
//	fmt.Println(prices)
//
//	fmt.Println(prices[0])
//	prices[1] = 4.33
//	fmt.Println(prices)
//
//	// 不确定
//	names := []string{"baixiaoji", "wangyisheng"}
//
//	fmt.Println(names[1])
//	foreignName := []string{"Kevin", "A Jiea"}
//	names = append(names, foreignName...)
//	fmt.Println(names)
//
//	products := []Product {
//		{1, "life of the love", 2.33},
//		{1, "life of the love", 2.44},
//	}
//
//	fmt.Println(products[1])
//
//
//	priceList := [4]float64 {12.11, 21.22, 32.33, 44.33}
//
//	newPriceList := priceList[1:]
//
//	fmt.Println(newPriceList)
//
//	fmt.Println(len(newPriceList), cap(newPriceList))
//
//	anotherPriceList := newPriceList[1:]
//	anotherPriceList[0] = 0.9
//	anotherPriceList = anotherPriceList[0:1]
//	fmt.Println(priceList)
//	fmt.Println(newPriceList)
//	fmt.Println(anotherPriceList)
//
//	fmt.Println(len(anotherPriceList), cap(anotherPriceList))
//
//}