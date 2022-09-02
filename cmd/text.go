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
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		areasResponse, err := client.SearchAreasByText(cmd.Flag("value").Value.String())
		if err == nil {
			areas := areasResponse.Areas
			tbl := table.New("Id", "Name", "Region")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, area := range areas {
				tbl.AddRow(area.Id, area.Name, area.Region)
			}
			tbl.Print()
		}
		return err
	},
}

func init() {
	searchCmd.AddCommand(textCmd)

	textCmd.Flags().String("value", "v", "Search text" )

}
