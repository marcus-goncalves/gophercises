package controllers

import api "photo_gallery.com/v1/server_side_app"

type Static struct {
	HomeView    *api.View
	ContactView *api.View
	FaqView     *api.View
}

func NewStatic() *Static {
	return &Static{
		HomeView:    api.NewView("main", "static/home"),
		ContactView: api.NewView("main", "static/contact"),
		FaqView:     api.NewView("main", "static/faq"),
	}
}
