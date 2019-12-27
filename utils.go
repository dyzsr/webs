package main

import (
	"io/ioutil"
	"time"
)

type fileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
}

func fileList() ([]fileInfo, error) {
	files, err := ioutil.ReadDir("media/")
	if err != nil {
		return nil, err
	}
	infos := []fileInfo{}
	for i := range files {
		infos = append(infos, fileInfo{
			Name:    files[i].Name(),
			Size:    files[i].Size(),
			ModTime: files[i].ModTime(),
		})
	}
	return infos, nil
}
