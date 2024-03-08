package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

const (
	file = ".todo" // a to-do list is saved per directory
	fileMode = 0644
)

type todoList []string

func usage() bool {
	if !flag("help") && !flag("-h") && !flag("--help") &&
	!flag("usage") && !flag("-u") && !flag("--usage") {
		return false
	}
	fmt.Println(
`usage: todo [top|pop|all|help]

'todo' prompts for a task to add to list and exits
  - most recent task added is popped (FIFO)
  - if the task begins with "!" it will be added to bottom

 subcommands:
    'top'	prints the next task to-do
    'pop'	 removes from list and prints next task
    'all'	adds multiple tasks, type "q" to finish
    'help'	will print this message.

the tasks are saved as '.todo' in current directory

type 'more .todo' to show all tasks

any list of arguments that does not start with a subcommand is treated as a new task`)
	fmt.Println()
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
	if todo.top() || todo.pop() || todo.start() || todo.swap(0) {
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
		if !confirm("no todo, create new?") {
			os.Exit(0)
		}
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
	if td == "" {
		return
	}
	td = strings.TrimSpace(td)
	if td[:1] == "!" { // safe, as we know which rune
		*todo = append(*todo, td[1:])
		return
	}
	*todo = append(todoList{td}, *todo...) // default fifo
}

func checkAndPrint(t todoList) {
	if len(t) < 1 {
		fmt.Println("-everything done!-")
		return
	}
	fmt.Println(t[0])
}

func (todo todoList) top() bool {
	if !flag("top") {
		return false
	}
	checkAndPrint(todo)
	return true
}

func (todo *todoList) pop() bool {
	if !flag("pop") {
		return false
	}
	if len(*todo) < 1 {
		return true
	}
	*todo = (*todo)[1:]
	if len(*todo) > 0 {
		fmt.Println("next:")
	}
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

// bonus feature
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
	if confirm("swap deeper?") {
		i++
		todo.swap(i)
	}
	return true
}

func confirm(prompt string) bool {
	fmt.Printf(prompt+" y/n/_  ")
	s := ""
	fmt.Scanln(&s)
	if s == "y" || s == "yes" {
		return true
	}
	return false
}
