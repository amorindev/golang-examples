package handler

import (
	"html/template"
	"net/http"
)

type User struct {
	Name     string
	Email string
	Avatar   string
	LastName string
	IsActive bool
}

var user = User{
	Name:     "",
	Email: "jhon@test.com",
	Avatar: "https://plus.unsplash.com/premium_photo-1689568126014-06fea9d5d341?q=80&w=870&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	LastName: "Doe",
	IsActive: true,
}

type Task struct {
	ID          string
	Text        string
	IsCompleted bool
}

var tasks = []*Task{
	{
		ID:          "1",
		Text:        "first task",
		IsCompleted: true,
	},
	{
		ID:          "2",
		Text:        "second article",
		IsCompleted: false,
	},
	{
		ID:          "3",
		Text:        "third article",
		IsCompleted: true,
	},
}

func HelloWorldPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"web/templates/hello_world.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		User  *User
		Tasks []*Task
	}{
		Tasks: tasks,
		User:  &user,
	}

	err = ts.ExecuteTemplate(w, "hello_world", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
