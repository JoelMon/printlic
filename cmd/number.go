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
	"errors"
	"fmt"
	"strings"

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
	numberCmd.Flags().StringVarP(&birth, "birth", "b", "", "Birth year: yyyy (Required)")
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

func inputCheck(firstName, lastName, middleName, birthYear, sex string) error {

	if firstName == "" {
		return errors.New("Error: \"First Name\" was left empty")
	}
	if lastName == "" {
		return errors.New("Error: \"Last Name\" was left empty")
	}
	if birthYear == "" {
		return errors.New("Error: \"Birth year\" was left empty")
	}
	if sex == "" {
		return errors.New("Error: \"Sex\" was left empty")
	}

	return nil
}

func eFirstName(firstName, middleName string) (string, error) {

	var nameTable = map[string]int{
		"albert":    20,
		"alice":     20,
		"ann":       40,
		"anna":      40,
		"anne":      40,
		"annie":     40,
		"arthur":    40,
		"bernard":   80,
		"bette":     80,
		"bettie":    80,
		"betty":     80,
		"carl":      120,
		"catherine": 120,
		"charles":   140,
		"clara":     140,
		"donald":    180,
		"dorthy":    180,
		"edward":    220,
		"elizabeth": 220,
		"florence":  260,
		"frank":     260,
		"george":    300,
		"grace":     300,
		"harold":    340,
		"harriet":   340,
		"harry":     360,
		"hazel":     360,
		"helen":     380,
		"henry":     380,
		"james":     440,
		"jane":      440,
		"jayne":     440,
		"jean":      460,
		"joan":      480,
		"john":      460,
		"joseph":    480,
		"margaret":  560,
		"martin":    560,
		"marvin":    580,
		"mary":      580,
		"melvin":    600,
		"mildred":   600,
		"patricia":  680,
		"paul":      680,
		"richard":   740,
		"robert":    760,
		"ruby":      740,
		"ruth":      760,
		"thelma":    820,
		"thomas":    820,
		"walter":    900,
		"wanda":     900,
		"william":   920,
		"wilma":     920,
	}

	var initName = map[rune]int{
		'a': 0,
		'b': 60,
		'c': 100,
		'd': 160,
		'e': 200,
		'f': 240,
		'g': 280,
		'h': 320,
		'i': 400,
		'j': 420,
		'k': 500,
		'l': 520,
		'm': 540,
		'n': 620,
		'o': 640,
		'p': 660,
		'q': 700,
		'r': 720,
		's': 780,
		't': 800,
		'u': 840,
		'v': 860,
		'w': 880,
		'x': 940,
		'y': 960,
		'z': 980,
	}

	var middleTable = map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 14,
		'p': 15,
		'q': 15,
		'r': 16,
		's': 17,
		't': 18,
		'u': 18,
		'v': 18,
		'w': 19,
		'x': 19,
		'y': 19,
		'z': 19,
	}

	firstName = strings.ToLower(firstName)
	middleName = strings.ToLower(middleName)

	_, ok := nameTable[firstName]

	if ok == true && middleName == "" {
		return pad(nameTable[firstName]), nil
	}

	// TODO: Refactor. mInitial is repeated twice.
	firstInit := rune(firstName[0])
	if ok == true && middleName != "" {
		mInitial := rune(middleName[0])
		return pad(nameTable[firstName] + middleTable[mInitial]), nil
	} else if ok == true && middleName == "" {
		return pad(nameTable[firstName]), nil
	} else if ok == false && middleName == "" {
		return pad(initName[firstInit]), nil
	} else if ok == false && middleName != "" {
		mInitial := rune(middleName[0])
		return pad(initName[firstInit] + middleTable[mInitial]), nil
	}

	return "ok", errors.New("something went wrong")
}

func pad(intName int) string {
	codedName := fmt.Sprintf("%03d", intName)
	return codedName

}
