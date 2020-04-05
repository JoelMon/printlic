package cmd

var inputCheckTestCases = []struct {
	description   string
	input         string
	expected      string
	errorExpected bool
}{
	{
		description:   "Empty First Name",
		input:         "",
		errorExpected: true,
	},
}
