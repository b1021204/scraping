package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"time"
)

func login(username string, password string) {
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
	if err := page.Navigate("https://sso.fun.ac.jp"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	// get ID and Password element. Set value
	elem_user := page.FindByName("username")
	elem_pass := page.FindByName("password")
	elem_user.Fill(username)
	elem_pass.Fill(password)
	// Submit
	if err := page.FindByClass("credentials_input_submit").Click(); err != nil {
		log.Fatalf("Failed to login:%v", err)
		return
	} else {
		fmt.Println("Succece login!!")
		elem_vm := page.FindByID("/Common/Manage_P")
		elem_vm.Click()
		time.Sleep(2 * time.Second)

		return
	}
}

func go_to_vmpage(username string, password string) {
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
	go_to_vmpage(username, password)

}
