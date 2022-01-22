package controllers

import (
	"ToDoAPP/models"
	"ToDoAPP/utilities"
	"github.com/spf13/cast"
)


func GetAllTaskFunction()(returnData utilities.ResponseJSON){
	allData, err := models.GetAllTask()
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: No task found!")
	} else {
		utilities.SuccessResponse(&returnData, allData)
	}
	return

}


func GetTaskByIdFunction(id int)(returnData utilities.ResponseJSON){
	data, err := models.GetTaskById(id)
	if err != nil {
		utilities.ErrorResponse(&returnData, "task does not exist!")
	} else {
		utilities.SuccessResponse(&returnData, data)
	}
	return

}


func GetAllTaskByUseridFunction(userid int)(returnData utilities.ResponseJSON){
	allData, err := models.GetAllTaskByUserid(userid)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: No task found!")
	} else {
		utilities.SuccessResponse(&returnData, allData)
	}
	return

}


func GetAllTaskByUserIdAndStatusFunction(userid int, status string)(returnData utilities.ResponseJSON){
	allData, err := models.GetAllTaskByUserIdAndStatus(userid,status)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: No task found!")
	} else {
		utilities.SuccessResponse(&returnData, allData)
	}
	return

}




func AddTaskFunction(task models.Task)(returnData utilities.ResponseJSON){

    user,err:= models.GetUserById(task.UserId)
	if err!=nil || user.Id==0{
		utilities.ErrorResponse(&returnData,"User Does not Exist please Register a user first")

	}

    task.Status= "pending"


	err = models.AddTask(&task)
	if err!=nil{
        utilities.ErrorResponse(&returnData,cast.ToString(err))
		return 
	}

	 utilities.SuccessResponse(&returnData, task)
	 return
}


func UpdateTaskFunction(task models.Task)(returnData utilities.ResponseJSON){
	if task.Id==0 {
		utilities.UnprocessableResponse(&returnData)
		return
	}

	oldTask,err := models.GetTaskById(task.Id)
	if err!=nil{
		utilities.ErrorResponse(&returnData,err.Error())
		return 

	}

	if task.UserId!=0 && task.UserId!= oldTask.UserId {
		utilities.ErrorResponse(&returnData,"Task Does not Exist")
		return

	}

	if utilities.IsEmpty(task.Description){
		task.Description= oldTask.Description

	}



	if task.DueDate.IsZero(){
		task.DueDate= oldTask.DueDate
		
	}

	if utilities.IsEmpty(task.Status){
		task.Status=oldTask.Status
	}else if task.Status!= oldTask.Status{
		task.Status="completed"
	}


	err=models.UpdateTask(&task)

	if err!=nil{
        utilities.ErrorResponse(&returnData,cast.ToString(err))
	}
	 utilities.SuccessResponse(&returnData,task)
	 return



	


}


func DeleteTaskFunction(task models.Task)(returnData utilities.ResponseJSON){
	if task.Id==0 {
		utilities.UnprocessableResponse(&returnData)
		return
	}
	_,err := models.GetTaskById(task.Id)
	if err!=nil{
		utilities.ErrorResponse(&returnData,err.Error())
		return 

	}

	err= models.DeleteTask(task.Id)

	if err!=nil{
        utilities.ErrorResponse(&returnData,cast.ToString(err))
		return 
	}

	 utilities.SuccessResponse(&returnData,nil)
	 return


}