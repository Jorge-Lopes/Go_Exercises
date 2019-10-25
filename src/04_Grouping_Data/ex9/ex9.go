package main

import "fmt"

func main() {
	var favThings = map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	favThings[`Jorge`] = []string{`boxing`, `bike`, `420`}

	for i, v := range favThings {
		fmt.Println(i, "favorite things are")
		for j, f := range v {
			fmt.Println(j, f)
		}
	}
}
