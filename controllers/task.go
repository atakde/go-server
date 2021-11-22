package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-server/core"
	"go-server/models"
	"net/http"
)

var db = core.Database()

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	// TODO :: handle here.
	if params["task"] == "" {
		return
	}

	prepareInsertion, err := db.Prepare("INSERT INTO tasks (task, is_completed) VALUES ('?', 0)")
	if err != nil {
		panic(err.Error())
	}

	prepareInsertion.Exec(params["task"])
	json.NewEncoder(w).Encode("Task created.")
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	_, err := db.Query("DELETE FROM tasks WHERE id = ?", params["id"])
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode("Task deleted.")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.Task
	jErr := json.NewDecoder(r.Body).Decode(&task)
	if jErr != nil {
		panic(jErr)
	}

	_, err := db.Query("UPDATE tasks SET task = ? WHERE id = ?", task.Task, task.Id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Task updated.")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var results []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Task, &task.IsCompleted)
		if err != nil {
			panic(err)
		}
		results = append(results, task)
	}

	json.NewEncoder(w).Encode(results)
}
