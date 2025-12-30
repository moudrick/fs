package finddup

import (
	"errors"
	"path"
	"strings"
)

type Node struct {
	Dir      bool
	Content  []byte
	Children map[string]*Node
}

func File(content string) *Node {
	return &Node{Dir: false, Content: []byte(content)}
}

func Dir(children map[string]*Node) *Node {
	return &Node{Dir: true, Children: children}
}

type MemFS struct {
	root *Node
}

func NewMemFS(root *Node) *MemFS {
	return &MemFS{root: root}
}

func (fs *MemFS) IsDir(p string) (bool, error) {
	n, err := fs.get(p)
	if err != nil {
		return false, err
	}
	return n.Dir, nil
}

func (fs *MemFS) ListDir(p string) ([]string, error) {
	n, err := fs.get(p)
	if err != nil {
		return nil, err
	}
	if !n.Dir {
		return nil, errors.New("not dir")
	}
	out := make([]string, 0, len(n.Children))
	for name := range n.Children {
		out = append(out, name)
	}
	return out, nil
}

func (fs *MemFS) ReadFile(p string) ([]byte, error) {
	n, err := fs.get(p)
	if err != nil {
		return nil, err
	}
	if n.Dir {
		return nil, errors.New("is dir")
	}
	return append([]byte(nil), n.Content...), nil
}

func (fs *MemFS) Join(elem ...string) string {
	return path.Join(elem...)
}

func (fs *MemFS) get(p string) (*Node, error) {
	if p == "/" || p == "" {
		return fs.root, nil
	}
	parts := strings.Split(strings.Trim(p, "/"), "/")
	cur := fs.root
	for _, part := range parts {
		if !cur.Dir {
			return nil, errors.New("invalid path")
		}
		n, ok := cur.Children[part]
		if !ok {
			return nil, errors.New("not found")
		}
		cur = n
	}
	return cur, nil
}
