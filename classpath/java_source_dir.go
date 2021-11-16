package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type javaSourceDir struct {
	p string
}

func (self *javaSourceDir) path() string {
	return self.p
}

func (self *javaSourceDir) readClass(className string) (has bool, data []byte, err error) {
	return self.readResource(className + ".class")
}

func (self *javaSourceDir) readResource(path string) (has bool, data []byte, err error) {
	absPath, err := filepath.Abs(self.p)
	if err != nil {
		return
	}
	fPath := filepath.Join(absPath, path)
	data, err = ioutil.ReadFile(fPath)
	has = err == nil
	return
}
