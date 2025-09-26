package ping_pong_game

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func sayPong(c chan string) {
	c <- "pong"
}

func sayPing(c chan string) {
	c <- "ping"
}

func PingPongGame() {
	c := make(chan string)
	for i := 0; i < 5; i++ {
		go sayPong(c)
		go sayPing(c)
		x := <-c
		fmt.Println(x)

		y := <-c
		fmt.Println(y)
	}

}
