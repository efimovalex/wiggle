# 🖱️ Wiggle

Wiggle is a tiny utility that keeps your computer awake by gently wiggling the mouse cursor every second. Perfect for staying online during meetings, preventing screen lock, or just vibing with your idle time.

## 🚀 Features

- 🌀 Automatically wiggles your mouse cursor every second.
- 🎯 Move your cursor to the **top-right corner** of the screen to toggle the wiggle mode on/off.
- 🔢 Hit the **Num Lock** key to also toggle the wiggler.
- 💡 Lightweight and simple — set it and forget it.

## 📦 Installation

### Go 
```
$ CGO_ENABLE=1 go install github.com/efimovalex/wiggle@latest

$ wiggle& 
```

## 🧠 Usage
Start the script.
Let the cursor do the hustle.
Disable the wiggling by:
Moving your mouse to the top-right corner of the screen
Or pressing the Num Lock key.

## 😂 Bonus

![wiggle](assets/wiggle-shimmy.gif)

"Wiggle wiggle wiggle... yeah!"
— Your mouse, every second.
(Don't blame us if your mouse starts breakdancing.)


📄 License

MIT License. Do whatever you want, just don’t forget to wiggle responsibly.

Made with ❤️ and a whole lot of unnecessary movement.

## 🐛 Issues

To enable alerts on macOS (they are there but just not allowed), open Script Editor and run the following command:
```
display notification "Hello World" with title "My Title"
```

This will ask for permission to send notifications
After that, you can run the program and it will send notifications when the wiggler is toggled