//Package helloUser shows the greetings
package FirstPackage

//Hello is a global public variable
var Hello = "Hello"

//PrintHello is a global public function
func PrintHello(name string) string {
	return Hello + "-" + name
}
