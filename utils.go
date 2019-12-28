package main

import (
	"io/ioutil"
	"time"
)

var (
	domainPath = "http://localhost:9999/"
)

type FileInfo struct {
	Name    string
	Path    string
	Size    int64
	ModTime time.Time
}

func fileList() ([]FileInfo, error) {
	files, err := ioutil.ReadDir("media/")
	if err != nil {
		return nil, err
	}
	infos := []FileInfo{}
	for i := range files {
		name := files[i].Name()
		infos = append(infos, FileInfo{
			Name:    name,
			Path:    domainPath + "media/" + name,
			Size:    files[i].Size(),
			ModTime: files[i].ModTime(),
		})
	}
	return infos, nil
}

type PageData struct {
	UploadPath string     // url for upload
	Files      []FileInfo // uploaded files
}

func NewPageData() (*PageData, error) {
	files, err := fileList()
	if err != nil {
		return nil, err
	}
	return &PageData{
		UploadPath: domainPath + "upload",
		Files:      files,
	}, nil
}
