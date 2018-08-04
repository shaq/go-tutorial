package main

import (
	"fmt"
)

// on 32-bit machines, it'll overflow
// on 64-bit machines, everything is fine
// use: uint64 instead.
func main() {
	var ageOfUniverse uint64 = 14e9 // 14e9 = 14 and 9 zeros which is 14 billions
	fmt.Printf("`ageOfUniverse` is %d.\n", ageOfUniverse)
}
