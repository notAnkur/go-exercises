package main

import (
	"flag"
	"fmt"
	// "os"
)

func main() {
	// extract commandline arguments from terminal
	// argsBuf := make([]byte, 100)
	var nFlag = flag.Int("n", 1234, "test flag")
	flag.Parse();
	// n, err := os.Stdin.Read(argsBuf)
	// if err != nil {
	// 	fmt.Println("Error reading stdin", err)
	// }
	// fmt.Printf("Read %d bytes", n)
	fmt.Println(*nFlag);
}
