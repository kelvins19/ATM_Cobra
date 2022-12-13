/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from current user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			user := viper.Get("username")

			if user == "" {
				fmt.Print("There is no user logged in. Please login first\n")
			} else {
				fmt.Printf("Goodbye %v\n", user)

				viper.Set("username", "")
				viper.WriteConfig()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
