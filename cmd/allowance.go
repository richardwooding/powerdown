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

// allowanceCmd represents the allowance command
var allowanceCmd = &cobra.Command{
	Use:   "allowance",
	Short: "Retrieves information about the allowance of your EskomSePush token",
	Long: `Retrieves information about the allowance of your EskomSePush token

See: https://documenter.getpostman.com/view/1296288/UzQuNk3E#10647b8e-c839-4d56-82a2-d9a406ae4f18"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		allowanceResponse, err := client.Allowance()
		if err == nil {
			allowance := allowanceResponse.Allowance
			tbl := table.New("Count", "Limit", "Type")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			tbl.AddRow(allowance.Count, allowance.Limit, allowance.Type)
			tbl.Print()
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(allowanceCmd)
}
