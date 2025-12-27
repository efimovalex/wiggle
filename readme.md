# 🖱️ Wiggle

**Wiggle** is a tiny utility that keeps your computer awake by gently wiggling the mouse cursor every few seconds.  
Perfect for staying online during meetings, avoiding idle status in work chats, preventing screen lock, dodging AFK kicks in GeForce Now, or just vibing during idle time.

---

## 📱 Platform Support

- **Tested** on macOS  
- **Untested** on Linux (should work, but not guaranteed)  
- **Planned** for Windows — currently not working due to [compilation issues](https://github.com/robotn/gohook/blob/master/hook.go#L20C34-L20C41)

Have a problem? [Open an issue](https://github.com/efimovalex/wiggle/issues)!

---

## Features

- Automatically wiggles the mouse cursor every few seconds
- Starts automatically after 30 seconds of user inactivity
- Pauses when you move the mouse, scroll, or press a key
- Lightweight and simple — set it and forget it
- Sends macOS notifications when the wiggler is toggled on/off

---

## Installation

### Via Go

```bash
$ CGO_ENABLE=1 go install github.com/efimovalex/wiggle@latest

$ wiggle &

```

## Usage
1. Start the script.
2. Let the cursor do the hustle.
3. Wiggling pauses on:
    - Any mouse movement or action
    - Any keyboard key press
4. If idle for 30 seconds, it resumes automatically.
5. To stop wiggling, simply close the script or use `Ctrl+C` in the terminal.
6. To run it in the background, use `wiggle &`. Use `fg` to bring it back to the foreground or `kill` to stop it.


## Options
```
$ wiggle -h
Usage: wiggle [options]

Options:
  -v                 Enable verbose logging
  -vvv               Enable debug logging
  -h                 Show this help message
  -idle-time         Set idle time before starting (e.g. -idle-time=5s, default: 30s)
  -wiggle-interval   Set wiggle interval (e.g. -wiggle-interval=2s, default: 5s)
```

## Bonus

![wiggle](assets/wiggle-shimmy.gif)

"Wiggle wiggle wiggle... yeah!"
— Your mouse, every second.
(Don't blame us if your mouse starts breakdancing.)


📄 License

MIT License. Do whatever you want, just don’t forget to wiggle responsibly.

Made with love and a whole lot of unnecessary movement.

## Issues

To enable alerts on macOS (they are there but just not allowed), open Script Editor and run the following command:
```
display notification "Hello World" with title "My Title"
```

This will ask for permission to send notifications
After that, you can run the program and it will send notifications when the wiggler is toggled
