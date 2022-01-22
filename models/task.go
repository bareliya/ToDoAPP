package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)



type Task struct {
	Id            int       `json:"id"; orm:"column(id);auto"`
	UserId        int       `json:"userId"; orm:"column(user_id)"`
	TaskName      string    `json:"taskname";orm:"column(task_name)"`
	Description   string    `json:"description";orm:"column(description)"`
	Status        string    `json:"status";orm:"column(status)"`
	DueDate       time.Time `json:"dueDate"; orm:"column(due_date);"`
}
	


func (t *Task) TableName() string {
	return "task"
}


func init() {
	// Need to register model in init
	orm.RegisterModel(new(Task))
}


func AddTask(task *Task)error{
	o := orm.NewOrm()
    _,err:=o.Insert(task)
	return err

}

func GetAllTask() (tasks []Task, err error) {
	o := orm.NewOrm()
	tasks = []Task{}
	_, err = o.QueryTable(new(Task)).RelatedSel().All(&tasks)
	fmt.Println(tasks)
	return tasks, err
}

func GetTaskById(id int) (task *Task, err error) {
	o := orm.NewOrm()
	task = &Task{}
	err = o.QueryTable(new(Task)).Filter("id", id).RelatedSel().One(task)
	if err == nil {
		return task, nil
	}
	return nil, err
}

func GetAllTaskByUserid(userId int) (v []Task, err error) {
	o := orm.NewOrm()
	v = []Task{}
	_, err = o.QueryTable(new(Task)).Filter("user_id", userId).RelatedSel().OrderBy("due_date").All(&v)
	return v, err
}

func GetAllTaskByUserIdAndStatus(userId int, status string) (v []Task, err error) {
	o := orm.NewOrm()
	v = []Task{}
	_, err = o.QueryTable(new(Task)).Filter("user_id", userId).Filter("status", status).RelatedSel().OrderBy("due_date").All(&v)
	return v, err
}


func UpdateTask(m *Task) (err error) {
	o := orm.NewOrm()
	v := Task{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteTask(id int) (err error) {
	o := orm.NewOrm()
	task := Task{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&task); err == nil {
		var num int64
		if num, err = o.Delete(&Task{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

