package main

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

func main() {
	var msn string
	msn = "08012345678"
	num, err := phonenumbers.Parse(msn, "JP")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s was parsed as [%v]\n", msn, num) // 08012345678 was parsed as [country_code:81  national_number:8012345678]

	msn = "818012345678"
	num, err = phonenumbers.Parse(msn, "JP")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s was parsed as [%v]\n", msn, num) // 818012345678 was parsed as [country_code:81  national_number:8012345678]

	msn = "+818012345678"
	num, err = phonenumbers.Parse(msn, "JP")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s was parsed as [%v]\n", msn, num) // +818012345678 was parsed as [country_code:81  national_number:8012345678]

	msn = "+818012345678"
	num, err = phonenumbers.Parse(msn, "US")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s was parsed as [%v]\n", msn, num) // +818012345678 was parsed as [country_code:81  national_number:8012345678]

	msn = "+18012345678"
	num, err = phonenumbers.Parse(msn, "US")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s was parsed as [%v]\n", msn, num) // +18012345678 was parsed as [country_code:1  national_number:8012345678]
}
