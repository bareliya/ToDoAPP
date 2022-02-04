# ToDoAPP

### This is a Simple ToDo Backend Code 
### One Instance of code i have alreday deployed on the lambda and created API GAteway Endpoints That can be be found in the below postman collection
https://www.getpostman.com/collections/526963b4efd334bca962


# for Running Code in the local system follow steps written below
1. clone the repository
2. set the mysql credential in conf->env.yml file ( make sure runmode shoud be dev )
3. create databse and tables, please refer mysql->dbSchema.sql ( just copy paste those command in your workbench nothing more :) )
4. run the code and follow the below postman collection 

https://www.getpostman.com/collections/e1a846af5ec01f9e0172



### --------------------------------------------------------------------------------------------------------------
## Code Overview

we have two table is our database tododb, one table is task, to store details of each task while other is user to store user details.<br>
Schema is given below

``` sql
--
-- Table structure for table `User`
--
create table user(
  id int auto_increment not null,
  first_name varchar(50) not null,
  last_name varchar(50) not null,
   primary key(id)
  

);

--
-- Table structure for table `Task`
--

create table task(
  id int auto_increment not null,
  user_id  int not null,
  task_name  varchar(100) not null,
  description text,
  status varchar(50) not null,
  due_date datetime NOT NULL DEFAULT current_timestamp(),
  primary key(id)
);
 
```

## Endpoint
 1. add user : <br>
    ### POST  /todo/user <br>
     Request Body:
     ``` json
     {
    "firstName": "hardic",
    "lastName": "pandya"
     }
     ```
this end point will add the the user in the user table <br>
Response Body:
  ``` json
          
    {
     "code": 200,
     "msg": "success",
     "model": {
         "id": 7,
         "firstName": "hardic",
         "lastName": "pandya"
          }
     }
   ```
you will get user id on adding a user(in this case user id is 7)<br>

2. get all users: <br>
   ### Get /todo/user<br>
  it will basically fetch all the user from the table<br>
  Resopnse Body
  ``` json
  {
    "code": 200,
    "msg": "success",
    "model": [
        {
            "id": 1,
            "firstName": "fname",
            "lastName": "lname"
        },
        {
            "id": 2,
            "firstName": "Rahul",
            "lastName": "Bareliya"
        },
        {
            "id": 3,
            "firstName": "test",
            "lastName": "user"
        },
        {
            "id": 4,
            "firstName": "virat",
            "lastName": "kohli"
        },
        {
            "id": 5,
            "firstName": "rohit",
            "lastName": "sharma"
        },
        {
            "id": 6,
            "firstName": "hardic",
            "lastName": "pandya"
        },
        {
            "id": 7,
            "firstName": "hardic",
            "lastName": "pandya"
        }
    ]
}
```

3. get a user by user id <br>
   ### Get todo/user?id=1<br>
   it will filter the table by user id and will return the details corrospondig to the userid<br>
   Response Body
   ``` json
      {
    "code": 200,
    "msg": "success",
    "model": {
        "id": 1,
        "firstName": "fname",
        "lastName": "lname"
       }
     }
     ```


4.  Add task  <br>
    ### Post todo/tsak<br>
    
   Request Body:
   ``` json
   {
    "userId": 6,
    "taskName": "BackEnd",
    "description": "an incredible task 2",
    "dueDate": "2022-02-23T17:58:53+05:30"
    }
```
task are added corrosponding to user <br>
so we need user id in request body (int above case is ,6) if we pass invalid user id we will get error accordingly
Response Body:
``` json
{
    "code": 200,
    "msg": "success",
    "model": {
        "id": 10,
        "userId": 6,
        "taskname": "BackEnd",
        "description": "an incredible task 2",
        "status": "pending",
        "dueDate": "2022-02-23T17:58:53+05:30"
    }
}
```

5. Update Task. <br>
   ### PUT todo/task<br>
    Request Body:
    ``` json
    {
      "id": 1,
        "taskname": "updated Task Name",
        "description": "an incredible task 2 updated",
        "dueDate": "2022-02-23T17:58:53+05:30"
   }
   ```
In this Case we must give a valid task id to update(in this case it is 1) other wise we will get error accordingly

Response Body:
``` json
{
    "code": 200,
    "msg": "success",
    "model": {
        "id": 1,
        "userId": 0,
        "taskname": "updated Task Name",
        "description": "an incredible task 2 updated",
        "status": "pending",
        "dueDate": "2022-02-23T17:58:53+05:30"
    }
}
```

6. Change Status <br> 
   ### PUT todo/task<br>
   Requst Body :
   ``` json 
   {
    "id": 7,
    "status": "completed"
   }
    ```
Agsin we need a valid task id to update the status<br>
response Body:
``` json
{
    "code": 200,
    "msg": "success",
    "model": {
        "id": 7,
        "userId": 0,
        "taskname": "",
        "description": "an incredible task 2",
        "status": "completed",
        "dueDate": "2022-02-23T17:58:53+05:30"
    }
  }
```
7. delete task  <br>
   ### DELETE todo/task<br>
   request body:
   ``` json
   {
    "id": 6

   }
    ```
we need to pass just task id in the request body<br>
response body:
``` json
{
    "code": 200,
    "msg": "success",
    "model": null
  }
```
8. get all task <br> 
   ### GET todo/task<br>
   Response Body:
   ``` json
   {
    "code": 200,
    "msg": "success",
    "model": [
        {
            "id": 1,
            "userId": 0,
            "taskname": "updated Task Name",
            "description": "an incredible task 2 updated",
            "status": "pending",
            "dueDate": "2022-02-23T17:58:53+05:30"
        },
        {
            "id": 2,
            "userId": 1,
            "taskname": "",
            "description": "an incredible task 2",
            "status": "completed",
            "dueDate": "2022-02-23T17:58:53+05:30"
        },
        {
            "id": 3,
            "userId": 1,
            "taskname": " TaskName",
            "description": "an incredible task 2",
            "status": "pending",
            "dueDate": "2022-02-23T17:58:53+05:30"
        },
        {
            "id": 4,
            "userId": 3,
            "taskname": "BournVita",
            "description": "an incredible task 2",
            "status": "pending",
            "dueDate": "2022-02-23T17:58:53+05:30"
        },
        {
            "id": 7,
            "userId": 0,
            "taskname": "",
            "description": "an incredible task 2",
            "status": "completed",
            "dueDate": "2022-02-23T17:58:53+05:30"
        },
        {
            "id": 8,
            "userId": 5,
            "taskname": "FrontEnd",
            "description": "an incredible task 2",
            "status": "pending",
            "dueDate": "2022-02-23T17:58:53+05:30"
        },
        {
            "id": 9,
            "userId": 6,
            "taskname": "BackEnd",
            "description": "an incredible task 2",
            "status": "pending",
            "dueDate": "2022-02-23T17:58:53+05:30"
        },
        {
            "id": 10,
            "userId": 6,
            "taskname": "BackEnd",
            "description": "an incredible task 2",
            "status": "pending",
            "dueDate": "2022-02-23T17:58:53+05:30"
        }
       ]
     }
   ```
   
   
9. get the task by user id <br> 
   ### GET todo/tsak?userId=1 
  ``` json
  {
    "code": 200,
    "msg": "success",
    "model": [
        {
            "id": 2,
            "userId": 1,
            "taskname": "",
            "description": "an incredible task 2",
            "status": "completed",
            "dueDate": "2022-02-23T17:58:53+05:30"
        },
        {
            "id": 3,
            "userId": 1,
            "taskname": " TaskName",
            "description": "an incredible task 2",
            "status": "pending",
            "dueDate": "2022-02-23T17:58:53+05:30"
        }
    ]
  }
```
we will get the all the task created by the user <br>
10. get the task by task id. <br> 
    ### GET task/todo?id=1<br>
   response body:
   ``` json 
   {
    "code": 200,
    "msg": "success",
    "model": {
        "id": 1,
        "userId": 0,
        "taskname": "updated Task Name",
        "description": "an incredible task 2 updated",
        "status": "pending",
        "dueDate": "2022-02-23T17:58:53+05:30"
    }
}
```
we will get the task corrosponding to the id<br>

11. get the task by user is and status<br>
    ### GET todo/task?userId=1&status=pending
  reponse body:
  ``` json
  {
    "code": 200,
    "msg": "success",
    "model": [
        {
            "id": 3,
            "userId": 1,
            "taskname": " TaskName",
            "description": "an incredible task 2",
            "status": "pending",
            "dueDate": "2022-02-23T17:58:53+05:30"
        }
    ]
}
```
this will give us the task have the given status corrosponding to the user id
