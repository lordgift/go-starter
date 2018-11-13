package lcdexample

const (
	zero = ` _ 
| |
|_|`

	one = `   
  |
  |`

	two = ` _ 
 _|
|_ `

	three = ` _ 
 _|
 _|`

	four = `   
|_|
  |`

	five = ` _ 
|_ 
 _|`

	six = ` _ 
|_ 
|_|`

	seven = ` _ 
  |
  |`

	eight = ` _ 
|_|
|_|`

	nine = ` _ 
|_|
  |`
)

var digits = []string{
	zero, one, two, three, four, five, six, seven, eight, nine,
}

func Lcd(n int) string {
	return digits[n]
}
