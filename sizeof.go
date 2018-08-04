
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ageOfUniverse uint64
	reservedBytes := unsafe.Sizeof(ageOfUniverse)

	fmt.Printf("Number of bytes reserved for `ageOfUniverse` variable is %d bytes.\n",
		reservedBytes)
}
