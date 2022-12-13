/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kelvins19/ATM/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login the user",
	Long:  `Login command to login the user based on the username`,
	Run: func(cmd *cobra.Command, args []string) {
		checkCurrentUser := viper.Get("username")

		if len(args) < 1 {
			fmt.Print("You should have at least one username\n")
		}

		if len(args) > 1 {
			fmt.Print("You can only have one username\n")
		}

		if len(args) == 1 {
			if checkCurrentUser != "" {
				if args[0] == fmt.Sprintf("%v", checkCurrentUser) {
					fmt.Printf("You have already logged in as %v. Please continue your activity\n", checkCurrentUser)
				} else {
					fmt.Printf("You have already been logged in as %v. Please logout first.\n", checkCurrentUser)
				}
			} else {
				user, _ := database.FindOrCreateUser(args[0])
				viper.Set("username", user.Username)
				viper.WriteConfig()
				fmt.Printf("Hello %v \n", user.Username)
				fmt.Printf("Your balance is $%v\n", user.Balance)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
