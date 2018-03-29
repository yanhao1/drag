package  main

import "fmt"
import "math"

const s string  ="contant"

func  main()  {
	fmt.Println(s)
	const n  = 5000000

	const d  = 3e20/n

	fmt.Println(d)

	fmt.Println(int(d))

	fmt.Println(math.Sin(n))
}