package main

import "fmt"

type person struct{
	name string
	age int
}

func main()  {
	fmt.Println(person{"yanhao",20})
	fmt.Println(person{name:"liju",age:29})

	fmt.Println(person{name:"yanmin"})

	fmt.Println(&person{name:"7lk",age:222})

	s:=person{name:"ces",age:33}
	fmt.Println(s.name)

	sp :=&s

	fmt.Println(sp.age)

	sp.age=3333

	fmt.Println(sp.age)
}