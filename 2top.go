package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:\n\ttotop <window id>")
		os.Exit(1)
	}

	display, err := xgb.NewConn()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open display: %v\n", err)
		os.Exit(1)
	}
	defer display.Close()

	id, err := strconv.ParseUint(os.Args[1], 0, 32) // Use 0 for automatic base detection
	if err != nil {
	    fmt.Fprintf(os.Stderr, "Invalid window ID: %v\n", err)
	    os.Exit(1)
	}

	window := xproto.Window(id)

	xproto.ConfigureWindow(display, window, xproto.ConfigWindowStackMode, []uint32{xproto.StackModeAbove})
	xproto.SetInputFocus(display, xproto.InputFocusNone, window, xproto.TimeCurrentTime)
}
