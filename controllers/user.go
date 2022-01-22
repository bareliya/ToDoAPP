package controllers

import(
	"ToDoAPP/utilities"
	"ToDoAPP/models"
	"github.com/spf13/cast"
)



func GetAllUserFunction()(returnData utilities.ResponseJSON){

	allData, err := models.GetAllUser()
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: No User found!")
	} else {
		utilities.SuccessResponse(&returnData, allData)
	}
	return

}


func GetUserByIdFunction(id int)(returnData utilities.ResponseJSON){

	data, err := models.GetUserById(id)
	if err != nil {
		utilities.ErrorResponse(&returnData, "user does not exist!")
	} else {
		utilities.SuccessResponse(&returnData, data)
	}
	return

}



func AddUserFunction(user models.User)(returnData utilities.ResponseJSON){

	if utilities.IsEmpty(user.FirstName) || utilities.IsEmpty(user.LastName){
		utilities.ErrorResponse(&returnData,"Error:First Name and Last Name Both are mandotory Fields")
		return 

	}


	err:= models.AddUser(&user)
	if err!=nil{
        utilities.ErrorResponse(&returnData,cast.ToString(err))
	}

	utilities.SuccessResponse(&returnData,user)
	return

}



func DeleteUserFunction(user models.User)(returnData utilities.ResponseJSON){

	if user.Id==0 {
		utilities.ErrorResponse(&returnData," Error: Id is Empty")

	}
   
    err:= models.DeleteUser(user.Id)
	if err!=nil{
        utilities.ErrorResponse(&returnData,"user does not exist")
	}


    utilities.SuccessResponse(&returnData,nil)
	return

}