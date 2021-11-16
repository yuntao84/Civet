package classpath

type javaSource interface {
	path() string
	readResource(path string) (has bool, data []byte, err error)
	readClass(className string) (has bool, data []byte, err error)
}
