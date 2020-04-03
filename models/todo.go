package models

import "github.com/jshzhj/ginframe/pkg/mysql"

type Todos struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//对Todos表的增删改查都写在这个文件里
func CreateATodo(todo *Todos) (err error) {
	if _, err = mysql.DB.Insert(todo); err != nil {
		return
	}
	return nil
}

func GetTodoList(todo *[]Todos) (err error) {
	if err = mysql.DB.Find(todo); err != nil {
		return
	}
	return nil
}

func GetATodo(id string, todo *Todos) (err error) {
	if _, err = mysql.DB.Where("id = ?", id).Get(todo); err != nil {
		return
	}
	return nil
}

func UpdateATodo(id string, todo *Todos) (err error) {
	if _, err = mysql.DB.Where("id = ?", id).Cols( "status").Update(todo); err != nil {
		return
	}
	return nil
}

func DeleteATodo(id string, todo *Todos) (err error) {
	if _, err = mysql.DB.Where("id = ?", id).Delete(todo); err != nil {
		return
	}
	return nil
}
