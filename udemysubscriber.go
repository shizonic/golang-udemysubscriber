package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/headzoo/surf.v1"
)

const (
	baseReferrerUrl string = "http://t3n.de/news/udemy-e-learning-kurse-kostenlos-833114/"
	baseUrl         string = "https://udemy.com"
)

func main() {
	// get udemy email and password
	fmt.Println("== udemy.com ==\n")
	fmt.Println("email: ")
	var email string
	fmt.Scanln(&email)
	fmt.Println("password (no output): ")
	password, err := terminal.ReadPassword(0)

	// login to udemy
	fmt.Println("\nstarting enrollment process...\n")
	bow := surf.NewBrowser()
	err = bow.Open("https://www.udemy.com/join/login-popup/?response_type=json")
	if err != nil {
		panic(err)
	}

	loginForm, _ := bow.Form("form#login-form")
	loginForm.Input("email", email)
	loginForm.Input("password", string(password))
	if loginForm.Submit() != nil {
		panic(err)
	}

	// parse checkout urls from t3n
	err = bow.Open(baseReferrerUrl)
	if err != nil {
		panic(err)
	}

	courseCounter := 0

	bow.Find("div#post-833114").Find("a").Each(func(_ int, link *goquery.Selection) {
		referrerUrl, _ := link.Attr("href")
		bow.Open(referrerUrl)
		checkoutUrl, _ := bow.Find("div.buy-box__element--buy-button").Find("a").Attr("href")

		// checkout and enroll course
		fmt.Println("Checking status of course:\n" + referrerUrl)
		if len(checkoutUrl) > 0 {
			fullCheckoutUrl := baseUrl + checkoutUrl
			bow.Open(fullCheckoutUrl)
			fmt.Println("=> Enrolling course:\n" + fullCheckoutUrl + "\n")
			courseCounter++
		} else {
			fmt.Println("=> Already enrolled\n")
		}
	})

	fmt.Printf("%v courses enrolled\n", courseCounter)
}
