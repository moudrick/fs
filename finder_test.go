package finddup

import (
	"reflect"
	"testing"
)

// --- table-driven tests ----------------------------------------------------

func TestFindDuplicatedFiles(t *testing.T) {
	fs0 := NewMemFS(Dir(map[string]*Node{
			"f1": File("*"),
			"f2": File("**"),
			"d1": Dir(map[string]*Node{
				"f3": File("***"),
			}),
			"d2": Dir(map[string]*Node{
				"f4": File("**"),
			}),
			"d3": Dir(map[string]*Node{
				"f5": File("*"),
				"f6": File("*"),
			}),
		}))
	tests := []struct {
		name string
		root string
		fs   *MemFS
		want [][]string
	}{
		{
			name: "root duplicates across subtrees",
			root: "/",
			fs: fs0,
			want: [][]string{
				{"/f1", "/d3/f5", "/d3/f6"},
				{"/f2", "/d2/f4"},
			},
		},

		{
			name: "subtree only (with siblings)",
			root: "/d3",
			fs: fs0,
			want: [][]string{
				{"/d3/f5", "/d3/f6"},
			},
		},
		{
			name: "subtree only (no siblings)",
			root: "/d3",
			fs: NewMemFS(Dir(map[string]*Node{
				"d3": Dir(map[string]*Node{
					"f5": File("*"),
					"f6": File("*"),
				}),
			})),
			want: [][]string{
				{"/d3/f5", "/d3/f6"},
			},
		},

		{
			name: "no duplicates",
			root: "/",
			fs: NewMemFS(Dir(map[string]*Node{
				"a": File("1"),
				"b": File("2"),
				"d": Dir(map[string]*Node{
					"c": File("3"),
				}),
			})),
			want: nil,
		},

		{
			name: "single file only",
			root: "/",
			fs: NewMemFS(Dir(map[string]*Node{
				"a": File("x"),
			})),
			want: nil,
		},

		{
			name: "empty directory",
			root: "/",
			fs: NewMemFS(Dir(map[string]*Node{})),
			want: nil,
		},

		{
			name: "deep nested duplicates",
			root: "/",
			fs: NewMemFS(Dir(map[string]*Node{
				"d1": Dir(map[string]*Node{
					"d2": Dir(map[string]*Node{
						"a": File("Z"),
					}),
				}),
				"d3": Dir(map[string]*Node{
					"b": File("Z"),
					"c": File("Z"),
				}),
			})),
			want: [][]string{
				{"/d1/d2/a", "/d3/b", "/d3/c"},
			},
		},

		{
			name: "duplicates with empty content",
			root: "/",
			fs: NewMemFS(Dir(map[string]*Node{
				"a": File(""),
				"b": File(""),
				"c": File("x"),
			})),
			want: [][]string{
				{"/a", "/b"},
			},
		},

		{
			name: "duplicate file names but different content",
			root: "/",
			fs: NewMemFS(Dir(map[string]*Node{
				"a": File("1"),
				"d": Dir(map[string]*Node{
					"a": File("2"),
				}),
			})),
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindDuplicatedFiles(tt.root, tt.fs)

			if !reflect.DeepEqual(
				normalize(got),
				normalize(tt.want),
			) {
				t.Fatalf(
					"unexpected result\nroot=%s\ngot =%v\nwant=%v",
					tt.root, got, tt.want,
				)
			}
		})
	}
}


// --- helpers ---------------------------------------------------------------

// normalize makes comparison order-independent.
// Used ONLY in tests.
func normalize(groups [][]string) [][]string {
	out := make([][]string, len(groups))
	for i, g := range groups {
		cp := append([]string(nil), g...)
		sortStrings(cp)
		out[i] = cp
	}
	sortGroups(out)
	return out
}

func sortStrings(s []string) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

func sortGroups(g [][]string) {
	for i := 0; i < len(g); i++ {
		for j := i + 1; j < len(g); j++ {
			if g[j][0] < g[i][0] {
				g[i], g[j] = g[j], g[i]
			}
		}
	}
}


// package finddup

// import (
// 	"reflect"
// 	"testing"
// )

// func normalize(g [][]string) [][]string {
// 	out := make([][]string, len(g))
// 	for i, x := range g {
// 		cp := append([]string(nil), x...)
// 		sortStrings(cp)
// 		out[i] = cp
// 	}
// 	sortGroups(out)
// 	return out
// }

// func sortStrings(s []string) {
// 	for i := 0; i < len(s); i++ {
// 		for j := i + 1; j < len(s); j++ {
// 			if s[j] < s[i] {
// 				s[i], s[j] = s[j], s[i]
// 			}
// 		}
// 	}
// }

// func sortGroups(g [][]string) {
// 	for i := 0; i < len(g); i++ {
// 		for j := i + 1; j < len(g); j++ {
// 			if g[j][0] < g[i][0] {
// 				g[i], g[j] = g[j], g[i]
// 			}
// 		}
// 	}
// }

// func TestFindDuplicatedFilesRoot(t *testing.T) {
// 	fs := NewMemFS(Dir(map[string]*Node{
// 		"f1": File("*"),
// 		"f2": File("**"),
// 		"d1": Dir(map[string]*Node{
// 			"f3": File("***"),
// 		}),
// 		"d2": Dir(map[string]*Node{
// 			"f4": File("**"),
// 		}),
// 		"d3": Dir(map[string]*Node{
// 			"f5": File("*"),
// 			"f6": File("*"),
// 		}),
// 	}))

// 	got := FindDuplicatedFiles("/", fs)
// 	want := [][]string{
// 		{"/f1", "/d3/f5", "/d3/f6"},
// 		{"/f2", "/d2/f4"},
// 	}

// 	if !reflect.DeepEqual(normalize(got), normalize(want)) {
// 		t.Fatalf("unexpected result\n got=%v\nwant=%v", got, want)
// 	}
// }

// func TestFindDuplicatedFilesSubtree(t *testing.T) {
// 	fs := NewMemFS(Dir(map[string]*Node{
// 		"d3": Dir(map[string]*Node{
// 			"f5": File("*"),
// 			"f6": File("*"),
// 		}),
// 	}))

// 	got := FindDuplicatedFiles("/d3", fs)
// 	want := [][]string{
// 		{"/d3/f5", "/d3/f6"},
// 	}

// 	if !reflect.DeepEqual(normalize(got), normalize(want)) {
// 		t.Fatalf("unexpected result")
// 	}
// }
