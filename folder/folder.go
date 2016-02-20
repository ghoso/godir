package folder
//
// folder.go
//
import (
//	"fmt"
	"path/filepath"
	"errors"
//	"encoding/json"
	"strings"
	"os"
)

const (
	FILE = 0
	DIR = 1
)

type File struct {
	IsDir bool `json: "IsDir"`
	Path string `json: "Path"`
	Name string `json: "Name"`
	Children map[string]*File `json: "Children"`
}

// Top directory
var top_directory string

// Eliminate empty string from file path slice.
func trimEmpty(path []string) []string{
	ret_s := make([]string,0)
	for i := 0;i < len(path);i++ {
		if len(path[i]) > 0 {
			ret_s = append(ret_s,path[i])
		}
	}
	return ret_s
}

// Create folder information
func NewFolder(path string) *File{
	_,name := filepath.Split(path)
	if len(top_directory) == 0 {
		top_directory = path
	}
	return &File{Name: name,Path: path, Children: make(map[string]*File), IsDir: true}
}

// Search folder informatin
func (f *File) GetFolder(path string) (error,*File) {
	// Make path absolute
	cpath,_ := filepath.Abs(filepath.Clean(path))
	// Path validation
	finfo,_ := os.Stat(cpath)
	if finfo == nil || finfo.IsDir() == false {
		return errors.New(cpath + " is not a directory."),nil
	}
	// if top dicrectory,return current File object
	if cpath == top_directory {
		return nil,f
	}
	// Path string convart into slices
	segments := trimEmpty(strings.Split(cpath,"/"))
	top_segments := trimEmpty(strings.Split(top_directory,"/"))
	fdir := f
	for i := len(top_segments) ;i < len(segments);i++ {
		dname := segments[i]
		fdir  = fdir.Children[dname]
		if fdir == nil {
			return errors.New("Could not find " + dname + " in "  + cpath),nil
		}
	}
  return nil,fdir
}

// Add folder information
func (f *File) AddFolder(path string) error{
	// Check if same name's directory or file exists
	_,file := filepath.Split(path)
	if f.Children[file] != nil {
			return errors.New(path + " already exists.")
	}		
	// If not exists,create new folder information
	f.Children[file] = NewFolder(path)
	return nil
}

// Delete folder information
func (f *File) DeleteFolder(path string) error{
	// Not yet implemented
	return nil
}

// Add file information
func (f *File) AddFile(path string) error{
	_,file := filepath.Split(path)
	// Check file name duplicate
	if f.Children[file] != nil {
			return errors.New(path + " already exists.")
	}
	// Add file information into directory
	f.Children[file] = &File{Name: file,Path: path, Children: nil, IsDir: false}
	return nil
}

// Delete file information
func (f *File) DeleteFile(path string) error {
	// Not yet implemented
	return nil
}
