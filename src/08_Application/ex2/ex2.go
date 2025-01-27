package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)
	var p1 []person

	fmt.Println(p1)

	err := json.Unmarshal([]byte(s), &p1)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range p1 {
		fmt.Println("\nPerson: ", i)
		fmt.Println(v.First, v.Last, v.Age)
		for _, saying := range v.Sayings {
			fmt.Println("\t", saying)
		}
	}

}
