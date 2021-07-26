package Models

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Config"
	"fmt"
)
//GetAllUsers Fetch all user data
func GetAllSubjects(user *[]Subject) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}
//CreateUser ... Insert New data
func CreateSubject(user *Subject) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetSubjectByID(user *Subject, id string) (err error) {
	if err = Config.DB.Where("ID = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
//UpdateUser ... Update user
func UpdateSubject(user *Subject, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}
//DeleteUser ... Delete user
func DeleteSubject(user *Subject, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}

