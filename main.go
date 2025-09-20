package main

import "fmt"

func main() {
	Read(NewDarwingPbCopy())
}

func Read(pbcopy Copy) {
	fmt.Println(pbcopy.ReadAll())
}
