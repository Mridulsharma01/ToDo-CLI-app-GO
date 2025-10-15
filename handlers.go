package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func getCall(getFlag *flag.FlagSet, getAll *bool, getId *int) {
	// read the json
	err := getFlag.Parse(os.Args[2:])
	checkError(err)

	if *getId > 0 {
		beautyPrint(getTaskByID(*getId))
	} else {
		if *getAll {

			beautyPrint(getAllTasks())
		} else {
			fmt.Println("Provide corect Id please or set --all=true")
		}
	}

}
func addCall(addFlag *flag.FlagSet, id *int, title *string, description *string, status *string) {
	// read the json
	err := addFlag.Parse(os.Args[2:])
	checkError(err)
	if *id < 1 || *title == "Nil" {
		fmt.Println("Id and Title are required to add a new task!")
		os.Exit(1)
	}
	tasks := getAllTasks()
	for _, task := range tasks {
		if task.Id == *id {
			fmt.Println("Id is already present. Please choose another Id")
			os.Exit(1)
		}
	}
	var newTask ToDo
	newTask.Id = *id
	newTask.Title = *title
	if description != nil {
		newTask.Description = *description
	}
	newTask.Status = "Not_Done"
	tasks = append(tasks, newTask)
	jsonData, err := json.Marshal(tasks)
	checkError(err)
	err = os.WriteFile("todo.json", jsonData, 0755)
	checkError(err)
	fmt.Println("Task Added!")
}
func updateCall(updateFlag *flag.FlagSet, id *int, title *string, description *string, status *string) {
	// read the json
	err := updateFlag.Parse(os.Args[2:])
	if *id < 1 {
		fmt.Println("Id cannot be less than 1")
		os.Exit(1)
	} else {
		tasks := getAllTasks()
		for i, task := range tasks {
			if *id == task.Id {
				if *title != "Nil" {

					task.Title = *title
				}
				if *description != "Nil" {
					task.Description = *description
				}
				if *status != "Nil" {
					task.Status = *status
				}
				tasks[i] = task
				fmt.Println("updated task!")
				break
			}
		}

		jsonData, err := json.Marshal(tasks)
		checkError(err)
		err = os.WriteFile("todo.json", jsonData, 0755)
		checkError(err)
	}
	checkError(err)
}
func deleteCall(deleteFlag *flag.FlagSet, id *int, deleteAll *bool) {
	// read the json
	err := deleteFlag.Parse(os.Args[2:])
	checkError(err)
	if *deleteAll == true {
		var emptyList []ToDo
		jsonData, err := json.Marshal(emptyList)
		checkError(err)
		err = os.WriteFile("todo.json", jsonData, 077)
		checkError(err)
		fmt.Println("Deleted All Tasks!")
		os.Exit(1)
	}
	tasks := getAllTasks()
	toDelete := -1
	for i, task := range tasks {
		if task.Id == *id {
			toDelete = i
			break
		}
	}
	if toDelete == -1 {
		fmt.Println("Task not found with the given Id!")
		os.Exit(1)
	}
	tasks = append(tasks[:toDelete], tasks[toDelete+1:]...)
	jsonData, err := json.Marshal(tasks)
	checkError(err)
	err = os.WriteFile("todo.json", jsonData, 0755)
	checkError(err)
	fmt.Println("Task Deleted!")
}
func statusCall(statusFlag *flag.FlagSet, id *int) {
	// read the json
	err := statusFlag.Parse(os.Args[2:])
	checkError(err)
	tasks := getAllTasks()
	for i, task := range tasks {
		if task.Id == *id {
			task.Status = "Done"
			tasks[i] = task
			fmt.Println("Status Marked Done!")
			break
		}
	}
	jsonData, err := json.Marshal(tasks)
	checkError(err)
	err = os.WriteFile("todo.json", jsonData, 0755)
	checkError(err)
}

func checkError(e error) {

	if e != nil {
		fmt.Println("Error Occured!", e)
		os.Exit(1)
	}

}

func getAllTasks() []ToDo {
	jsonData, err := os.ReadFile("todo.json")
	checkError(err)
	var todoData []ToDo

	if len(jsonData) == 0 {
		var emptyList []ToDo
		data, err := json.Marshal(emptyList)
		checkError(err)
		os.WriteFile("todo.json", data, 0755)
		return todoData
	}
	err = json.Unmarshal(jsonData, &todoData)
	checkError(err)

	return todoData
}

func getTaskByID(id int) []ToDo {
	books := getAllTasks()
	var result []ToDo
	for i := range books {
		if books[i].Id == id {
			result = append(result, books[i])
			break
		}
	}

	return result
}

func beautyPrint(tasks []ToDo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Title", "Description", "Status"})

	for _, task := range tasks {
		var entry []string
		entry = append(entry, fmt.Sprintf("%d", task.Id))
		entry = append(entry, task.Title)
		entry = append(entry, task.Description)
		entry = append(entry, task.Status)

		table.Append(entry)
	}
	table.Render()
}
