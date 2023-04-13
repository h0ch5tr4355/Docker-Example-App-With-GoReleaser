package helloworld

import (
	"fmt"
	"os"
)

func OpenMyFile() {

	data, err := os.ReadFile("data/mydata.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content of mydata.txt: ", string(data))
}
