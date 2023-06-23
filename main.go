package main

import (
	"backend/logic"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/repeale/fp-go"
)


func PerformSearch(path string, recursive bool, search string, mode string) map[string][]os.FileInfo {
	result := make(map[string][]os.FileInfo)

	directories := []string{path}

	if (recursive) {
		files, err := logic.GetAllSubFolders(path)
		if err != nil {panic(err)}
		directories = append(directories, (fp.Map(func(x logic.Directory) string {return x.Position + "/" + x.Name})(files))...)
	}

	for _, dir := range directories {
		files, err := logic.GetFilesInPath(dir)
		if err != nil {panic(err)}
		result[dir] = *files
	}

	for key, value := range result {
		result[key] = logic.SearchByName(value, search)
		if mode == "name" {result[key] = logic.SortByName(value)}
		if mode == "size" {result[key] = logic.SortBySize(value)}
		if mode == "date" {result[key] = logic.SortByDate(value)}
	}

	return result
}

type File struct {
	Name string
	Dir string
	Date string
	IsFolder bool
	Size int64
}


func main() {
	path := flag.String("p", ".", "Directory to search")
	subS := flag.Bool("r", false, "Search trough child folders")
	name := flag.String("s", "", "Search by name")
	mode := flag.String("m", "", "Search mode: name, size, date")
	flag.Parse()

	search := PerformSearch(*path, *subS, *name, *mode)

	results := make([]File, 0)

	for key,value := range search {
		for _, file := range value {
			file_new := File {
				Name: file.Name(),
				Dir: key,
				Date: file.ModTime().Format("2006-01-02"),
				IsFolder: file.IsDir(),
				Size: file.Size(),
			}
			results = append(results, file_new)
		}
	}

	response, err := json.Marshal(results)
	if err != nil {panic(err)}
	fmt.Println(string(response))

}
