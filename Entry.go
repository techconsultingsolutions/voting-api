// This is entry point.
package main

import ( 
	"fmt"
	"github.com/techconsultingsolutions/voting-api/models"
)
/*
This is multiline
comment.
*/
func main() {
	fmt.Println("Hi");
	u := models.User{
		ID: 2,
	}

	fmt.Println(u);
}