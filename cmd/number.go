/*
Copyright © 2020 Joel Montes de Oca

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var firstName string
var lastName string
var middleName string
var birth string
var sex string

// numberCmd represents the number command
var numberCmd = &cobra.Command{
	Use:   "number",
	Short: "Generate a Florida driver license number.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("number called")
		fmt.Println(firstName + lastName + middleName + birth + sex)
	},
}

func init() {
	rootCmd.AddCommand(numberCmd)

	numberCmd.Flags().StringVarP(&firstName, "first", "f", "", "First name (Required)")
	numberCmd.Flags().StringVarP(&lastName, "last", "l", "", "Last name -- wrap in \"\" if contains white space. (Required)")
	numberCmd.Flags().StringVarP(&middleName, "middle", "m", "", "Middle name -- Leave blank if N/A")
	numberCmd.Flags().StringVarP(&birth, "birth", "b", "", "Birth date: mmddyyyy (Required)")
	numberCmd.Flags().StringVarP(&sex, "sex", "s", "", "The person's sex: m or f (Required)")
	numberCmd.MarkFlagRequired("first")
	numberCmd.MarkFlagRequired("last")
	numberCmd.MarkFlagRequired("birth")
	numberCmd.MarkFlagRequired("sex")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// numberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// numberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func inputCheck(fn, ln, mn, bd, sex string) (string, error) {
	return "Joel", nil
}
