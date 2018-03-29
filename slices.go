package main

import (
	"fmt"
)

func main()  {
	s :=make([]string,3)
	fmt.Println("emp",s)

	s[0] ="a"
	s[1]="b"
	s[2]="c"

	fmt.Println("set:",s)
	fmt.Println("get",s)

	s = append(s,"d")
	s=append(s,"e","f")

	fmt.Println("apend:",s)


	c :=make([]string,len(s))

	copy(c,s)

	fmt.Println("cpy",c)


	l:=s[2:5]
	fmt.Println("sl2",l)

	towd := make([][]int,3)
	for i:=0;i<3;i++ {
		innerLen :=i+1
		towd[i] = make([]int,innerLen)
		for j :=0;j<innerLen ;j++  {
			towd[i][j] = i +j
		}

	}
	fmt.Println("2d:",towd)
}
