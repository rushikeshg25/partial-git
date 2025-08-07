package internal

import "fmt"

type Flags struct {
	Set   string
	Auth  bool
	Check bool
	Unset bool
}

func Run(flags Flags, args []string) {
	fmt.Println("=== User Input Summary ===")

	// Print command line arguments
	if len(args) > 0 {
		fmt.Printf("Arguments: %v\n", args)
	} else {
		fmt.Println("Arguments: none")
	}

	// Print flag values
	fmt.Printf("Set flag: %q\n", flags.Set)
	fmt.Printf("Auth flag: %t\n", flags.Auth)
	fmt.Printf("Check flag: %t\n", flags.Check)
	fmt.Printf("Unset flag: %t\n", flags.Unset)

	fmt.Println("========================")
}
