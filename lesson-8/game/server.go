package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

var (
	answered        bool
	currentQuestion string
	currentAnswer   int
	operators       = []func(a, b int) (int, string){
		func(a, b int) (int, string) {
			return a + b, fmt.Sprintf("%d + %d = ?", a, b)
		},
		func(a, b int) (int, string) {
			return a - b, fmt.Sprintf("%d - %d = ?", a, b)
		},
		func(a, b int) (int, string) {
			return a * b, fmt.Sprintf("%d * %d = ?", a, b)
		},
	}
)

func main() {
	rand.Seed(time.Now().Unix())

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			if len(clients) == 1 {
				go nextQuestion()
			} else {
				cli <- fmt.Sprintf("The current question is: %s", currentQuestion)
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		answer, err := strconv.Atoi(input.Text())
		if err != nil {
			ch <- "incorrect input"
			continue
		}
		if answer != currentAnswer {
			ch <- "incorrect answer"
			continue
		}
		if answered {
			ch <- "question is already answered"
			continue
		}

		answered = true
		messages <- fmt.Sprintf("%s wins! The answer is: %d", who, answer)

		go nextQuestion()
	}
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func nextQuestion() {
	log.Println("generating next question")

	operator := operators[rand.Intn(len(operators))]

	currentAnswer, currentQuestion = operator(rand.Intn(10), rand.Intn(10))

	answered = false

	messages <- fmt.Sprintf("The next question is: %s", currentQuestion)
}
