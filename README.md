# todo

`todo` is a minimalist CLI utility for efficient workflows

>
	usage: todo [pop|done|all|help]

	'todo' prompts for a task to add to list and exits
	  - most recent task added is popped (FIFO)
	  - if the task begins with "!" it will be added to bottom)

	 subcommands:
		'pop'	prints the number of tasks and the next task to-do
		'done'	prints and removes from list
		'all'	adds multiple tasks, type "q" to finish
		'help'	will print this message.

	the tasks are saved as '.todo' in current directory
	'less .todo' to show all tasks

Suggested use case: use todo to collect typos or other diversions to return to later on, to keep the current commit clean

Tested on FreeBSD, expected to be compatible with any unix-like/POSIX system

Code review comments or bug reports are very much welcome
