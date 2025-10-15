# ToDo Cli GO

It is golang command line application to list, add, update and delete tasks using flag, json, os package


## Usage

```bash

# clone a repo
git clone https://github.com/Mridulsharma01/ToDo-CLI-app-GO.git

# build
go build

# run

# get tasks
./todo get --all=true
./todo get --id=1

# add a task with id ,title, description, status
./todo add --id=2 --title="Read 10 pages " --description="book-LLM from scratch" 

# update a task with id ,title, description, status
./todo update --id=2 --title="Read 20 pages " --description="book-GoLang for begineers" 

# delete a task by --id
./todo delete --id=2

# change status to DOne
./todo status --id=2


```
