package finddup

type FS interface {
	IsDir(path string) (bool, error)
	ListDir(path string) ([]string, error)
	ReadFile(path string) ([]byte, error)
	Join(elem ...string) string
}
