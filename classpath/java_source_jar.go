package classpath

import (
	"archive/zip"
	"io/ioutil"
	"path/filepath"
)

type javaSourceJar struct {
	p string
}

func (self *javaSourceJar) path() string {
	return self.p
}

func (self *javaSourceJar) readClass(className string) (has bool, data []byte, err error) {
	return self.readResource(className + ".class")
}

func (self *javaSourceJar) readResource(path string) (has bool, data []byte, err error) {
	absPath, err := filepath.Abs(self.p)
	if err != nil {
		return
	}
	zipRC, err := zip.OpenReader(absPath)
	if err != nil {
		return
	}
	defer zipRC.Close()
	for _, f := range zipRC.File {
		if f.Name == path {
			has = true
			file, err := f.Open()
			if err != nil {
				return has, nil, err
			}
			data, err = ioutil.ReadAll(file)
			return has, data, err
		}
	}
	return
}
