package main

import (
	"fmt"
	"repos"
)

func main() {
	grade := repos.GetRepo()
	fmt.Println(grade)
}
