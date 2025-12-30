package main

import (
	"flag"
	"fmt"

	"finddup"
)

func main() {
	root := flag.String("root", "/", "root path")
	mem := flag.Bool("mem", true, "use in-memory fs")
	flag.Parse()

	var fs finddup.FS
	if *mem {
		fs = demoFS()
		fmt.Println("MemFS demo")
	} else {
		fs = finddup.OSFS{}
	}

	res := finddup.FindDuplicatedFiles(*root, fs)
	for i, g := range res {
		fmt.Printf("Group %d:\n", i+1)
		for _, p := range g {
			fmt.Println(" ", p)
		}
	}
}

func demoFS() *finddup.MemFS {
	return finddup.NewMemFS(finddup.Dir(map[string]*finddup.Node{
		"f1": finddup.File("*"),
		"f2": finddup.File("**"),
		"d1": finddup.Dir(map[string]*finddup.Node{
			"f3": finddup.File("***"),
		}),
		"d2": finddup.Dir(map[string]*finddup.Node{
			"f4": finddup.File("**"),
		}),
		"d3": finddup.Dir(map[string]*finddup.Node{
			"f5": finddup.File("*"),
			"f6": finddup.File("*"),
		}),
	}))
}
