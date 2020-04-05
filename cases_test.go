package cmd

var inputCheckTestCases = []struct {
	description   string
	fname         string
	lname         string
	mname         string
	birth         string
	sex           string
	expected      string
	errorExpected bool
}{
	{
		description:   "All empty fields",
		fname:         "",
		lname:         "",
		mname:         "",
		birth:         "",
		sex:           "",
		expected:      "1",
		errorExpected: true,
	},
}
