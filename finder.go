package finddup

import (
	"crypto/sha256"
	"encoding/hex"
)

func FindDuplicatedFiles(root string, fs FS) [][]string {
	hashToPaths := make(map[string][]string)

	var walk func(string)
	walk = func(p string) {
		isDir, err := fs.IsDir(p)
		if err != nil {
			return
		}

		if isDir {
			children, err := fs.ListDir(p)
			if err != nil {
				return
			}
			for _, name := range children {
				walk(fs.Join(p, name))
			}
			return
		}

		data, err := fs.ReadFile(p)
		if err != nil {
			return
		}

		h := hashBytes(data)
		hashToPaths[h] = append(hashToPaths[h], p)
	}

	walk(root)

	var res [][]string
	for _, g := range hashToPaths {
		if len(g) >= 2 {
			res = append(res, g)
		}
	}
	return res
}
 
func hashBytes(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
