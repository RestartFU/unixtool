package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var out string
	switch os.Args[1] {
	case "brightness_percent":
		out = brightness_percent()
	}
	fmt.Println(out)
}

func brightness_percent() string {
	cmd := exec.Command("brightnessctl", "get")

	out := bytes.NewBuffer([]byte{})
	cmd.Stdout = out

	cmd.Run()

	n, err := strconv.Atoi(strings.TrimSpace(out.String()))
	if err != nil {
		return "could not fetch brightness level"
	}

	percent := float64(n) / 2.55
	return fmt.Sprintf("%.0f", percent)
}
