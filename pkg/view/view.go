package view

import (
	"net/http"
	"net/url"

	"github.com/saspallow/gonews/pkg/model"
)

type IndexData struct {
	List []*model.News
}

// Index Renders index View
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// News Renders news View
func News(w http.ResponseWriter, data *model.News) {
	render(tpNews, w, data)
}

type AdminLoginData struct {
	Flash url.Values
}

// AdminLogin renders Admin Login View
func AdminLogin(w http.ResponseWriter, data *AdminLoginData) {
	render(tpAdminLogin, w, data)
	data.Flash.Del("errors")
}

// AdminRegister renders Admin Register View
func AdminRegister(w http.ResponseWriter, data *IndexData) {
	render(tpAdminRegister, w, data)
}

type AdminListData struct {
	List []*model.News
}

// AdminList renders Admin Login View
func AdminList(w http.ResponseWriter, data *AdminListData) {
	render(tpAdminList, w, data)
}

// AdminCreate renders Admin Login View
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}

// AdminEdit renders Admin Login View
func AdminEdit(w http.ResponseWriter, data *model.News) {
	render(tpAdminEdit, w, data)
}
