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

Suggested use case: use `todo` to collect typos or other diversions within a project to return to later on, to keep the current commit clean

Tested on FreeBSD, expected to be compatible with any unix-like/POSIX system

The design maximises effectiveness by placing focus on the only the next task to-do, and also by implicitly categorising tasks per project directory

Any errors are simply printed to stdout, as this is a human-facing utlitity and not intended to be used in scripts

Beta feature: `todo swap` will shunt away a task, by swapping up one from deeper in the stack. Will prompt to go deeper, as required

It might make more sense for the present `pop` to be called `top` and `done` to be `pop`, as that is more semantic

User feedback, code review comments or bug reports are very much welcome
