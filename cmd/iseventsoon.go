/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// iseventsoonCmd represents the iseventsoon command
var iseventsoonCmd = &cobra.Command{
	Use:   "is-event-soon",
	Short: "Returns zero exit code if loadsheeding event is soon, non-zero if not or could not determine due to error",
	Long: `Returns zero exit code if loadsheeding event is soon, non-zero if not or could not determine due to error`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		id := viper.GetString("id")
		suggestShutdownTime, err := cmd.Flags().GetDuration("suggest-shutdown-time")
		if err != nil {
			return err
		}
		println("Retrieving information for area:", id)
		println()
		simulateEvent, _ := cmd.Flags().GetString("simulate-event")
		areaResponse, err := client.SearchArea(id, simulateEvent )
		if err == nil {
			if len(areaResponse.Events) > 0 {
				nextEvent := areaResponse.Events[0]
				timeTillNextEventStarts := time.Until(nextEvent.Start)
				fmt.Printf("%v will start in %v at %v\n", nextEvent.Note, timeTillNextEventStarts, nextEvent.Start)
				println()
				if timeTillNextEventStarts <= suggestShutdownTime {
					println("Recommend shutting down")
				} else {
					println("Do not recommend shutting down")
					os.Exit(1)
				}
			}
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(iseventsoonCmd)

	iseventsoonCmd.Flags().String("id", "", "Area id")
	iseventsoonCmd.Flags().String("simulate-event", "", "Simulate an event (current/future)")
	iseventsoonCmd.Flags().Duration("suggest-shutdown-time", 1 * time.Hour, "Suggest shutdown if event is in less than or equal to this time")

	if err := viper.BindPFlag("id", areaCmd.Flags().Lookup("id")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}
}
