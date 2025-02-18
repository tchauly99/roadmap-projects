# Task Tracker

Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## How to run

Clone the repository:

```bash
git clone https://github.com/tchauly99/roadmap-projects.git
cd task-tracker
```

Init the tasktracker module:

```bash
go mod init tasktracker
```

Build task-cli executable:

```bash
go build -o task-cli
```

## Usage

```bash
./task-cli --help # To see usage
```

### To list out all tasks
```bash
./task-cli list
```

### To list out tasks based on status
```bash
./task-cli list todo
```
```bash
./task-cli list in-progress
```
```bash
./task-cli list done
```

### To add one/multiple task
```bash
./task-cli add <Task 1 description> <Task 2 description> ...
```

### To update a task description
```bash
./task-cli update <Task ID> <New description>
```

### To delete one/multiple task
```bash
./task-cli delete <Task ID 1> <Task ID 2> ...
```

### To mark one/multiple task as in progress/done/todo
```bash
./task-cli mark-in-progress <Task ID 1> <Task ID 2> ...
```
```bash
./task-cli mark-done <Task ID 1> <Task ID 2> ...
```
```bash
./task-cli mark-todo <Task ID 1> <Task ID 2> ...
```

## References
[Cobra User Guide](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)

[JSON in GO](https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go)

[Parsing JSON file with GO](https://tutorialedge.net/golang/parsing-json-with-golang/)

[encoding/json document](https://pkg.go.dev/encoding/json)

[Referred solution by arikchakma](https://github.com/arikchakma/backend-projects/tree/main/task-tracker)

