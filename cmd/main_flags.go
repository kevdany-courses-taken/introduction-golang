package main

// Package main shows how to use flags in golang.

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// )

// var beers = map[string]string {
// 	"pilsner": "A classic German pale lager.",
// 	"hefeweizen": "German wheat beer with a rich, malty",
// }

// Flags of "os" library
// Run function: $ go run main.go beers
//
// 	func main() {
// 		param := os.Args[1] // Start on the position 1, why position 0 is the application
// 		if param == "beers" {
// 			fmt.Println(beers)
// 		}
// 		fmt.Println(param)
// 	}

// Flags of "flag" library, standar library
// Run function: $ go run main.go --beers
//
// func main(){
//		// Params
//		// 1- name of flag
// 		// 2- value for default this flag
// 		// 3- description of flag, show help
// 		b := flag.Bool("beers", false, "print beers") // return a pointer
// 		flag.Parse() // read flags of console, else return value default
// 		if *b {
// 			fmt.Println(beers)
// 		}
//  }

// Flags of "flags" library, with flagset
// Run function: $ go run main.go beers --id=pilsner
//
// func main(){
// 		beersCmd := flag.NewFlagSet("beers", flag.ExitOnError)
// 		flag.Parse()

// 		if flag.NArg() == 0 {
// 			flag.PrintDefaults()
// 			os.Exit(2)
// 		}

// 		switch flag.Arg(0) {
// 		case "beers":
// 			ID := beersCmd.String("id", "", "find by ID")
// 			beersCmd.Parse(os.Args[2:])

// 			if *ID != "" {
// 				fmt.Println(beers[*ID])
// 			} else {
// 				fmt.Println(beers)
// 			}
// 		}
// }
