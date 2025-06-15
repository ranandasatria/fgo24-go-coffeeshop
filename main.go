package main

import "fgo24-go-weeklytask/menu"

func main() {
	menu.MainMenu()
}





// mutex example
// var (
// 	counter int
// 	mutex   sync.Mutex
// )

// func main() {
// 	for i := 0; i < 1000; i++ {
// 		go func() {
// 			mutex.Lock()
// 			fmt.Println("Counter:", counter)
// 			mutex.Unlock()
// 			counter++
// 		}()
// 	}

// 	time.Sleep(time.Second)
// }
