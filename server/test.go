package main

import (
	"fmt"
	"math"
	"strings"
	"regexp"
)
func main() {
	res := math.Log2(8)
	fmt.Println("binary log:", res)

	s := "test"
	passCalc("abc")
	passCalc("abcDeF")
	passCalc("abc1")
	passCalc("abcDeF12!")
}

func passCalc(password string) {
	pool := float64(0)
	l := float64(len(password))

	check := map[string]string {
		"lower": "abcdefghijklmnopqrstuvwxyz",
		"upper": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"number": "0123456789",
		"other": "`~!@#$%^&*()-=_+[{]}\\|;':\",.<>/?",
	}

	if strings.ContainsAny(password, check["other"]) {
		pool += 32
	} 
	if strings.ContainsAny(password, check["number"]) {
		pool += 10
	} 
	if strings.ContainsAny(password, check["upper"]) {
		pool += 26
	} 
	if strings.ContainsAny(password, check["lower"]) {
		pool += 26
	} 

	checkRegex := map[string]string {
		"moderate":"(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])",
		"medium":"(?=.*[a-z])(?=.*[A-Z])",
		"weak1":"(?=.*[a-z])",
		"weak2":"(?=.*[A-Z])",
		"weak3":"(?=.*\\d)",		
	}

	matchMed, _ := regexp.MatchString(checkRegex["medium"], password)

	matchMod, _ := regexp.MatchString(checkRegex["moderate"], password) 

	matchW1, _ := regexp.MatchString(checkRegex["weak1"], password) 

	matchW2, _ := regexp.MatchString(checkRegex["weak2"], password) 

	matchW3, _ := regexp.MatchString(checkRegex["weak3"], password) 

	if matchMed {
		pool -= 4
	} else if matchMod {
		pool -= 2
	}

	ent := l * math.Log2(pool) 
	fmt.Println("password entropy: ", math.Round(ent))
	
}
