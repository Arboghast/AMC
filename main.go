package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/playwright-community/playwright-go"
)

func main() {
	fmt.Println("Launching Playwright...")

	pw, _ := playwright.Run()

	for {
		browser, _ := pw.Chromium.Launch()
		page, _ := browser.NewPage()

		// December 6th is the only confirmed date - there may be other showings throughout the week
		interstellarID := "interstellar-76729"
		date := "2024-12-06"
		lincolnCenterShowtimes := "https://www.amctheatres.com/movies/%s/showtimes?date=%s&theatre=amc-lincoln-square-13"
		_, _ = page.Goto(fmt.Sprintf(lincolnCenterShowtimes, interstellarID, date))

		noShowtimesMessage := "Sorry, no showtimes have been announced yet for this theatre. Showtimes for Friday and beyond are usually posted by Wednesday afternoon."
		content := page.GetByText(noShowtimesMessage)
		if content == nil {
			// there are showings available, so make a system notification
			for {
				_ = beeep.Beep(beeep.DefaultFreq/2, beeep.DefaultDuration)
				time.Sleep(5 * time.Second)
			}
		}

		currTime := time.Now().Format(time.Kitchen)
		fmt.Println("No Showtimes are available!", currTime)
		_ = browser.Close()

		time.Sleep(30 * time.Second)
	}
}
