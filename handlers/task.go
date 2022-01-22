package handlers

import (
	"ToDoAPP/controllers"
	"ToDoAPP/utilities"
	"ToDoAPP/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/spf13/cast"
)


func TaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		utilities.URLReturnResponseJson(w, getTask(w, r))
	} else if r.Method == "POST" {
		utilities.URLReturnResponseJson(w, addTask(w, r))
	} else if r.Method == "PUT" {
		utilities.URLReturnResponseJson(w, updateTask(w, r))
	}else if r.Method == "DELETE"{
		utilities.URLReturnResponseJson(w, deleteTask(w, r))
	}else{
		utilities.URLReturnResponseJson(w, "error")
	}
   
}

func getTask(w http.ResponseWriter, r *http.Request)(returnData utilities.ResponseJSON){

	Id := r.URL.Query().Get("id")
	userId := r.URL.Query().Get("userId")
	status := r.URL.Query().Get("status")


	switch true {

	case utilities.IsEmpty(Id) && utilities.IsEmpty(userId) &&  utilities.IsEmpty(status):
		return controllers.GetAllTaskFunction()

	case !utilities.IsEmpty(Id) && utilities.IsEmpty(userId) &&  utilities.IsEmpty(status):
		return controllers.GetTaskByIdFunction(cast.ToInt(Id))
	
	case  utilities.IsEmpty(Id) && !utilities.IsEmpty(userId) &&  utilities.IsEmpty(status):
		return controllers.GetAllTaskByUseridFunction(cast.ToInt(userId))
	
		
	case  utilities.IsEmpty(Id) && !utilities.IsEmpty(userId) &&  !utilities.IsEmpty(status):
		return controllers.GetAllTaskByUserIdAndStatusFunction(cast.ToInt(userId),status)

	default:
		utilities.UnprocessableResponse(&returnData)
		}
	return
}



func addTask(w http.ResponseWriter, r *http.Request)(returnData utilities.ResponseJSON){

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: incorrect data!")
		return
	}

	inputobj := models.Task{}
	json.Unmarshal(body, &inputobj)


	returnData = controllers.AddTaskFunction(inputobj)
	return

}
func updateTask(w http.ResponseWriter, r *http.Request)(returnData utilities.ResponseJSON){

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: incorrect data!")
		return
	}

	inputobj := models.Task{}
	json.Unmarshal(body, &inputobj)

	returnData = controllers.UpdateTaskFunction(inputobj)
	return


}

func deleteTask(w http.ResponseWriter, r *http.Request)(returnData utilities.ResponseJSON){

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: incorrect data!")
		return
	}

	inputobj := models.Task{}
	json.Unmarshal(body, &inputobj)


	returnData = controllers.DeleteTaskFunction(inputobj)
	return


}