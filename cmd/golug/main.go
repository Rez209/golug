package main

import (
	"fmt"
	"os"

	"github.com/Rez209/golug/internal/generator"
	"github.com/spf13/cobra"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Cyan  = "\033[36m"
)

const logo = Cyan + `
  ____       _                 
 / ___| ___ | |   _   _  __ _  
| |  _ / _ \| |  | | | |/ _` + "`" + ` | 
| |_| | (_) | |__| |_| | (_| | 
 \____|\___/|_____\__,_|\__, | 
                        |___/  
` + Reset

func main() {
	var rootCmd = &cobra.Command{
		Use:   "golug",
		Short: "GoLug - Multi-language microservice generator",
		Long:  logo + "\nGoLug is a powerful CLI tool for rapid microservice architecture generation.",
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Shows the current version of GoLug",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Cyan + "GoLug v1.0.0" + Reset)
		},
	}

	var port string
	var lang string
	var module string

	var newCmd = &cobra.Command{
		Use:   "new [service_name]",
		Short: "Initializes a new microservice",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(logo)
				fmt.Println(Red + "❌ Error: provide a service name! Example: golug new my-api" + Reset)
				return
			}
			serviceName := args[0]

			actualModule := module
			if lang == "go" && actualModule == "" {
				actualModule = "github.com/user/" + serviceName
			}

			fmt.Printf(Cyan+"🚀 Initializing microservice: %s (Language: %s, Port: %s, Module: %s)..."+Reset+"\n", serviceName, lang, port, actualModule)

			err := generator.GenerateService(serviceName, port, lang, actualModule)
			if err != nil {
				fmt.Println(Red+"❌ Error:", err, Reset)
			}
		},
	}

	newCmd.Flags().StringVarP(&port, "port", "p", "8080", "Network port for the microservice")
	newCmd.Flags().StringVarP(&lang, "lang", "l", "go", "Programming language (go, python, cpp, js)")
	newCmd.Flags().StringVarP(&module, "module", "m", "", "Go module path (e.g., github.com/user/project) [Only for Go]")

	rootCmd.AddCommand(versionCmd, newCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(Red, err, Reset)
		os.Exit(1)
	}
}
