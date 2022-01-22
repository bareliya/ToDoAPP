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



func UserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		utilities.URLReturnResponseJson(w, getUser(w, r))
	} else if r.Method == "POST" {
		utilities.URLReturnResponseJson(w, addUser(w, r))
	} else if r.Method == "DELETE" {
		utilities.URLReturnResponseJson(w, deleteUser(w, r))
	}else {
		utilities.URLReturnResponseJson(w, "error")
	}
   
}

func getUser(w http.ResponseWriter, r *http.Request)(returnData utilities.ResponseJSON){
	Id := r.URL.Query().Get("id")

	switch true {

	case utilities.IsEmpty(Id) :
		return controllers.GetAllUserFunction()

	case !utilities.IsEmpty(Id):
		return controllers.GetUserByIdFunction(cast.ToInt(Id))

	default:
		utilities.UnprocessableResponse(&returnData)
		}

	return

}
func addUser(w http.ResponseWriter, r *http.Request)(returnData utilities.ResponseJSON){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: incorrect data!")
		return
	}

	inputobj := models.User{}
	json.Unmarshal(body, &inputobj)


	returnData = controllers.AddUserFunction(inputobj)
	return


}
func deleteUser(w http.ResponseWriter, r *http.Request)(returnData utilities.ResponseJSON){
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: incorrect data!")
		return
	}

	inputobj := models.User{}
	json.Unmarshal(body, &inputobj)


	returnData = controllers.DeleteUserFunction(inputobj)
	return


}