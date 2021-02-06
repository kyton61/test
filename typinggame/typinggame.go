package typinggame

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const timelimit = 30

var correctAnsNum = 0
var animals = [10]string{"cat", "dog", "pigeion", "giraffe", "rat", "snake", "lion", "rhinoceros", "peacock", "flamingo"}
var countires = [10]string{"japan", "america", "china", "russia", "korea", "australia", "hungary", "brazil", "sweden", "turkey"}

// assert recieve error type and if reciever have error, print error message and exit function
func assert(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sendQuestion(chqes chan<- string, chans <-chan int) {
	for {
		<-chans
		i := rand.Int() % 10
		question := countries[i]
		chqes <- question
	}
}

func recieveQuestion(chqes <-chan string, chans chan<- int) {
	for {
		question := <-chqes
		var answer string
		fmt.Println("question:", question)
		fmt.Print("type:")
		fmt.Scanln(&answer)
		if question == answer {
			correctAnsNum++
			fmt.Println("good!!")
		} else {
			fmt.Println("bad!")
		}
		chans <- 0
	}
}

// Play typinggame
func Play() {
	chqes := make(chan string)
	chans := make(chan int)
	correctAnsNum := 0
	ansNum := 0
	rand.Seed(time.Now().UnixNano())

	// debug
	/*
		f, err := os.Create("trace.out")
		assert(err)
		defer f.Close()
		trace.Start(f)
		defer trace.Stop()
	*/

	// 問題を送り続ける
	go sendQuestion(chqes, chans)

	// 問題を受信して標準出力に出題
	go func(chqes <-chan string, chans chan<- int) {
		for {
			// 1ずつ取り出し
			question := <-chqes
			ansNum++
			var answer string
			fmt.Println("question:", question)
			fmt.Print("type:")
			fmt.Scanln(&answer)
			if question == answer {
				correctAnsNum++
				fmt.Println("good!!")
			} else {
				fmt.Println("bad!")
			}
			chans <- 0
		}
	}(chqes, chans)

	// 問題出力スタート
	fmt.Println("Start!")
	chans <- 0

	time.Sleep(timelimit * time.Second)
	fmt.Println("finish!!!!")
	fmt.Printf("Result:%d points", correctAnsNum)
}

// Play2 is refuctered from Play function
func Play2() {
	chqes := make(chan string)
	chans := make(chan int)
	chtime := make(chan int)
	rand.Seed(time.Now().UnixNano())
	// 問題を送り続ける
	go sendQuestion(chqes, chans)
	// 問題を出題し続ける
	go recieveQuestion(chqes, chans)
	// 問題出力スタート
	fmt.Println("Start!")
	chans <- 0

	select {
	case <-chtime:
	case <-time.After(timelimit * time.Second):
		close(chans)
		close(chqes)
		close(chtime)
		fmt.Println("finish!!!")
		fmt.Printf("Result:%d points", correctAnsNum)
		os.Exit(0)
	}

}
