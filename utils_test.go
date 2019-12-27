package main

import "testing"

func Test_fileList(t *testing.T) {
	files, err := fileList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(files)
}
