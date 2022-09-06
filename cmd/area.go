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
	"fmt"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// areaCmd represents the area command
var areaCmd = &cobra.Command{
	Use:   "area",
	Short: "Retrieve information for your area",
	Long: `Retrieve information for your area

Ref: https://documenter.getpostman.com/view/1296288/UzQuNk3E#1881472b-c959-4259-b574-177feb5e0cda`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id := viper.GetString("id")
		println("Retrieving information for area:", id)
		println()
		areaResponse, err := client.SearchArea(id)
		if err == nil {
			info := areaResponse.Info
			tbl := table.New("Name", "Region")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			tbl.AddRow(info.Name, info.Region)
			tbl.Print()
			println()
			println("Events")
			tbl = table.New("Note", "Start", "End")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, event := range areaResponse.Events {
				tbl.AddRow(event.Note, event.Start, event.End)
			}
			tbl.Print()
			println()
			println("Schedule via ", areaResponse.Schedule.Source)
			tbl = table.New("Date", "Day of week", "Stages")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, day := range areaResponse.Schedule.Days {
				stages := day.Stages
				if len(stages) == 0 {
					println(0)
					tbl.AddRow(day.Date, day.Name, "")
				} else {
				    tbl.AddRow(day.Date, day.Name, stageString(stages[0]))
				    if len(stages) > 1 {
						for _, stage := range stages[1:] {
							tbl.AddRow("", "", stageString(stage))
						}
					}
				}
			}
			tbl.Print()
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(areaCmd)

	areaCmd.Flags().String("id", "", "Area id")
	if err := viper.BindPFlag("id", areaCmd.Flags().Lookup("id")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}
}

func stageString(stage []string) string {
	if len(stage) == 0 {
		return ""
	} else if len(stage) == 1 {
		return stage[0]
	} else {
		return fmt.Sprintf("%s %s", stage[0], stage[1])
	}
}