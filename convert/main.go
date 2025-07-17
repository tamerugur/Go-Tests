package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	convertCmd := flag.NewFlagSet("convert", flag.ExitOnError)
	// define flags for fahrenheit and celsius, -f and -c, default to 0.0
	// by default it returns a pointer to the flag
	// ":=" is short variable declaration, it can only be used inside functions and type is inferred from right-hand side
	f := convertCmd.Float64("f", 0.0, "Fahrenheit")
	c := convertCmd.Float64("c", 0.0, "Celsius")

	if len(os.Args) < 2 {
		fmt.Println("please use subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "convert":
		convertCmd.Parse(os.Args[2:])
		if convertCmd.Parsed() {
			if *f != 0 {
				fmt.Println("°C:", (*f-32)*5/9)
			} else if *c != 0 {
				fmt.Println("°F:", (*c*9/5)+32)
			}
		}
	default:
		fmt.Println("unknown subcommand:", os.Args[1])
		os.Exit(1)
	}
}