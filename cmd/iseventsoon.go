/*
Copyright © 2022 Richard Wooding richard.wooding@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
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
		if id == "" {
			return errors.New("No id specified")
		}
		println("Retrieving information for area:", id)
		println()
		simulateEvent, _ := cmd.Flags().GetString("simulate-event")
		areaResponse, err := sePushClient.SearchArea(id, simulateEvent )
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
			} else {
				println("Do not recommend shutting down")
				os.Exit(1)
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

	if err := viper.BindPFlag("id", iseventsoonCmd.Flags().Lookup("id")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}
}
