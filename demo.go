// package main

// import "fmt"

// //this is a single line comment

// func main() {

// 	/*this is
// 	a multi line
// 	comment*/
// 	fmt.Println("Hello World")
// }
// package main

// import "fmt"

// func main() {
// 	name := "bharathi"
// 	fmt.Println(name)

// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var greeting string = "Hello world"
// 	fmt.Println(greeting)

// }
// package main

// import (
// 	"fmt"
// )

// func main() {

// 	fmt.Print("hello world!!!!")

// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var name string = "kodekloud"
// 	var user string = "bharathi"

// 	fmt.Print("Welcome to ", name, ", ", user)
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var name string = "kodekloud"
// 	var user string = "bharathi"
//fmt.Print(name, "\n"),fmt.Println(name)the both are printing the same new line.
// 	fmt.Print(name, "\n")
// 	fmt.Print(user)
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var name string = "bharathi"
// 	var i int = 90
// 	fmt.Printf("Hey, %v! you have scored %d/100 in computer", name, i)
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var S, t string = "foo", "bar"
// 	fmt.Println(S)
// 	fmt.Println(t)
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var (
// 		S string = "foo"
// 		i int    = 5
// 	)
// 	fmt.Println(S)
// 	fmt.Println(i)
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	name := "sravani"
// 	name = "bharathi"
// 	fmt.Println(name)
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	city := "londan"
// 	{
// 		country := "uk"
// 		fmt.Println(country)
// 		fmt.Println(city)
// 	}
// 	fmt.Println(city)
// }
//local variable example.
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	name := "bharathi"
// 	fmt.Println(name)

// }
//global variable example.
// package main

// import (
// 	"fmt"
// )

// var name string = "bharathi"

// func main() {
// 	fmt.Println(name)

// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	//var i int
// 	//fmt.Printf("%d", i)
// 	var f1 float64
// 	fmt.Printf("%.2f", f1)
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var name string
// 	fmt.Print("Enter your name: ")
// 	fmt.Scanf("%s", &name)
// 	fmt.Println("Hey there, ", name)
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var name string
// 	var is_muggle bool
// 	fmt.Print("Enter your name & are you a muggle: ")
// 	fmt.Scanf("%s %t", &name, &is_muggle)
// 	fmt.Println(name, is_muggle)

// }
//multipul input
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a string
// 	var b int
// 	fmt.Print("Enter a string and a namber: ")
// 	count, err := fmt.Scanf("%s %d", &a, &b)
// 	fmt.Println("count : ", count)
// 	fmt.Println("error: ", err)
// 	fmt.Println("a: ", a)
// 	fmt.Println("b: ", b)
// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var grades int = 30
// 	var message string = "hello world"
// 	var ischeck bool = true
// 	var amount float32 = 33.65
// 	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
// 	fmt.Printf("variable message = '%v' is of type %T \n", message, message)
// 	fmt.Printf("variable ischeck = '%v' is of type %T \n", ischeck, ischeck)
// 	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)
// }
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
// 	fmt.Printf("Type: %v \n", reflect.TypeOf("bharathi"))
// 	fmt.Printf("Type: %V \n", reflect.TypeOf(12.0))
// 	fmt.Printf("Type: %v \n", reflect.TypeOf(true))

// }
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	var grades int = 30
// 	var message string = "hello world"

// 	fmt.Printf("variable grades = %v is of type %v \n", grades, reflect.TypeOf(grades))
// 	fmt.Printf("variable message = '%v' is of type %v \n", message, reflect.TypeOf(message))
// }
//integer to float
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var i int = 60
// 	var f float64 = float64(i)
// 	fmt.Printf("%.2f\n", f)
// }
//floate to integer
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var f float64 = 45.89
// 	var i int = int(f)
// 	fmt.Printf("%v\n", i)
// }
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var i int = 42
// 	var s string = strconv.Itoa(i) // convert int to string
// 	fmt.Printf("%q", s)

// }
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var s string = "200"
// 	i, err := strconv.Atoi(s)
// 	fmt.Printf("%v, %T \n", i, i)
// 	fmt.Printf("%v, %T", err, err)
// }
// we create a invalid syntex error
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var s string = "200abc"
// 	i, err := strconv.Atoi(s)
// 	fmt.Printf("%v, %T \n", i, i)
// 	fmt.Printf("%v, %T", err, err)
// }
// package main

// import "fmt"

// func main() {
// 	const name = "bharathi"
// 	const is_muggle = false
// 	const age = 20
// 	fmt.Printf("%v: %T \n", name, name)
// 	fmt.Printf("%v: %T \n", is_muggle, is_muggle)
// 	fmt.Printf("%v: %T \n", age, age)

// }
// package main

// import "fmt"

// const PI float64 = 3.14 //gloable constant
// func main() {
// 	var radius float64 = 5.0
// 	var area float64
// 	area = PI * radius * radius
// 	fmt.Println("Area of circle is : ", area)
// }
////////////////operators and control flow
// this is == operator(comparsion operator)
// package main

// import "fmt"

// func main() {
// 	var city string = "kolkata"
// 	var city_2 string = "colkata"
// 	fmt.Println(city == city_2)
// }
//this is != operator
// package main

// import "fmt"

// func main() {
// 	var city string = "kolkata"
// 	var city_2 string = "colkata"
// 	fmt.Println(city != city_2)
// }
// package main
// import "fmt"
// func main(){
// 	var a, b int =5, 10
// 	fmt.Println(a < b)
// }
// package main

// import "fmt"

// func main() {
// 	var a, b int = 10, 10
// 	fmt.Println(a <= b)
// }
// package main

// import "fmt"

// func main() {
// 	var a, b int = 20, 10
// 	fmt.Println(a > b)
// }
// package main

// import "fmt"

// func main() {
// 	var a, b int = 10, 10
// 	fmt.Println(a >= b)
// }
//arithetic operator
// package main

// import "fmt"

// func main() {
// 	var a, b string = " bharathi", "pallapu"
// 	var c, d int = 5, 10
// 	fmt.Println(a + b)
// 	fmt.Println(c + d)
// }
// package main

// import "fmt"

// func main() {
// 	var a, b float64 = 77.00, 44.87
// 	fmt.Printf("%.2f", a-b)
// }
// package main

// import "fmt"

// func main() {
// 	var c, d int = 5, 10
// 	fmt.Println(c * d)
// }
// package main

// import "fmt"

// func main() {
// 	var c, d int = 12, 10
// 	fmt.Println(c / d)
// }

// package main

// import "fmt"

// func main() {
// 	var c, d int = 18, 10
// 	fmt.Println(c % d)
// }
// package main

// import "fmt"

// func main() {
// 	var i int = 1
// 	i++
// 	fmt.Println(i)
// }
// package main

// import "fmt"

// func main() {
// 	var i int = 1
// 	i--
// 	fmt.Println(i)
// }
// package main

// import "fmt"

// func main() {
// 	var X int = 10
// 	fmt.Println((X < 100) && (X < 200))
// 	fmt.Println((X < 300) && (X < 0))
// }
// package main

// import "fmt"

// func main() {
// 	var X int = 10
// 	fmt.Println((X < 0) || (X < 200))
// 	fmt.Println((X < 0) && (X > 200))
// }
// package main

// import "fmt"

// func main() {
// 	var x, y int = 10, 20
// 	fmt.Println(!(x > y))
// 	fmt.Println(!(true))
// 	fmt.Println(!(false))

// }
// package main

// import "fmt"

// func main() {
// 	var x int = 10
// 	var y int
// 	y = x
// 	fmt.Println(y)
// }
// package main

// import "fmt"

// func main() {
// 	var x, y int = 10, 20
// 	x += y
// 	fmt.Println(x)
// }
// package main

// import "fmt"

// func main() {
// 	var x, y int = 10, 20
// 	x -= y
// 	fmt.Println(x)
// }
// package main

// import "fmt"

// func main() {
// 	var x, y int = 10, 20
// 	x *= y
// 	fmt.Println(x)
// }
// package main

// import "fmt"

// func main() {
// 	var x, y int = 200, 20
// 	x /= y
// 	fmt.Println(x)
// }
// package main

// import "fmt"

// func main() {
// 	var x, y int = 210, 20
// 	x %= y
// 	fmt.Println(x)
// }
//  package main

//  import "fmt"

//  func main() {
//  	var x, y int = 12, 25
//  	z := x & y
//  	fmt.Println(z)
//  }
// package main

// import "fmt"

// func main() {
// 	var x, y int = 12, 25
// 	z := x | y
// 	fmt.Println(z)
// }
// package main

// import "fmt"

// func main() {
// 	var x, y int = 12, 25
// 	z := x ^ y
// 	fmt.Println(z)
// }
// package main

// import "fmt"

// func main() {
// 	var x int = 212
// 	z := x << 1
// 	fmt.Println(z)
// }
// package main

// import "fmt"

// func main() {
// 	var x int = 212
// 	z := x >> 2
// 	fmt.Println(z)
// }
// package main

// import "fmt"

// func main() {
// 	var a string = "happy"
// 	if a == "happy" {
// 		fmt.Println(a)
// 	}
// }
// package main

// import "fmt"

// func main() {
// 	var country string = "andra pradesh"
// 	if country == "india" {
// 		fmt.Println("india is a country")
// 	} else {
// 		fmt.Println("andra pradesh is not a country")
// 	}
// }
// package main

// import "fmt"

// func main() {
// 	country := "andra pradesh"
// 	if country == "india" {
// 		fmt.Println("india is a country")
// 	} else if country == "Ravachoty" {
// 		fmt.Println("andra pradesh is not a country")
// 	} else {
// 		fmt.Println("no appetite")
// 	}
// }
// package main

// import "fmt"

// func main() {
// 	var a, b string = "foo", "bar"
// 	if a+b == "foo" {
// 		fmt.Println("foo")
// 	} else if a+b == "bar" {
// 		fmt.Println("bar")
// 	} else if a+b == "foobar" {
// 		fmt.Println("foobar")
// 	} else {
// 		fmt.Println("None matched")
// 	}
// 	fmt.Println("thank you!")
// }
// package main

// import "fmt"

// func main() {
// 	var a, b string = "foo", "bar"
// 	if a+b == "foo" {
// 		fmt.Println("foo")
// 	} else if a+b == "foobar" {
// 		fmt.Println("bar")
// 	} else if a+b == "foobar" {
// 		fmt.Println("foobar")
// 	} else {
// 		fmt.Println("None matched")
// 	}
// 	fmt.Println("thank you!")

// }
// package main

// import "fmt"

// func main() {
// 	var i int = 100
// 	switch i {
// 	case 10:
// 		fmt.Println("i is 10")
// 	case 100, 200:
// 		fmt.Println("i is either 100 or200")
// 	default:
// 		fmt.Println("i is neither 0, 100 or 200")

// 	}
// }
// package main

// import "fmt"

// func main() {
// 	var i int = 10
// 	switch i {
// 	case -5:
// 		fmt.Println("-5")
// 	case 10:
// 		fmt.Println("10")
// 		fallthrough
// 	case 20:
// 		fmt.Println("20")
// 		fallthrough
// 	default:
// 		fmt.Println("default")

// 	}
// }
// package main

// import "fmt"

// func main() {
// 	var a, b int = 10, 20
// 	switch {
// 	case a+b == 30:
// 		fmt.Println("is equal to 30")
// 	case a+b <= 30:
// 		fmt.Println("lessthan or equal to 30")

// 	default:
// 		fmt.Println("greater than 30")

// 	}
// }
// package main

// import "fmt"

// func main() {
// 	var a, b = 100, 5
// 	switch {
// 	case a/b == 10:
// 		fmt.Println("10")
// 	case a/b == 20:
// 		fmt.Println("20")
// 	case a/b == 10:
// 		fmt.Println("30")
// 	default:
// 		fmt.Println("default")
// 	}

// }
// package main

// import "fmt"

// func main() {
// 	day := "sunday"
// 	switch day {
// 	case "monday":
// 		fmt.Println("monday")
// 	case "tuesday":
// 		fmt.Println("tuesday")
// 	case "wednesday":
// 		fmt.Println("wednesday")
// 	case "thursday":
// 		fmt.Println("thursday")
// 	case "friday":
// 		fmt.Println("friday")
// 	case "saturday", "sunday":
// 		fmt.Println("weekend")
// 	default:
// 		fmt.Println("default")
// 	}

// }
// package main

// import "fmt"

// func main() {
// 	day := "wednesday"
// 	switch day {
// 	case "monday":
// 		fmt.Println("monday")
// 	case "tuesday":
// 		fmt.Println("tuesday")
// 	case "wednesday":
// 		fmt.Println("wednesday")
// 		fallthrough
// 	case "thursday":
// 		fmt.Println("thursday")
// 		fallthrough
// 	case "friday":
// 		fmt.Println("friday")
// 	case "saturday", "sunday":
// 		fmt.Println("weekend")
// 	default:
// 		fmt.Println("default")
// 	}

// }
// package main

// import "fmt"

// func main() {
// 	var i, j = 10, 50

// 	switch {
// 	case i+j == 60:
// 		fmt.Println("equal to 60")
// 	case i+j <= 60:
// 		fmt.Println("less than or equal to 60")
// 		fallthrough
// 	default:
// 		fmt.Println("greater than 60")
// 	}

// }
// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i * i)
// 	}
// }
// package main

// import "fmt"

// func main() {
// 	i := 1
// 	for i <= 5 {
// 		fmt.Println(i * i)
// 		i += 1
// 	}
// }
// package main

// import "fmt"

// func main() {
// 	sum := 0
// 	for {
// 		sum++ //repeated forword
// 	}
// 	fmt.Println(sum) //never reached
// }
// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		if i == 3 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}

// }
// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		if i == 3 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}

// }
// package main

// import "fmt"

// func main() {
// 	var grades [5]int
// 	fmt.Println(grades)
// 	var fruits [3]string
// 	fmt.Println(fruits)
// }
// package main

// import "fmt"

// func main() {
// 	var fruits [2]string = [2]string{"apple", "orange"}
// 	fmt.Println(fruits)

// 	marks := [3]int{20, 30, 40}
// 	fmt.Println(marks)

// 	names := [...]string{"bharathi", "sravani", "priya"}
// 	fmt.Println(names)
// // }
// package main

// import "fmt"

// func main() {
// 	var fruits [2]string = [2]string{"apple", "orange"}
// 	fmt.Println(len(fruits))
// }

// package main

// import "fmt"

// func main() {
// 	var grades [5]int = [5]int{50, 60, 70, 80, 90}
// 	fmt.Println(grades)
// 	grades[1] = 100
// 	fmt.Println(grades)
// }
// package main

// import "fmt"

// func main() {
// 	var grades [5]int = [5]int{50, 60, 70, 80, 90}
// 	for i := 0; i < len(grades); i++ {
// 		fmt.Println(grades[i])
// 	}
// }
// package main

// import "fmt"

// func main() {
// 	var grades [5]int = [5]int{50, 60, 70, 80, 90}
// 	for index, element := range grades {
// 		fmt.Println(index, "=>", element)
// 	}
// }
//in thes in the place of 1 is 64
// package main

// import "fmt"

// func main() {
// 	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
// 	fmt.Println(arr[2][1])
// }
// package main

// import "fmt"

// func main() {
// 	slice := []int{10, 20, 30}
// 	fmt.Println(slice)
// }
// package main

// import "fmt"

// func main() {
// 	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
// 	slice := arr[1:8]
// 	fmt.Println(slice)
// }
