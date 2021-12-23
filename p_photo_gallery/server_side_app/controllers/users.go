package controllers

import (
	"net/http"

	api "photo_gallery.com/v1/server_side_app"
)

type Users struct {
	NewView *api.View
}

func NewUsers() *Users {
	return &Users{
		NewView: api.NewView("main", "server_side_app/views/users/new.html"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}