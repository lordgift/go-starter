package lcd

import (
	"strconv"
	"strings"
)

// ._.   ...   ._.   ._.   ...   ._.   ._.   ._.   ._.   ._.
// |.|   ..|   ._|   ._|   |_|   |_.   |_.   ..|   |_|   |_|
// |_|   ..|   |_.   ._|   ..|   ._|   |_|   ..|   |_|   ..|

var digits = [10][]string{
	0: {
		" _ ",
		"| |",
		"|_|",
	},
	1: {
		"   ",
		"  |",
		"  |",
	},
	2: {
		" _ ",
		" _|",
		"|_ ",
	},
	3: {
		" _ ",
		" _|",
		" _|",
	},
	4: {
		"   ",
		"|_|",
		"  |",
	},
	5: {
		" _ ",
		"|_ ",
		" _|",
	},
	6: {
		" _ ",
		"|_ ",
		"|_|",
	},
	7: {
		" _ ",
		"  |",
		"  |",
	},
	8: {
		" _ ",
		"|_|",
		"|_|",
	},
	9: {
		" _ ",
		"|_|",
		"  |",
	},
}

func LCD(n int) string {
	ns := strconv.Itoa(n)
	var ds [][]string
	for i := 0; i < len(ns); i++ {
		ds = append(ds, digits[ns[i]-'0'])
	}
	var num []string
	for i := 0; i < 3; i++ {
		var row []string
		for _, v := range ds {
			row = append(row, v[i])
		}
		num = append(num, strings.Join(row, " "))
	}

	return strings.Join(num, "\n")
}
