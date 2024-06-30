package main

import (
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	// ブラウザはChromeを指定して起動
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	} // go to login page
	if err := page.Navigate("https://sso.fun.ac.jp/my.policy"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	// get ID and Password element. Set value
	elem_user := page.FindByName("usesrname")
	elem_pass := page.FindByName("password")
	elem_user.Fill("b1021204")
	elem_pass.Fill("EPa6ouQ2")
	// Submit
	if err := page.FindByClass("credentials_input_submit").Click(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}
}
