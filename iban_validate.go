package main

import (
	"fmt"
	"iban_validate/yaml_reader"
	"log"
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

func validate_country(country_code string, iban_length int) bool {
	settings, err := yaml_reader.ReadYAML("settings.yaml")

	if err != nil {
		log.Fatalf("error: %v", err) // Logs the error and exits the program
	}

	if iban_length < settings.IBAN_rules.MinLength || iban_length > settings.IBAN_rules.MaxLength {
		fmt.Println("iban length is incorrect.")
		return false
	}

	for _, country := range settings.Countries {
		if strings.ToLower(country.Prefix) == country_code && country.Length != iban_length {
			fmt.Println("iban length is incorrect for:", country.Name)
			return false
		}
	}

	return true
}

func iban_isvalid(iban string) bool {
	iban = strings.ToLower(strings.TrimSpace(iban))

	rune_iban := []rune(iban)
	iban_length := len(rune_iban)

	country_code := iban[:2]

	if !validate_country(country_code, iban_length) {
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

	return n.Mod(n, big.NewInt(97)).Cmp(big.NewInt(1)) == 0

	// GB82WEST12345698765432
	// gb82WeSt12345698765432
}
