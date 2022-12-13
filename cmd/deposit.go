/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/kelvins19/ATM/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// depositCmd represents the deposit command
var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Command to deposit balance to logged in user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		user := viper.Get("username")
		username := fmt.Sprintf("%v", user)

		if user != "" {
			userData, _ := database.FindUser(username)
			depositBalance, _ := strconv.ParseFloat(args[0], 64)

			if depositBalance > 0 {
				userData.Balance += depositBalance

				database.UpdateUser(userData)
				fmt.Printf("Your balance is %v ", userData.Balance)
			} else {
				fmt.Printf("Please deposit more than 0")
			}
		} else {
			fmt.Printf("Please login first.")
		}
	},
}

func init() {
	rootCmd.AddCommand(depositCmd)
}
