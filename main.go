// Wiggle
// This program will wiggle the mouse cursor every second
// Moving the mouse cursor to the top right corner of the screen will toggle the wiggler on and off
// Pressing the 'num lock' key will also toggle the wiggler on and off

// To enable alerts on macOS, open Script Editor and run the following command:
// display notification "Hello World" with title "My Title"
// This will ask for permission to send notifications
// After that, you can run the program and it will send notifications when the wiggler is toggled
package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var (
	step = 10

	sx       int
	sy       int
	displays = robotgo.DisplaysNum()
)

var wigglerFlag = true

func main() {
	robotgo.MouseSleep = 100
	computeScreenSize()

	go func() {
		nod := robotgo.DisplaysNum()
		if nod != displays {
			fmt.Println("Display no changed")
			computeScreenSize()
			displays = nod
		}
	}()

	go func() {
		var locked bool
		for {
			select {
			case <-time.After(time.Millisecond * 100):
				locX, locY := robotgo.Location()
				// fmt.Println("X: ", locX, "Y: ", locY)
				if locX >= sx-8 && locY < sy+8 {
					if !locked {
						locked = true
						toogleWiggler()
					}
				} else {
					locked = false
				}
			}
		}
	}()

	go func() {

		evChan := hook.Start()
		defer hook.End()

		for ev := range evChan {
			// Ingore mouse events and key up
			if ev.Kind != hook.KeyDown {
				continue
			}
			// fmt.Println("hook: ", ev)
			if ev.Rawcode == 71 && ev.Keychar == 27 {
				toogleWiggler()
			}
		}
	}()

	for {
		select {
		case <-time.After(time.Second * 1):
			if wigglerFlag {
				wiggle()
			}
		}
	}
}

func toogleWiggler() {
	if wigglerFlag {
		fmt.Println("Wiggler is off")
		beeep.Alert("Wiggler", "Wiggler is off", "./assets/icon.png")
	} else {
		fmt.Println("Wiggler is on")
		beeep.Alert("Wiggler", "Wiggler is on", "./assets/icon.png")
	}

	wigglerFlag = !wigglerFlag
}

func wiggle() {
	fmt.Println("Wiggling mouse")

	robotgo.MoveSmoothRelative(step, 0)
	robotgo.MoveSmoothRelative(-2*step, 0)
	robotgo.MoveSmoothRelative(step, 0)

	robotgo.MoveSmoothRelative(step/2, 0)
	robotgo.MoveSmoothRelative(-step, 0)
	robotgo.MoveSmoothRelative(step/2, 0)

}

func computeScreenSize() {
	// Get the max witdh of all displays
	// This is used to determine the top right corner of the screen
	for i := 0; i < displays; i++ {
		x, y, w, _ := robotgo.GetDisplayBounds(i)
		if x+w > sx {
			sx = x + w
		}
		if y < sy {
			sy = y
		}
	}

}
