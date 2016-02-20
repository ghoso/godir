package main
//
// godir.go
//
// Show directory information in JSON format
// Usage: godir [directory]
import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
	"./folder"
)

func dir2json(dir string) (error,[]byte) {
	dir_clean,_ := filepath.Abs(filepath.Clean(dir))
	d := folder.NewFolder(dir_clean)
	err :=filepath.Walk(dir_clean,
		func(path string, info os.FileInfo, err error) error {
			pdir,_ := filepath.Split(path)
			// Proceed directory
			if info.IsDir(){
				// Ignore if top directory
				if dir_clean == path {
					return nil
				// Proceed child directory
				} else {
					err,cd := d.GetFolder(pdir)
					if err != nil {
						return err
					}
					cd.AddFolder(path)
				}
			// Proceed file
			}else{
				err,cd := d.GetFolder(pdir)
				if err != nil {
					return err
				}
				cd.AddFile(path)
			}
			return nil
		})
	if err !=nil {
		return err,nil
	} else {
		json_data,_ := json.Marshal(d)
		return nil,json_data
	}
}

func main(){
	// directory := "./tmp"
	// Argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: godir [directory]")
		os.Exit(1)
	} else {
		directory = os.Args[1]
	}
	
	err,dir_json := dir2json(directory)
	if err != nil {
		fmt.Printf("godir Error: %s\n ",err)
		os.Exit(1)
	}
	fmt.Printf("%s\n",dir_json)
	os.Exit(0)
}
