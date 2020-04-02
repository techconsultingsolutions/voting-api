// This is entry point.
package mainx

import ( "fmt" )

/*
This is multiline
comment.
*/
func mainx() {
	fmt.Println("Declare variable in go!!!!");

	//These are value type data type.
	var i int;
	i = 40;
	fmt.Println(i);

	var f float32 = 1.2;
	fmt.Println(f);

	//Implicity initialization
	myName := "Hemant"
	fmt.Println(myName);
	b := true;
	fmt.Println(b);

	//Complex datatype
	c := complex(4,5);
	fmt.Println(c);

	//Multiple assignment

	r , img := real(c) , imag(c);
	fmt.Println(r,img);

	//Pinter datatype.
	var firstNamePtr * string = new (string);
	*firstNamePtr = "This is Value at Pointer";
	fmt.Println(*firstNamePtr);
	
	//Pointer arithmatic is not allowed in GO 

	ptrToPtr := &firstNamePtr;

	fmt.Println(ptrToPtr,*ptrToPtr,**ptrToPtr);


	//Constants
	const pi = 3.14;
	fmt.Println(pi);

	const g int = 10;
	fmt.Println(g);

	//Constant expressions.


}