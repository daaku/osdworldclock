// Osdworldclock briefly shows the time in a few cities as an on-screen
// overlay, then quits.
package main

import (
	"log"
	"time"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

var cities = []struct{ name, zone string }{
	{"Mumbai", "Asia/Kolkata"},
	{"Antwerp", "Europe/Brussels"},
	{"Dubai", "Asia/Dubai"},
	{"London", "Europe/London"},
	{"NYC", "America/New_York"},
	{"SF", "America/Los_Angeles"},
}

func timeIn(t time.Time, zone string) string {
	l, err := time.LoadLocation(zone)
	if err != nil {
		log.Fatal(err)
	}
	return t.In(l).Format("3:04pm")
}

func view(c *ui.Context) {
	now := time.Now()
	white := ui.RGBA(255, 255, 255, 1)
	c.Root().Background(ui.RGBA(13, 13, 13, 0.75))
	ui.Column(c).Fill().Center().Children(func() {
		ui.Text(c, "\U0001f557").FontSize(128)
		ui.Row(c).Children(func() {
			for _, city := range cities {
				ui.Column(c).AlignItems(ui.Center).Padding(0, 28, 16, 16).Children(func() {
					ui.Text(c, city.name).FontSize(32).Bold().TextColor(white)
					ui.Text(c, timeIn(now, city.zone)).FontSize(24).Bold().TextColor(white)
				})
			}
		})
	})
}

func main() {
	mygo.App.WhenReady(func() {
		w := mygo.NewWindow(mygo.WindowOptions{
			Title:         "World Clock",
			Panel:         true, // an overlay layer surface on Wayland
			Frameless:     true,
			Transparent:   true,
			DisableResize: true,
			Hidden:        true,
			Width:         1000,
			Height:        320,
			Content:       ui.View(view),
		})
		w.ShowInactive() // don't take the keyboard
		time.AfterFunc(5*time.Second, mygo.App.Quit)
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
