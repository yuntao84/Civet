package classpath

import (
	"github.com/yuntao84/Civet/tool"
	"os"
	"path/filepath"
)

type javaSourceWildcard struct {
	p       string
	sources []javaSource
}

func newWildcardJavaSource(p string) (*javaSourceWildcard, error) {
	js := &javaSourceWildcard{p: p}
	baseDir := p[:len(p)-1]
	absPath, err := filepath.Abs(baseDir)
	if err != nil {
		return nil, err
	}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != absPath {
			return filepath.SkipDir
		}
		if tool.StringHasSuffix(path, ".jar", ".zip") {
			js.sources = append(js.sources, &javaSourceJar{
				p: filepath.Join(baseDir, info.Name()),
			})
		}
		return nil
	}
	err = filepath.Walk(absPath, walkFn)
	if err != nil {
		return nil, err
	}
	return js, nil
}

func (self *javaSourceWildcard) path() string {
	return self.p
}

func (self *javaSourceWildcard) readClass(className string) (has bool, data []byte, err error) {
	return self.readResource(className + ".class")
}

func (self *javaSourceWildcard) readResource(path string) (has bool, data []byte, err error) {
	for _, source := range self.sources {
		has, data, err = source.readResource(path)
		if has {
			return has, data, err
		}
	}
	return
}
