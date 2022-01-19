package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)


func makeCall(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	elapsed := time.Since(start)
	log.Printf("status code: %d, start: %s end: %s\n", resp.StatusCode, start.Format("15:04:05"), elapsed.String())
}

func Trigger(cmd *cobra.Command, args []string) error{
	url, err := cmd.Flags().GetString("url")
	if err != nil {
		os.Exit(1)
		return err
	}
	for {
		trigger := time.Now().Second()
		if trigger == 0 || trigger == 20 || trigger == 40 {
			log.Println("----")
			var wg sync.WaitGroup
			for i := 0; i < 10; i++ {
				wg.Add(1)
				go makeCall(url, &wg)
			}
			wg.Wait()
			time.Sleep(time.Second)
		}
	}
}
