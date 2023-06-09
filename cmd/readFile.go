/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	assignments "github.com/91diego/platform-challenge/pkg"
	"github.com/spf13/cobra"
)

// readFileCmd represents the readFile command
var readFileCmd = &cobra.Command{
	Use:   "readFile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		shipmentFile, err := cmd.Flags().GetString("shipmentFile")
		if err != nil {
			log.Fatal(err)
		}
		driversFile, err := cmd.Flags().GetString("driverFile")
		if err != nil {
			log.Fatal(err)
		}

		driverAssignments, err := assignments.ShipmentAssignments(&shipmentFile, &driversFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(driverAssignments)
	},
}

func init() {
	rootCmd.AddCommand(readFileCmd)

	readFileCmd.Flags().StringP("shipmentFile", "s", "", "param to upload shimpement file")
	readFileCmd.Flags().StringP("driverFile", "d", "", "param to upload drivers file")
	rootCmd.MarkFlagRequired("shipmentfile")
	rootCmd.MarkFlagRequired("driverFile")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
