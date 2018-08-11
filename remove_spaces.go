/*
created by Rodrigo Stuchi for demonstration purposes
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var re = regexp.MustCompile(`[\s\p{Zs}]+`)
	v := " a  b  c "
	s := re.ReplaceAllString(v, " ")
	s = strings.Trim(s, " ")

	av := make([]uint8, len(v))
	for i, k := range v {
		av[i] = uint8(k)
	}

	fmt.Printf( "original value : \"%v\"\n", v)
	fmt.Printf( "changed value  : \"%v\"\n\n", s)

	fmt.Println("---before---")
	fmt.Println("spaces details : [.  a  .  .  b  .     .     c  . ]")
	fmt.Printf( "char codes     : %v\n\n", av)

	fmt.Println("---after---")

	an := make([]uint8, len(s))
	for i, k := range s {
		an[i] = uint8(k)
	}
	fmt.Println("[a  .  b  .  c]")
	fmt.Printf("%v\n", an)
}
