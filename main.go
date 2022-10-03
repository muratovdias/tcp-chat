package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	port := "8989"
	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) != 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if err := NewServer().Run(port); err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	mut      sync.Mutex
	users    map[net.Conn]string
	messages chan message
}

type message struct {
	text string
	name string
	conn net.Conn
}

func NewServer() *Server {
	return &Server{
		users:    make(map[net.Conn]string),
		messages: make(chan message),
	}
}

func newMessage(msg string, userName string, conn net.Conn) message {
	return message{
		text: userText(userName) + msg,
		name: userName,
		conn: conn,
	}
}

func (s *Server) Run(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	if port == "8989" {
		fmt.Println("Listening on the port :8989")
	} else {
		fmt.Printf("Listening on the port :%s\n", port)
	}
	file, err := os.Create("history.txt")
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		conn.Write([]byte("Welcome to TCP-Chat!\n"))
		content, err := ioutil.ReadFile("logo.txt")
		if err != nil {
			fmt.Println(err)
		}
		// create .txt file to save the history of chat
		conn.Write(content)
		go s.broadcaster(conn)
		go s.handleConn(conn, file)
	}
}

func (s *Server) handleConn(myConn net.Conn, file *os.File) error {
	defer myConn.Close()
	reader := bufio.NewReader(myConn)
	var userName, text string
	// take a user name and check it
	for {
		myConn.Write([]byte("[ENTER YOUR NAME]:"))
		userName, _ = reader.ReadString('\n')
		userName = userName[:len(userName)-1]
		flag := checkUserName(userName, s.users, myConn)
		if flag {
			continue
		}
		s.mut.Lock()
		if len(s.users) > 9 {
			myConn.Write([]byte("Server is full, try to join later\n"))
			myConn.Close()
		}
		s.users[myConn] = userName
		s.mut.Unlock()
		break
	}
	history, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	myConn.Write(history)
	text = fmt.Sprintf("\r %s has joined our chat...  \n", userName)
	saveHistory(text, file) // store the history of chat
	s.messages <- newMessage(text, userName, myConn)
	// take a massage and check it
	for {
		msg, err := reader.ReadString('\n')
		if len(msg) == 1 {
			myConn.Write([]byte(userText(userName)))
			continue
		}
		if err != nil {
			s.mut.Lock()
			delete(s.users, myConn)
			s.mut.Unlock()
			text = fmt.Sprintf("\r %s has left our chat...  \n", userName)
			saveHistory(text, file) // store the history of chat
			s.messages <- newMessage(text, userName, myConn)
			break
		} else if strings.TrimSpace(msg) == "" {
			continue
		} else {
			historyText := userText(userName) + msg
			saveHistory(historyText, file)
			s.messages <- newMessage(msg, userName, myConn)
		}
	}
	return nil
}

func (s *Server) broadcaster(myConn net.Conn) {
	for {
		msg := <-s.messages
		s.mut.Lock()
		for conn, name := range s.users {
			if msg.conn == conn {
				msg.conn.Write([]byte(userText(msg.name)))
				continue
			}
			conn.Write([]byte(clear(msg.text) + msg.text + userText(name)))
		}
		s.mut.Unlock()
	}
}

func checkUserName(userName string, users map[net.Conn]string, myConn net.Conn) bool {
	if strings.TrimSpace(userName) == "" {
		myConn.Write([]byte("The username cannot be empty...\n"))
		return true
	}
	for _, name := range users {
		if name == userName {
			myConn.Write([]byte("Username already taken...\n"))
			return true
		}
	}
	return false
}

func userText(userName string) string {
	return fmt.Sprintf("\r[%s][%s]:", time.Now().Format("01-02-2006 15:04:05"), userName)
}

func clear(a string) string {
	return "\r" + strings.Repeat(" ", 27+len(a)) + "\r"
}

func saveHistory(h string, f *os.File) {
	if _, err := f.WriteString(h); err != nil {
		log.Fatal()
	}
}
