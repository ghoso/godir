package main

import (
	"fmt"
)

func main(){
	s := make([]string,0)
	str1 := "One"
	str2 := "Two"
	str3 := "Three"
	s = append(s,str1)
	s = append(s,str2)
	s = append(s,str3)
	fmt.Printf("segments = %q\n",s)
	fmt.Printf("last segment = %s\n",s[len(s)-1])
}
