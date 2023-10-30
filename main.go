package main

import(
    "fmt"
)

func main() {
    /*fmt.Println(name)
    fmt.Println(&name)*/ // name değişkeninin ramdeki adresi

    //x := 22

    /*fmt.Println(x)
    fmt.Println(&x)

    fmt.Println()

    fmt.Printf("%T, %v\n", x, x)
    fmt.Printf("%T, %v\n", &x, &x)

    var y *int = &x 
    fmt.Println(y)
    fmt.Printf("%T, %v\n", y, y)*/

    /*z := &name
    fmt.Printf("%T, %v\n", z, z)*/

    /*fmt.Println(x)
    fmt.Println(&x)
    fmt.Println(*(&x))
    fmt.Println(&(*(&x)))*/

    /*x1 := 10
    x2 := x1
    fmt.Println(x1, x2)
    x1 = 5
    fmt.Println(x1, x2)*/

    /*x1 := 10
    x2 := &x1
    fmt.Println(x1, *x2)
    x1 = 5
    fmt.Println(x1, *x2)
    *x2 = 3
    fmt.Println(x1, *x2)*/

    /*x1 := 10
    x3 := &x1
    *x3 = 5
    // fmt.Println(x1, *x3)*/

    // x1 := [4]int{1, 10, 100, 1000} // array pass by value 
    // x1 := []int{1, 10, 100, 1000} // slice pass by reference
    // x2 := x1

    // fmt.Println(x1, x2)
    
    // x2[0] = 3
    // fmt.Println(x2)
    // fmt.Println(x1)

    // x := 5
    // fmt.Println(x)
    // doubleIt(x)
    // fmt.Println(x)
  
    /* myArr := [3]int{1, 10, 100}

    fmt.Println(myArr)

    doubleIt(myArr)

    fmt.Println(myArr) */

    /*x := 5
    fmt.Println(x)
    doubleIt(&x)
    fmt.Println(x) */

    // UYGULAMA
    /* type myType int
    var x myType 
    x = 10
    fmt.Printf("%T, %v\n", x, x) */

    // UYGULAMA
    /* x := 10
    y := &x
    *y = 5
    fmt.Println(x) */

    //UYGULAMA
    r1 := Rectangle{3, 8}
    fmt.Println(r1.area())
    fmt.Println(r1.circumference()) 
    
    interfaceFunc(r1)

    /* families := showFamily()
    for i := 0; i < len(families); i++ {
        fmt.Printf("%v, %T\n", families[i], families[i])
    } */



}

func interfaceFunc (i shape) {
    fmt.Println(i) 
    fmt.Println(i.area())
    fmt.Println(i.circumference())
    fmt.Printf("%T", i)
    fmt.Println()
}

type shape interface {
    area() int
    circumference() int
}

/* func showFamily() []family {
    family1 := family{
        name : "Yusuf",
        age : 23,
        isMarried: false,
    }

    family2 := family {
        name : "Tuğba",
        age : 29,
    }

    family3 := family {
        "Satı",
        45,
        true,
    }

    var family4 family
    family4.name = "Lokman"
    family4.age = 53
    family4.isMarried = true

    return []family {family1, family2, family3, family4}
}

type family struct {
    name string
    age int
    isMarried bool
} */

type Rectangle struct {
    a int
    b int
}

func (r Rectangle) display() {
    fmt.Println("Kenar uzunlukları: ", r.a, r.b, " olan dikdörtgen")
}

func (r Rectangle) area() int {
    return r.a*r.b
}

func (r Rectangle) circumference() int {
    return r.a * 2 + r.b * 2 
} 

/* func doubleIt(num *int) {
    fmt.Println(num)
    *num *= 2
    fmt.Println(*num)
} */