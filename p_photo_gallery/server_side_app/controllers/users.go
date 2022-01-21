package controllers

import (
	"fmt"
	"net/http"

	api "photo_gallery.com/v1/server_side_app"
	"photo_gallery.com/v1/server_side_app/models"
)

type Users struct {
	NewView *api.View
	us      *models.UserService
}

type SignUpForm struct {
	Name     string `schema: name`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: api.NewView("main", "users/new"),
		us:      us,
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignUpForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user := models.User{
		Name:  form.Name,
		Email: form.Email,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User: ", user)
}
