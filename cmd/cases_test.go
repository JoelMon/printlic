package cmd

var inputCheckTestCases = []struct {
	description   string
	firstName     string
	lastName      string
	middleName    string
	birth         string
	sex           string
	errorExpected bool
}{
	{
		description:   "Empty First Name",
		firstName:     "",
		lastName:      "Robert",
		middleName:    "James",
		birth:         "09341995",
		sex:           "m",
		errorExpected: true,
	},
	{
		description:   "Empty Last Name",
		firstName:     "Jose",
		lastName:      "",
		middleName:    "James",
		birth:         "09341995",
		sex:           "m",
		errorExpected: true,
	},
	{
		description:   "Empty Middle Name",
		firstName:     "Jose",
		lastName:      "Rojas",
		middleName:    "",
		birth:         "09341995",
		sex:           "m",
		errorExpected: false,
	},
	{
		description:   "Empty Birth",
		firstName:     "Jose",
		lastName:      "Rojas",
		middleName:    "Miras",
		birth:         "",
		sex:           "m",
		errorExpected: true,
	},
	{
		description:   "Empty Sex",
		firstName:     "Jose",
		lastName:      "Rojas",
		middleName:    "Miras",
		birth:         "03231988",
		sex:           "",
		errorExpected: true,
	},
}

var eFirstNameTestCases = []struct {
	description   string
	firstName     string
	middleName    string
	expected      string
	errorExpected bool
}{
	{
		description: "Encode Wilma",
		firstName:   "Wilma",
		middleName:  "",
		expected:    "920",
	},
	{
		description: "Encode Betty",
		firstName:   "Betty",
		middleName:  "",
		expected:    "080",
	},
	{
		description: "Encode William Andrew",
		firstName:   "William",
		middleName:  "Andrew",
		expected:    "921",
	},
	{
		description: "Encode Andrew Jones",
		firstName:   "Andrew",
		middleName:  "Jones",
		expected:    "010",
	},
	{
		description: "Encode Joel",
		firstName:   "Joel",
		middleName:  "",
		expected:    "420",
	},
}
var padTestCases = []struct {
	description   string
	codedName     int
	expected      string
	errorExpected bool
}{
	{
		description: "5 => 005",
		codedName:   5,
		expected:    "005",
	},
	{
		description: "01 => 001",
		codedName:   01,
		expected:    "001",
	},
	{
		description: "50 => 050",
		codedName:   50,
		expected:    "050",
	},
	{
		description: "100 => 100",
		codedName:   100,
		expected:    "100",
	},
	{
		description: "999 => 999",
		codedName:   999,
		expected:    "999",
	},
}
