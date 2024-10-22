# Todo CLI Application

A simple command-line interface (CLI) application for managing tasks. This CLI allows you to add, list, complete, and delete tasks via commands in the terminal.

## Features

- Add a new task.
- List tasks (with an option to list all tasks or just pending tasks).
- Mark a task as complete.
- Delete a task.

## Prerequisites

Before you can run this project, make sure you have the following installed on your machine:

- [Go (1.23.2 or higher)](https://golang.org/doc/install)
- A terminal or command-line interface (e.g., zsh, bash)

## Project Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/gurshaan17/go-lang/todo-cli.git
   cd todo-cli

	2.	Install Dependencies:
The project uses Go modules for dependency management. To download and install the required dependencies, run:

go mod tidy


	3.	Build the Application:
To compile the CLI application into a binary, run the following command from the project root:

go build -o tasks

This will generate an executable file named tasks.

	4.	Make the Binary Executable (Optional):
If necessary, make sure the tasks binary has execution permissions:

chmod +x tasks



Running the Application

Once the project is set up and the binary is built, you can use the following commands to manage your tasks.

Add a Task

To add a task, use the add command followed by the task description in quotes:

./tasks add "My new task"

List Tasks

To list uncompleted tasks:

./tasks list

To list all tasks (including completed ones):

./tasks list --all

Mark a Task as Complete

To mark a task as completed, use the complete command followed by the task ID:

./tasks complete <task-id>

Delete a Task

To delete a task, use the delete command followed by the task ID:

./tasks delete <task-id>

File Locking

To prevent issues with multiple processes accessing the task list at the same time, the application uses file locking when reading and writing tasks to the CSV file. This ensures that tasks are written and read safely.

Example Data File

The task data is stored in a CSV file. Here is an example of what the file looks like:

ID,Description,CreatedAt,IsComplete
1,Tidy up my desk,2024-07-27T16:45:19-05:00,false
2,Write up documentation for the new project feature,2024-07-27T16:45:26-05:00,true
3,Change my keyboard mapping to use escape/control,2024-07-27T16:45:31-05:00,false

Known Issues

	•	Make sure to always use the correct format when running the commands.
	•	Ensure the tasks binary is built for your architecture. If you encounter issues related to architecture, rebuild the binary using the appropriate GOARCH and GOOS settings.
