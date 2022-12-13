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

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Command to transfer balance to another username",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		user := viper.Get("username")
		username := fmt.Sprintf("%v", user)

		if len(args) == 2 {
			// check if user logged in or not
			if user != "" {
				userData, _ := database.FindUser(username)

				transferUsername := fmt.Sprintf("%v", args[0])
				transferUserData, _ := database.FindUser(transferUsername)

				// check if transferred user is exist
				if transferUserData != nil {
					transferBalance, _ := strconv.ParseFloat(args[1], 64)

					// Check if the transfer username is same with logged in username
					if username == transferUsername {
						fmt.Printf("You cannot transfer to your own username\n")
					} else {
						// Check if current user balance is less than transferBalance
						if userData.Balance < transferBalance {
							owedBalance := transferBalance - userData.Balance

							transferUserData.Balance += userData.Balance
							userData.Balance -= userData.Balance

							database.UpdateUser(transferUserData)
							database.UpdateUser(userData)

							fmt.Printf("Transferred $%v to %v\n", transferBalance, transferUsername)
							fmt.Printf("Your balance is $%v\n", userData.Balance)
							fmt.Printf("Owed $%v to %v\n", owedBalance, transferUserData.Username)

						} else {
							userData.Balance -= transferBalance

							transferUserData.Balance += transferBalance

							database.UpdateUser(userData)
							database.UpdateUser(transferUserData)

							fmt.Printf("Transferred $%v to %v\n", transferBalance, transferUsername)
							fmt.Printf("Your balance is $%v\n", userData.Balance)
						}
					}

				} else {
					fmt.Printf("User %v not found\n", transferUsername)
				}

			} else {
				fmt.Printf("Please login first.\n")
			}
		} else {
			fmt.Printf("Please make sure to pass correct arguments.\n")
		}

	},
}

func init() {
	rootCmd.AddCommand(transferCmd)
}
