package logic

import "os"

func GetFilesInPath(path string) (*[]os.FileInfo, error) {
	dir, err := os.Open(path)
	if err != nil {;return nil, err}
	files, err := dir.Readdir(0)
	if err != nil {return nil, err}
	return &files, nil
}
