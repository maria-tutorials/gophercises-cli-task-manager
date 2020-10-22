# gophercises-cli-task-manager

CLI tool that can be used to manage your TODOs in the terminal. 

#### The basic usage of the tool is going to look roughly like this:
```
$ task
task is a CLI for managing your TODOs.

Usage:
  task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
Added "review talk proposal" to your task list.

$ task add clean dishes
Added "clean dishes" to your task list.

$ task list
You have the following tasks:
1. review talk proposal
2. some task description

$ task do 1
You have completed the "review talk proposal" task.

$ task list
You have the following tasks:
1. some task description
```


### Part 1
Build the CLI shell

### Part 2
Write the BoltDB interactions

### Part 3
Put it all together

### Extras
- The `rm` command will delete a task instead of completing it.

- The `completed` command will list out any tasks completed in the same day. You can define this however you want (last 12hrs, last 24hrs, or the same calendar date).
