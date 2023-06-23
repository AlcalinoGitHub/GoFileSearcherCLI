package logic

import (
	"os"
	"fmt"
	"github.com/repeale/fp-go"
)

type Directory struct {
	Name string
	Position string
}

func GetAllSubFolders(path string)([]Directory, error) {
	all, err := GetFilesInPath(path)
	if err != nil {return []Directory{}, err}

	files := fp.Filter(func(x os.FileInfo) bool {
		return x.IsDir()
	})(*all)
	
	names := (fp.Map(func(x os.FileInfo) string {return x.Name()})(files))

	folders := fp.Map(func(x string) Directory  {return Directory{Name: x, Position: path}})(names)

	for _, dir := range folders {
		newPath := fmt.Sprintf("%s/%s",path,dir.Name)
		newFolders, err := GetAllSubFolders(newPath)
		if err != nil {return folders, err}
		folders = append(folders, newFolders...)
	}

	return folders, nil
}