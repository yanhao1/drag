package main

import (
	"errors"
	"fmt"
)

func f1(arg int)(int,error){
	if arg ==42{
		return -1,errors.New("cant work with 42")
	}
	return arg + 3,nil
}



type argEroor struct {
	age int
	prob string
}


func (e *argEroor) error() string{
	return fmt.Sprintf("%d - %s",e.age,e.prob)
}


func f2(arg int)(int , error){
	if arg == 42{
		return -1,&argEroor{arg,"cant work width it"}
	}
	return arg +3,nil

	}

func main()  {
	for _,i := range []int{7,42}{
		if r,e :=f1(i);e !=nil {
			fmt.Println("f1 failed:",e)
		}else {
			fmt.Println("f1 worked:",r)
		}
	}


}