package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var iban string

	fmt.Println("Enter IBAN to validate: ")

	fmt.Scanln(&iban)

	if iban_isvalid(iban) {
		fmt.Println("Valid")
	} else {
		fmt.Println("INVALID")
	}

}

func iban_isvalid(iban string) bool {
	iban = strings.ToLower(strings.TrimSpace(iban))

	rune_iban := []rune(iban)
	country_code := iban[:2]
	iban_length := len(rune_iban)

	if (country_code == "eg" && iban_length != 29) || 
	   (country_code == "it" && iban_length != 27) ||
	   (country_code == "no" && iban_length != 15) {
		return false
	}

	rune_iban = append(rune_iban[4:], rune_iban[:4]...)
	str_num := make([]string, len(rune_iban))

	for i := 0; i < len(rune_iban); i++ {
		if unicode.IsLetter(rune_iban[i]) {
			str_num[i] = strconv.Itoa(int(rune_iban[i] - 87))
		} else {
			str_num[i] = string(rune_iban[i])
		}
		//fmt.Println(strings.Join(str_num, ""))
	}

	n := new(big.Int)
	n.SetString(strings.Join(str_num, ""), 10)
	//fmt.Println(n.Mod(n, big.NewInt(97)))

	if n.Mod(n, big.NewInt(97)).Cmp(big.NewInt(1)) == 0 {
		return true
	}
	return false
	// GB82WEST12345698765432
	// gb82WeSt12345698765432
}
