package main

import (
	"github.com/sclevine/agouti"
	"log"
	"time"
)

func vmpage(username string, password string) {
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
	if err := page.Navigate("https://manage.p.fun.ac.jp/server_manage"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	time.Sleep(1 * time.Second)

	elem_user := page.FindByName("username")
	elem_pass := page.FindByName("password")
	elem_user.Fill(username)
	elem_pass.Fill(password)
	// Submit
	if err := page.FindByClass("credentials_input_submit").Click(); err != nil {
		log.Fatalf("Failed to login:%v", err)
		return
	}
	elem_choice := page.FindByXPath("/html/body/div/div/main/div/form/div[2]/div/span")
	elem_choice.Click()
	time.Sleep(2 * time.Second)
}

func main() {
	username := "b1021204"
	password := "EPa6ouQ2"
	//login(username, password)
	vmpage(username, password)

}
