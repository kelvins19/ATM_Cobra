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

// withdrawCmd represents the withdraw command
var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "Command to withdraw balance from logged in user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		user := viper.Get("username")
		username := fmt.Sprintf("%v", user)

		if user != "" {
			userData, _ := database.FindUser(username)
			withdrawBalance, _ := strconv.ParseFloat(args[0], 64)

			if withdrawBalance <= userData.Balance {
				userData.Balance -= withdrawBalance

				database.UpdateUser(userData)
				fmt.Printf("Your balance is $%v \n", userData.Balance)
			} else {
				fmt.Printf("You don't have enough funds to do the withdraw.\n")
			}
		} else {
			fmt.Printf("Please login first.\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(withdrawCmd)
}
