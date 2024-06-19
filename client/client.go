package client

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"net"
	"os"
	"strings"
	"sync"
)

var consoleMutex sync.Mutex

func ConnectToServer(address string) {
	count := 0

	conn, err := net.Dial("tcp", address)
	if err != nil {
		err = fmt.Errorf("Error when connecting the TChatP server: %w", err)
		panic(err)
	}
	defer conn.Close()

	go func() {
		reader := bufio.NewReader(conn)
		for {
			receivedMessage, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from server: ", err)
				return
			}
			writeToConsole(receivedMessage, &count)
		}
	}()

	for {

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		consoleMutex.Lock()
		fmt.Print("Enter your message: ")
		consoleMutex.Unlock()

		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing to server: ", err)
			return
		}
	}
}

func writeToConsole(text string, count *int) {
	consoleMutex.Lock()
	defer consoleMutex.Unlock()
	var col *color.Color
	text = strings.TrimSpace(text)
	if *count == 0 {
		col = color.New(color.FgMagenta)
		col.Println(text)
	} else {
		col = color.New(color.FgBlue)
		text = "\n" + text
		col.Println(text)

		fmt.Print("Enter your message: ")
	}

	*count += 1
}
