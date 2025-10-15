package main

import (
	"flag"
	"fmt"
	"os"
)

// "Flag"
type ToDo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func main() {

	// flag setup
	getFlag := flag.NewFlagSet("get", flag.ExitOnError)
	deleteFlag := flag.NewFlagSet("delete", flag.ExitOnError)
	addFlag := flag.NewFlagSet("add", flag.ExitOnError)
	updateFlag := flag.NewFlagSet("update", flag.ExitOnError)
	statusFlag := flag.NewFlagSet("status", flag.ExitOnError)

	getAll := getFlag.Bool("all", true, "Lists all the To-Do tasks")
	getById := getFlag.Int("id", 0, "Id of the to-do task")

	// adding arguments flags
	addId := addFlag.Int("id", 0, "Id of the task")
	addTitle := addFlag.String("title", "Nil", "Title here..")
	addDescription := addFlag.String("description", "Nil", "description here..")
	addStatus := addFlag.String("status", "Not_Done", "Status here..")

	// update arguments flags
	updateId := updateFlag.Int("id", 0, "Id of the task")
	updateTitle := updateFlag.String("title", "Nil", "Title here..")
	updateDescription := updateFlag.String("description", "Nil", "description here..")
	updateStatus := updateFlag.String("status", "Nil", "Status here..")

	// delete ..
	deleteId := deleteFlag.Int("id", 0, "Delete Id here")
	deleteAll := deleteFlag.Bool("all", false, "Delete all Tasks")

	//status ..
	statusId := statusFlag.Int("id", 0, "status id")

	if len(os.Args) < 2 {
		fmt.Println("Enter sufficient arguments!")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		getCall(getFlag, getAll, getById)

	case "add":
		addCall(addFlag, addId, addTitle, addDescription, addStatus)

	case "update":
		updateCall(updateFlag, updateId, updateTitle, updateDescription, updateStatus)

	case "delete":
		deleteCall(deleteFlag, deleteId, deleteAll)

	case "status":
		statusCall(statusFlag, statusId)
	}

}
