package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Select day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt dailed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

}
