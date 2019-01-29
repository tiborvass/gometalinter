// +build generate

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	if err := generate("tools.go"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func generate(outputFile string) error {
	l, err := strconv.Atoi(os.Getenv("GOLINE"))
	if err != nil {
		return err
	}
	f, err := os.Open(os.Getenv("GOFILE"))
	if err != nil {
		return err
	}
	defer f.Close()

	g := grep{f: f, from: l}

	list, err := g.reduce("InstallFrom:", reduceFn)
	if err != nil {
		return err
	}

	sort.Strings(list)

	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()
	tmpl := template.Must(template.New("tools.go").Parse(toolsTemplate))
	return tmpl.Execute(out, list)
}

func reduceFn(s string) (string, error) {
	k1 := strings.IndexByte(s, '"')
	if k1 < 0 {
		return "", fmt.Errorf(`Couldn't find opening quote character '"' in: %s`, s)
	}
	s2 := s[k1+1:]
	k2 := strings.IndexByte(s2, '"')
	if k2 < 0 {
		return "", fmt.Errorf(`Couldn't find closing quote character '"' in: %s`, s)
	}
	return s2[:k2], nil
}

type grep struct {
	f    io.Reader
	from int
}

func (g *grep) reduce(filter string, fn func(s string) (string, error)) ([]string, error) {
	list := make([]string, 0, 16)
	i := 0
	s := bufio.NewScanner(g.f)
	lineReached := false
	for s.Scan() {
		i++
		if i == g.from {
			lineReached = true
			continue
		}
		if lineReached {
			t := s.Text()
			if strings.Contains(t, filter) {
				reduction, err := fn(t)
				if err != nil {
					return nil, err
				}
				list = append(list, reduction)
			}
		}
	}
	return list, s.Err()
}

const toolsTemplate = `// +build tools

package main

import ({{ range . }}
	_ {{ printf "%q" . }}{{ end }}
)
`
