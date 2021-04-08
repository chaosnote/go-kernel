package file

// @see https://golang.org/pkg/path/filepath/

import (
	"io/ioutil"
	"os"
	"path"
)

//-------------------------------------------------------------------------------------------------

// Exist file_path p
func Exist(p string) bool {
	i, e := os.Stat(p)
	if os.IsNotExist(e) {
		return false
	}
	return !i.IsDir()
}

//-------------------------------------------------------------------------------------------------

// Read ...
func Read(dirPath, fileName string) ([]byte, error) {
	p := path.Join(dirPath, fileName)
	d, e := ioutil.ReadFile(p)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//-------------------------------------------------------------------------------------------------

// Write ...
func Write(dirPath, fileName string, content []byte) error {
	e := os.MkdirAll(dirPath, os.ModePerm)
	if e != nil {
		return e
	}

	p := path.Join(dirPath, fileName)
	return ioutil.WriteFile(p, content, os.ModePerm)
}

// WriteAppend ...
func WriteAppend(dirPath, fileName string, content []byte) error {
	e := os.MkdirAll(dirPath, os.ModePerm)
	if e != nil {
		return e
	}

	p := path.Join(dirPath, fileName)
	_, e = os.Stat(p)

	if e == nil { // 檔案已存在

		f, e := os.OpenFile(p, os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if e != nil {
			return e
		}
		defer f.Close()

		_, e = f.Write(content)
		return e

	} else if os.IsNotExist(e) { // 檔案不存在

		return ioutil.WriteFile(p, content, os.ModePerm)
	}
	return e
}
