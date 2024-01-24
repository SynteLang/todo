package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	file = ".todo" // a to-do list is saved per directory
	fileMode = 0644
)

type todoList []string

func usage() bool {
	if !flag("help") && !flag("-h") && !flag("--help") {
		return false
	}
	fmt.Println(
`usage: todo [pop|done|all|help]

'todo' prompts for a task to add to list and exits
  - most recent task added is popped (FIFO)
  - if the task begins with "!" it will be added to bottom)

 subcommands:
    'pop'	prints the number of tasks and the next task to-do
    'done'	prints and removes from list
    'all'	adds multiple tasks, type "q" to finish
    'help'	will print this message.

the tasks are saved as '.todo' in current directory
'less .todo' to show all tasks`)
	return true
}

func main() {
	if usage() {
		return
	}
	todo, f := load(file)
	if f == nil {
		return
	}
	defer todo.save()
	if len(os.Args) < 2 {
		t := bufio.NewScanner(os.Stdin)
		t.Scan()
		todo.pushString(t.Text())
		return
	}
	if todo.pop() || todo.done() || todo.start() || todo.swap(0) {
		return
	}
	s := collate(os.Args)
	todo.pushString(s)
}

var flag = flagFunc // for testing purposes
func flagFunc(s string) bool {
	if len(os.Args) > 2 {
		return false
	}
	if len(os.Args) == 2 && os.Args[1] == s {
		return true
	}
	return false
}

func collate(a []string) (s string) {
	for i := 1; i < len(a); i++ {
		s += a[i] + " " // powered by a[i]
	}
	return s
}

func open(file string) (*os.File) {
	f, err := os.Open(file)
	if err != nil {
		f, err = os.Create(file)
		if err != nil {
			return nil
		}
	}
	return f
}

func load(file string) (todoList, *os.File) {
	f := open(file)
	defer f.Close()
	if f == nil {
		fmt.Printf("unable to open %q\n", file)
		return nil, nil
	}
	s := bufio.NewScanner(f)
	todo := todoList{}
	for s.Scan() {
		todo = append(todo, s.Text())
	}
	return todo, f
}

func (todo *todoList) save() {
	var td string
	for _, v := range *todo {
		td += v + "\n"
	}
	err := os.WriteFile(file, []byte(td), fileMode)
	if err != nil {
		fmt.Printf("file not saved :( %v\n", err)
	}
	os.Exit(0)
}

func (todo *todoList) pushString(td string) {
	if len(td) < 1 {
		return
	}
	if td[:1] != "!" { // safe, as we know which rune
		*todo = append(todoList{td}, *todo...) // default fifo
	} else {
		*todo = append(*todo, td[1:])
	}
	return
}

func checkAndPrint(t todoList) {
	l := len(t)
	if l < 1 {
		return // nothing to print
	}
	fmt.Println(t[0])
}

func (todo todoList) pop() bool {
	if !flag("pop") {
		return false
	}
	checkAndPrint(todo)
	return true
}

func (todo *todoList) done() bool {
	if !flag("done") {
		return false
	}
	if len(*todo) < 1 {
		return true
	}
	fmt.Println("next:")
	*todo = (*todo)[1:]
	checkAndPrint(*todo)
	return true
}

func (todo *todoList) start() bool {
	if !flag("all") {
		return false
	}
	s := bufio.NewScanner(os.Stdin)
	td := []string{}
	for s.Scan() {
		txt := s.Text()
		if txt == "q" || txt == ":wq" {
			break
		}
		td = append(td, txt)
	}
	*todo = append(td, *todo...)
	return true
}

// undocumented bonus feature
func (todo *todoList) swap(i int) bool {
	if !flag("swap") {
		return false
	}
	l := len(*todo) 
	if l < i+2 {
		fmt.Println("unable to swap, not enough items")
		return true
	}
	(*todo)[0], (*todo)[i+1] = (*todo)[i+1], (*todo)[0]
	fmt.Println((*todo)[0])
	fmt.Printf("swap deeper? y/n  ")
	s := ""
	fmt.Scanln(&s)
	if s == "y" || s == "yes" {
		i++
		todo.swap(i)
	}
	return true
}
