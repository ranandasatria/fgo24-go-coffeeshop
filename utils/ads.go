package utils

import (
	"math/rand"
	"time"
)


func AdsChannel() chan string {
	ch := make(chan string)
	go func() {
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		ads := []string{
			"Coba Nasi Goreng Kampung, pedasnya mantap!\n",
			"Es Teh Kampul segar buat temen makan!\n",
			"Pisang Goreng enak buat ngemil!\n",
			"Pudding Coklat manisnya pas!\n",
		}
		for {
			ch <- ads[rng.Intn(len(ads))]
			time.Sleep(3 * time.Second)
		}
	}()
	return ch
}
