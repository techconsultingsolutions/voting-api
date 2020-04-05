package controllers

import (
	"fmt"
	"net/http"
)

/*
RegisterControllers
*/
func RegisterControllers() {
	//uc := newUserController()
	var uc userController
	fmt.Println(uc)
	http.Handle("/users", uc)
	http.Handle("/users/", uc)
}
