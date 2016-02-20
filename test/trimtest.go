package main

import "fmt"
import "strings"

func trimEmpty(path []string) []string{
	ret_s := make([]string,0)
	for i := 0;i < len(path);i++ {
		if len(path[i]) > 0 {
			ret_s = append(ret_s,path[i])
		}
	}
	return ret_s
}

func main() {
	dir := "/usr/local/bin/test"
	segments := trimEmpty(strings.Split(dir,"/"))
	fmt.Printf("slice = %q\n",segments)
}
