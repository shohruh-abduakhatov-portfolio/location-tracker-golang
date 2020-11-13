// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os/exec"

	cmd "gitlab.com/logitab/back-end-team/location-tracker-go/cmd"
)

func main() {
	fmt.Println("Location Tracker Go - Starting...")
	cmd.Execute()

}

func runPythonTest() {
	cmd := exec.Command("pip3 install websocket", " && ", "python3 python.py")
	if err := cmd.Run(); err != nil {
		fmt.Println("Cannot run python script")
	}
}
