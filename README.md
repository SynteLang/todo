# todo

`todo` is a minimalist CLI utility for efficient workflows

if you like, you can pronounce it like "dodo"

>
    usage: todo [top|pop|all|help]

    'todo' prompts for a task to add to list and exits
      - most recent task added is popped (FIFO)
      - if the task begins with "!" it will be added to bottom)

     subcommands:
        'top'   prints the next task to-do
        'pop'   removes from list and prints next task
        'all'   adds multiple tasks, type "q" to finish
        'help'  will print this message. Aliased to 'usage'

    the tasks are saved as '.todo' in current directory

    'more .todo' to show all tasks

    any list of arguments that does not start with a subcommand is treated as a new task

Suggested use case:
+   use `todo` to collect typos or other diversions within a project to return to later on, to keep the current commit clean

Tested on FreeBSD, expected to be compatible with any unix-like/POSIX system

Installation:
+ install `go` if you don't have it - [go.dev/doc/install](https://go.dev/doc/install)
+ check `echo $GOPATH` (usually '~/go/bin/') is in your $PATH
+ `go install github.com/syntelang/todo`

The design maximises effectiveness by placing focus on the only the next task to-do, and also by implicitly categorising tasks per project directory

## FAQ
+ If you have a task that takes more than one line to describe, you probably need to break it down into smaller tasks. This will help your productivity.  
Alternatively you could write the task to a file and `todo name-of-file.txt`, then later you can simply, for example `vim (todo top)` (fish shell)
+ Why not just write `// TODO` comments in your code? Well that can work well for some use cases. The benefit of `todo` is that it keeps your code - and your commits - clean, and in one place.  

Any errors are simply printed to stdout, as this is a human-facing utlitity and not intended to be used in scripts

Bonus feature:
`todo swap` will shunt away a task, by swapping up one from deeper in the stack. Will prompt to go deeper, as required. Type "y" to accept, "n" or nothing to finish.

User feedback, code review comments or bug reports are very much welcome :)
