package main

import (
	"fmt"
	"log"

	"github.com/gen2brain/beeep"
	"github.com/playwright-community/playwright-go"
)

func main() {
	fmt.Println("hello")

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// December 6th is the only confirmed date - there may be other showings throughout the week
	interstellarID := "gladiator-ii-72641"
	lincolnCenterShowtimes := "https://www.amctheatres.com/movies/%s/showtimes?date=2024-12-06&theatre=amc-lincoln-square-13"
	if _, err = page.Goto(fmt.Sprintf(lincolnCenterShowtimes, interstellarID)); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	noShowtimesMessage := "Sorry, no showtimes have been announced yet for this theatre. Showtimes for Friday and beyond are usually posted by Wednesday afternoon."
	content := page.GetByText(noShowtimesMessage)
	if content != nil {
		fmt.Println("No element could be found!")
		// continue wrap in for loop
	}

	// there are showings available, so send a message
	err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration*2)
	if err != nil {
		log.Fatalf("Could not issue audio")
	}
	// send discord notification here too

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
