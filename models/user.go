package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
)



type User struct {
	Id             int       `json:"id";orm:"column(id);auto"`
	FirstName      string       `json:"firstName";orm:"column(first_name)"`
	LastName       string    `json:"lastName";orm:"column(last_name)"`
}


func (u *User) TableName() string {
	return "user"
}


func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
}


func AddUser(user *User)error{
	o := orm.NewOrm()
    _,err:=o.Insert(user)
	return err

}

func GetAllUser() (users []User, err error) {
	o := orm.NewOrm()
	users = []User{}
	_, err = o.QueryTable(new(User)).RelatedSel().All(&users)
	fmt.Println(users)
	return users, err
}

func GetUserById(id int) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{}
	err = o.QueryTable(new(User)).Filter("id", id).RelatedSel().One(user)
	if err == nil {
		return user, nil
	}
	return nil, err
}




func UpdateUser(u *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: u.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	user := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&user); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}