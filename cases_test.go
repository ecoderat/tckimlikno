package tckimlikno

var testCases = []struct {
	description string
	input       string
	ok          bool
}{
	{
		"first digit can not be zero",
		"01234567890",
		false,
	},
	{
		"single digit strings can not be valid",
		"1",
		false,
	},
	{
		"a single zero is invalid",
		"0",
		false,
	},
	{
		"valid strings with punctuation included become invalid",
		"180-714-701-10",
		false,
	},
	{
		"valid number can not be negative",
		"-01234567890",
		false,
	},
	{
		"using ascii value for doubled non-digit isn't allowed",
		":9876543210",
		false,
	},
	{
		"last digit invalid",
		"14948892946",
		false,
	},
	{
		"last second digit invalid",
		"14948892937",
		false,
	},
	{
		"valid number 1",
		"18071470110",
		true,
	},
	{
		"valid number 2",
		"74091671050",
		true,
	},
	{
		"valid number 3",
		"68461383032",
		true,
	},
	{
		"valid number 4",
		"56673392584",
		true,
	},
	{
		"valid number 5",
		"93212606504",
		true,
	},
	{
		"valid number 6",
		"64404737702",
		true,
	},
}
