package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
)

var (
	typeProject string
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	color.NoColor = false
	writer := bufio.NewWriter(conn)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a type of project").
				Options(
					huh.NewOption("Projects Group", ".groupproj"),
					huh.NewOption("Projects", ".dproj"),
				).
				Value(&typeProject),
		),
	)
	err = form.Run()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(typeProject)
	// os.Stdout.Sync()
	_, _ = writer.WriteString(typeProject)
	writer.Flush()
}
