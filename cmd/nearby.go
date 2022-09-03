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
)

// nearbyCmd represents the nearby command
var nearbyCmd = &cobra.Command{
	Use:   "nearby",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		lat, _ := cmd.Flags().GetFloat64("lat")
		lon, _ := cmd.Flags().GetFloat64("lon")
		fmt.Sprintf("Search area nearby: %f %f\n", lat, lon)
		println()
		areasResponse, err := client.SearchAreasByLatLong(lat, lon)
		if err == nil {
			areas := areasResponse.Areas
			tbl := table.New("Count", "Id")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, area := range areas {
				tbl.AddRow(area.Count, area.Id)
			}
			tbl.Print()
		}
		return err
	},
}

func init() {
	searchCmd.AddCommand(nearbyCmd)

	nearbyCmd.Flags().Float64("lat", -33.9249, "Latitude")
	nearbyCmd.Flags().Float64("lon", 18.4241, "Longitude")
}
