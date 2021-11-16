package classpath

import (
	"github.com/yuntao84/Civet/tool"
	"os"
	"path/filepath"
	"strings"
)

type Classpath struct {
	sources []javaSource
	path    string
}

func NewClasspath(path string) (*Classpath, error) {
	cp := &Classpath{path: path}
	err := cp.parseClasspath()
	if err != nil {
		return nil, err
	}
	return cp, nil
}

func (self *Classpath) ReadClass(className string) (data []byte, err error) {
	if has, data, err := self.read(self.sources, className); has {
		return data, err
	}
	return nil, err
}

func (self *Classpath) read(sources []javaSource, className string) (has bool, data []byte, err error) {
	for _, source := range sources {
		has, data, err = source.readClass(className)
		if has {
			return has, data, err
		}
	}
	return
}

func (self *Classpath) parseClasspath() (err error) {
	for _, path := range strings.Split(self.path, string(os.PathListSeparator)) {
		if strings.TrimSpace(path) == "" {
			continue
		}
		if tool.StringHasSuffix(path, string(os.PathSeparator)+"*") {
			dir := path[:len(path)-1]
			if !tool.FileExists(dir) || !tool.FileIsDir(dir) {
				continue
			}

			source, err := self.parseWildcardSource(path)
			if err != nil {
				return err
			}
			self.sources = append(self.sources, source)
		}

		if !tool.FileExists(path) {
			continue
		}
		if tool.StringHasSuffix(path, ".jar", ".zip") {
			if tool.FileIsDir(path) {
				continue
			}
			source, err := self.parseJarSource(path)
			if err != nil {
				return err
			}
			self.sources = append(self.sources, source)
		}

		if !tool.FileIsDir(path) {
			continue
		}
		source, err := self.parseDirSource(path)
		if err != nil {
			return err
		}
		self.sources = append(self.sources, source)
	}
	return
}

func (self *Classpath) parseJarSource(path string) (source javaSource, err error) {
	_, err = filepath.Abs(path)
	if err != nil {
		return
	}
	source = &javaSourceJar{
		p: path,
	}
	return
}

func (self *Classpath) parseDirSource(path string) (source javaSource, err error) {
	_, err = filepath.Abs(path)
	if err != nil {
		return
	}
	source = &javaSourceDir{
		p: path,
	}
	return
}

func (self *Classpath) parseWildcardSource(path string) (source javaSource, err error) {
	baseDir := path[:len(path)-1]
	_, err = filepath.Abs(baseDir)
	if err != nil {
		return
	}
	source, err = newWildcardJavaSource(path)
	return
}
