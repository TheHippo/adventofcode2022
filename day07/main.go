package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var matchDir = regexp.MustCompile(`dir\s([[:graph:]]+)`)
var matchFile = regexp.MustCompile(`([[:digit:]]+)\s([[:graph:]]+)`)

type direction uint8

const (
	_ direction = iota
	root
	up
	to
)

type CommmandDir struct {
	current          *Tree
	to               string
	commandDirection direction
}

func (d *CommmandDir) Execute(dir *Tree, cmd string) *Tree {
	d.current = dir
	switch cmd {
	case "/":
		d.commandDirection = root
		for d.current.parent.parent != nil {
			d.current = d.current.parent
		}
	case "..":
		d.commandDirection = up
		if d.current.parent.parent == nil {
			log.Fatal("can go up")
		}
		d.current = d.current.parent
	default:
		// going into a folder
		d.to = cmd
		d.commandDirection = to
		if dir, exists := d.current.dirs[d.to]; !exists {
			log.Fatalf("sub directry %s does not exist", d.to)
		} else {
			d.current = dir
		}
	}
	return d.current
}

func (d *CommmandDir) Input(s string) {
	// there is no output
}

type ListDir struct {
	current *Tree
}

func (l *ListDir) Execute(dir *Tree, cmd string) *Tree {
	l.current = dir
	return l.current
}

func (l *ListDir) Input(s string) {
	if matches := matchFile.FindAllStringSubmatch(s, -1); len(matches) == 1 {
		size, err := strconv.ParseUint(matches[0][1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		l.current.files[matches[0][2]] = uint(size)
	} else if matches := matchDir.FindAllStringSubmatch(s, -1); len(matches) == 1 {
		l.current.dirs[matches[0][1]] = &Tree{
			parent: l.current,
			dirs:   map[string]*Tree{},
			files:  map[string]uint{},
		}
	} else {
		log.Fatalf("unknown input %s", s)
	}
}

type Command interface {
	Execute(dir *Tree, cmd string) *Tree
	Input(s string)
}

type Tree struct {
	parent   *Tree
	dirs     map[string]*Tree
	files    map[string]uint
	selfSize uint
}

func (t *Tree) walk(f func(name string, t *Tree)) {
	for n, d := range t.dirs {
		f(n, d)
		d.walk(f)
	}
}

func (t *Tree) size() (result uint) {
	if t.selfSize != 0 {
		return t.selfSize
	}
	for _, d := range t.dirs {
		result += d.size()
	}
	for _, f := range t.files {
		result += f
	}
	t.selfSize = result
	return
}

func (t *Tree) print(out io.Writer, depth int) {
	prefix := strings.Repeat("  ", depth)
	for dirName, dir := range t.dirs {
		fmt.Fprintf(out, "%s- %s (dir)\n", prefix, dirName)
		dir.print(out, depth+1)
	}
	for filename, file := range t.files {
		fmt.Fprintf(out, "%s- %s (file, size=%d)\n", prefix, filename, file)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	root := &Tree{
		dirs:  map[string]*Tree{},
		files: map[string]uint{},
	}

	absRoot := &Tree{
		dirs: map[string]*Tree{
			"/": root,
		},
	}
	root.parent = absRoot

	currentDir := root
	var currentCommand Command

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "$") {
			option := ""
			// command
			if strings.HasPrefix(line, "$ cd") {
				option = strings.TrimPrefix(line, "$ cd ")
				currentCommand = &CommmandDir{}
			} else if strings.HasPrefix(line, "$ ls") {
				currentCommand = &ListDir{}
			} else {
				log.Fatalf("unknown command: %s", line)
			}
			currentDir = currentCommand.Execute(currentDir, option)
		} else {
			// command output
			currentCommand.Input(line)
		}

	}

	absRoot.print(os.Stdout, 0)

	var total uint
	root.walk(func(name string, t *Tree) {
		if t.size() <= 100000 {
			log.Printf("%s - %d", name, t.size())
			total += t.size()
		}
	})

	log.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
