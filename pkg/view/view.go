package view

import "net/http"

// Index Renders index View
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}

// AdminLogin renders Admin Login View
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminList renders Admin Login View
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminList, w, data)
}

// AdminCreate renders Admin Login View
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}

// AdminEdit renders Admin Login View
func AdminEdit(w http.ResponseWriter, data interface{}) {
	render(tpAdminEdit, w, data)
}
