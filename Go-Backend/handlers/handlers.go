package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetTasks(rw http.ResponseWriter, r *http.Request) {

	tasks := models.Tasks{}
	db.Database.Find(&tasks)
	sendData(rw, tasks, http.StatusOK)

}

func GetTask(rw http.ResponseWriter, r *http.Request) {
	if task, err := getTaskById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, task, http.StatusOK)
	}

}

func getTaskById(r *http.Request) (models.Task, *gorm.DB) {
	//Obtener ID
	vars := mux.Vars(r)
	taskId, _ := strconv.Atoi(vars["id"])
	task := models.Task{}

	if err := db.Database.First(&task, taskId); err.Error != nil {
		return task, err
	} else {
		return task, nil
	}
}

func CreateTask(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	task := models.Task{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&task); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&task)
		sendData(rw, task, http.StatusCreated)
	}
}

func UpdateTask(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	var taskId int64

	if task_ant, err := getTaskById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		taskId = task_ant.Id

		task := models.Task{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&task); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			task.Id = taskId
			db.Database.Save(&task)
			sendData(rw, task, http.StatusOK)
		}
	}

}

func DeleteTask(rw http.ResponseWriter, r *http.Request) {

	if task, err := getTaskById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		db.Database.Delete(&task)
		sendData(rw, task, http.StatusOK)
	}
}
