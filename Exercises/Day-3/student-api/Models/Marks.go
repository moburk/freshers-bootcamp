package Models

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Config"
	"fmt"
)
//GetAllUsers Fetch all user data
func GetAllMarks(user *[]Marks) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}
//CreateUser ... Insert New data
func CreateMarks(user *Marks) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil

}

func GetMarksByID(marks *Marks, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(marks).Error; err != nil {
		return err
	}
	return nil
}

//GetMarksByStudentID ... Fetch only one user by Id
func GetMarksByStudentID(marks *Marks, id string) (err error) {
	if err = Config.DB.Where("StudentID = ?", id).Find(marks).Error; err != nil {
		return err
	}
	return nil
}

func GetMarksBySubjectID(marks *Marks, id string) (err error) {
	if err = Config.DB.Where("SubjectID = ?", id).Find(marks).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateMarks(user *Marks, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}
//DeleteUser ... Delete user
func DeleteMarks(user *Marks, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
