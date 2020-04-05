// This is entry point.
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/techconsultingsolutions/voting-api/controllers"
	"github.com/techconsultingsolutions/voting-api/models"
)

/*
This is multiline
comment.
*/
func main() {
	fmt.Println("Hi")
	u := models.User{
		ID: 2,
	}

	fmt.Println(u)

	port := 3000
	retry := 2

	startWebServer(port, retry)
	startWebServerWithShortInitSyntax(port, retry)

	sumOfNumbers := sum(1, 2)
	fmt.Println("Sum of two numbers is : ", sumOfNumbers)

}

func startWebServer(port int, numberOfRetry int) {
	fmt.Println("Starting webserver...")

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	fmt.Println("Server has started and listing on ...", port, numberOfRetry)
}

/*
Shorter initialization syntax.
*/
func startWebServerWithShortInitSyntax(port, numberOfRetry int) {
	fmt.Println("Starting webserver...")

	//Do something

	fmt.Println("Server has started and listing on ...", port, numberOfRetry)
}

func sum(a, b int) int {
	return a + b
}

func div(n, d int) (int, error) {
	r := n / d
	return r, errors.New("something went wrong...")
}
