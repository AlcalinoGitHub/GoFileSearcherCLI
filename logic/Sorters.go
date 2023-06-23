package logic

import (
	"os"
	"sort"
	"strings"
)

func SortByName(files []os.FileInfo) []os.FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	return files
}

func SortBySize(files []os.FileInfo) []os.FileInfo {
	sort.Slice(files, func(i, j int) bool { 
		return files[i].Size() < files[j].Size()
	})
	return files
}

func SortByDate(files []os.FileInfo) []os.FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Before(files[j].ModTime())
	})
	return files
}

func SearchByName(files []os.FileInfo, search string) []os.FileInfo {
	var results []os.FileInfo

	for _, file := range files {
		if strings.Contains(file.Name(), search) {results = append(results, file)}
	}
	return results
}