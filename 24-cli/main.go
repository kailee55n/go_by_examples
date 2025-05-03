package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("=== CLI Examples in Go ===")
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go basic")
		fmt.Println("  go run main.go flags -name Alice -age 25 -verbose")
		fmt.Println("  go run main.go cobra get users --all")
		return
	}

	// Dispatch based on first argument
	switch os.Args[1] {
	case "basic":
		demoBasicArgs()
	case "flags":
		demoStandardFlags(os.Args[2:])
	case "cobra":
		demoCobraCommands(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

// Basic os.Args demo
func demoBasicArgs() {
	fmt.Println("\n=== Basic Arguments ===")
	args := os.Args
	fmt.Println("Program name:", args[0])
	fmt.Println("Arguments:", args[1:])
}

// Flag package demo
func demoStandardFlags(args []string) {
	fmt.Println("\n=== Standard Flags ===")

	fs := flag.NewFlagSet("flags", flag.ExitOnError)
	name := fs.String("name", "World", "Name to greet")
	age := fs.Int("age", 0, "Age of the person")
	verbose := fs.Bool("verbose", false, "Enable verbose output")
	color := fs.String("color", "blue", "Favorite color")

	fs.Parse(args)

	fmt.Printf("Hello, %s!\n", *name)
	if *age > 0 {
		fmt.Printf("Age: %d\n", *age)
	}
	if *verbose {
		fmt.Println("Verbose mode enabled")
		fmt.Printf("Favorite color: %s\n", *color)
		fmt.Printf("Remaining args: %v\n", fs.Args())
	}
}

// Cobra command demo
func demoCobraCommands(args []string) {
	var rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "Sample CLI with Cobra",
	}

	var getCmd = &cobra.Command{
		Use:   "get [resource]",
		Short: "Get a resource",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Getting resource: %s\n", args[0])
			all, _ := cmd.Flags().GetBool("all")
			if all {
				fmt.Println("Fetching all items")
			}
		},
	}
	getCmd.Flags().BoolP("all", "a", false, "Get all resources")

	var createCmd = &cobra.Command{
		Use:   "create [resource]",
		Short: "Create a resource",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Creating resource: %s\n", args[0])
			force, _ := cmd.Flags().GetBool("force")
			if force {
				fmt.Println("Force flag enabled")
			}
		},
	}
	createCmd.Flags().BoolP("force", "f", false, "Force creation")

	rootCmd.AddCommand(getCmd, createCmd)
	rootCmd.SetArgs(args) // Provide args explicitly
	rootCmd.Execute()
}
