package main 

import "os"
import "log"

func main() {
	
file, err := os.Create("sample.txt")

if err != nil {
	log.Fatal(err)
}
	
	file.WriteString("Hello, World!")
	file.Close

}