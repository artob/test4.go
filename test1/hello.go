package test1

import "fmt"
import "github.com/near/borsh-go"

func Add(x, y int) int {
	res, err := borsh.Serialize(x)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	return x + y
}
