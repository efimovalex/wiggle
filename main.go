package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"log/slog"

	"github.com/gen2brain/beeep"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"golang.org/x/sync/errgroup"
)

var (
	verboseFlag     *bool
	degugFlag       *bool
	hotCornerEnable *bool
	idleDetection   *bool
	idleTime        *time.Duration
	wiggleInterval  *time.Duration

	step = 10

	sx       int
	sy       int
	displays = robotgo.DisplaysNum()

	lastTimeMoved = time.Now()

	wigglerFlag = true
	m           = sync.Mutex{}

	lastX *int16
	lastY *int16

	logLvl *slog.LevelVar = &slog.LevelVar{}
	logger *slog.Logger   = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLvl,
	}))
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	// Check for screen size changes
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(time.Second * 2):
				nod := robotgo.DisplaysNum()
				if nod != displays {
					logger.Info("Displays changed")
					computeScreenSize()
					displays = nod
				}
			}
		}
	})

	// Detect mouse in hot corner
	eg.Go(func() error {
		var locked bool
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(time.Millisecond * 100):
				locX, locY := robotgo.Location()
				if locX >= sx-20 && locY < sy+20 {
					if !locked {
						locked = true
						toogleWiggler()
					}
				} else {
					locked = false
				}
			}
		}
	})

	// Detect key presses and mouse actions
	// This will detect the num lock key and toggle the wiggler on and off
	// It will also detect mouse actions and turn off the wiggler if the mouse is moved
	eg.Go(func() error {
		evChan := hook.Start()
		defer hook.End()
		for {
			select {
			case <-ctx.Done():
				return nil
			case ev := <-evChan:
				// logger.Debug("Event", "event", ev)
				switch ev.Kind {
				case hook.KeyDown:
					if ev.Rawcode == 71 && ev.Keychar == 27 {
						toogleWiggler()
					}
					fallthrough
				case hook.MouseDown:
					fallthrough
				case hook.MouseHold:
					fallthrough
				case hook.MouseUp:
					fallthrough
				case hook.MouseWheel:
					fallthrough
				case hook.MouseDrag:
					m.Lock()
					lastTimeMoved = time.Now()
					m.Unlock()
					if wigglerFlag {
						logger.Info("Mouse action detected, turning off wiggler")
						toogleOff()
					}
				case hook.MouseMove:
					if lastX == nil || lastY == nil {
						lastX = &ev.X
						lastY = &ev.Y
						lastTimeMoved = time.Now()
						continue
					}
					if ev.Y != *lastY {
						m.Lock()
						lastY = &ev.Y
						lastX = &ev.X
						lastTimeMoved = time.Now()
						m.Unlock()
						if wigglerFlag {
							logger.Info("Mouse moved on Y axis, turning off wiggler")
							toogleOff()
						}
					}
				}

			}
		}
	})

	// Check if idle
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil

			case <-time.After(time.Second * 1):
				if !wigglerFlag && time.Since(lastTimeMoved) > *idleTime {
					logger.Info("Idle time exceeded, turning on wiggler")
					toogleOn()
				}
			}
		}
	})

	// Wiggler
	// This will wiggle the mouse every second if the wiggler is on
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(*wiggleInterval):
				if wigglerFlag {
					wiggle()
				}
			}
		}
	})

	err := eg.Wait()
	if err != nil {
		logger.Error("Error", "error", err)
	}
}

func toogleOff() {
	m.Lock()
	defer m.Unlock()
	wigglerFlag = false
	logger.Info("Wiggler is off")
	err := beeep.Notify("Wiggler", "Wiggler is off", "./assets/icon.png")
	if err != nil {
		logger.Error("Error sending notification", "error", err)
	}
}

func toogleOn() {
	m.Lock()
	defer m.Unlock()
	wigglerFlag = true
	logger.Info("Wiggler is on")
	err := beeep.Notify("Wiggler", "Wiggler is on", "./assets/icon.png")
	if err != nil {
		logger.Error("Error sending notification", "error", err)
	}
}

func toogleWiggler() {
	if wigglerFlag {
		toogleOff()
	} else {
		toogleOn()
	}
}

func wiggle() {
	logger.Debug("Wiggling mouse")

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

func init() {
	logLvl.Set(slog.LevelError)

	flag.Usage = func() {
		fmt.Print("\nUsage: wiggler [options] \n\nOptions:\n" +
			" -v                 Enable verbose logging\n" +
			" -vvv               Enable debug logging\n" +
			" -hotcorner         Enable hot corner detection\n" +
			" -h                 Show this help message \n" +
			" -idle              Disable idle detection. Enabled by default. Example: -idle=false\n" +
			" -idle-time         Set idle time in seconds. Default is 10 seconds. Example: -idle-time=5s\n" +
			" -wiggle-interval   Set wiggle interval in seconds. Default is 1 second. Example: -wiggle-interval=2s\n")
	}
	// read flags
	verboseFlag = flag.Bool("v", false, "verbose")
	degugFlag = flag.Bool("vvv", false, "debug")
	hotCornerEnable = flag.Bool("hotcorner", false, "enable hot corner")
	idleDetection = flag.Bool("idle", true, "enable idle detection")
	idleTime = flag.Duration("idle-time", 10*time.Second, "idle time in seconds")
	wiggleInterval = flag.Duration("wiggle-interval", 1*time.Second, "wiggle interval in seconds")

	flag.Parse()
	if *verboseFlag {
		logLvl.Set(slog.LevelInfo)
	} else if *degugFlag {
		logLvl.Set(slog.LevelDebug)
	}

	robotgo.MouseSleep = 100
	computeScreenSize()

}
