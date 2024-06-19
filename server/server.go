package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/fatih/color"
)


type User struct {
	username string // Username
	isAdmin  bool   // Defines if the user is administrador
	conn     net.Conn
}

var users []User

var serverMutex sync.Mutex

func ListenConn(address string) {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		err = fmt.Errorf("Error when starting the TChatP Listener: %w", err)
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			err = fmt.Errorf("Error when accepting connections for TChatP: %w", err)
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	green := color.New(color.FgGreen)
	green.Println("New Connection from ", conn.RemoteAddr().String())

	conn.Write([]byte("Tell us your nickname, so we can register you: \n"))

	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading new user username: ", err)
		return
	}

	nameTrimmed := strings.TrimSpace(name)

	user := User{
		username: nameTrimmed,
		isAdmin:  false,
		conn:     conn,
	}

	serverMutex.Lock()
	users = append(users, user)
	serverMutex.Unlock()

	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Error while reading User %s message: %e", user.username, err)
			return
		}

		message := strings.TrimSpace(msg)

		fullMessage := fmt.Sprintf("%s says: %s \n", user.username, message)
        yellow := color.New(color.FgYellow)
		yellow.Printf("%s says: %s \n", user.username, message)

		go broadcastMessage(user, []byte(fullMessage))
	}

}

func broadcastMessage(sender User, msg []byte) {
	serverMutex.Lock()
	defer serverMutex.Unlock()

	for _, user := range users {
		if user.conn == sender.conn {
			continue
		}

		conn := user.conn
		_, err := conn.Write(msg)
		if err != nil {
			red := color.New(color.FgRed)
			red.Println("Error while writing to user: ", err)
			return
		}
	}
}
