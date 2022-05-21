package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"math"
	"net/http"
	//"regexp"
	"strings"
)

func passCalc(password string) int {
	//txt := ""
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
	
	/*
	checkRegex := map[string]string {
		"strong": "(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])(?=.*\\W)",
		"moderate":"(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])",
		"medium":"",
		"weak":"",
	}

	matchS, _ := regexp.MatchString(checkRegex["strong"], password)
	matchM, _ := regexp.MatchString(checkRegex["strong"], password) 
	
	regex buat ngecek kalo isinya alfabet kecil/alfabet kapital/numerik doang,
	nanti nilai entropinya berkurang kalo gaada variasi 
	*/
	

	ent := l * math.Log2(pool) 
	fmt.Println("password entropy: ", math.Round(ent))
	
	res := int(ent) + int(l)
	
	if res >= 60 {
		
	}
	return int(math.Round(ent)) + int(l)
}

func renderHTML(w http.ResponseWriter, file string, data interface{}) {
	t, err := template.ParseFiles(file)
	if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
        return
	}
}

func calculate(w http.ResponseWriter, r *http.Request) {
	
	body, _ := ioutil.ReadAll(r.Body)

	userPass := string(body)[9:]

	fmt.Println(userPass)


	if passCalc(userPass) < 25 {
		res := map[string]interface{}{
			"PwdTxt": userPass, 
			"CalcRes":"Your password is weak",
		}
		renderHTML(w, "../client/index.html", res)
	} else {
		res := map[string]interface{}{
			"PwdTxt": userPass, 
			"CalcRes":"Your password is strong",
		}
		renderHTML(w, "../client/index.html", res)
	}
	
	//http.Redirect(w, r, "/results", http.StatusSeeOther)
}

func main() {
	fileServer := http.FileServer(http.Dir("../client"))
	http.Handle("/", http.StripPrefix("/", fileServer))
	http.HandleFunc("/calculate", calculate)
	http.ListenAndServe(":8080", nil)
}