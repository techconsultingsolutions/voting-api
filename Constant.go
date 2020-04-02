// This is entry point.
package main

import ( "fmt" )

const pi = 3.14;

const (
	first = iota	
	second 
)

const (
	Employee = iota
	Manager = iota
)

/*
This is multiline
comment.
*/
func main() {
	fmt.Println("Declare variable in go!!!!");
	fmt.Println(pi);

	fmt.Println(first,second);

	fmt.Println(Manager);

	//iota and constant blocks

}