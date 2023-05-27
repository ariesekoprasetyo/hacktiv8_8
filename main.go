package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go reloadData()

	fmt.Println("Press Ctrl+C to stop...")

	<-stop

	fmt.Println("Stopping the application...")
}

func reloadData() {
	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		var waterStatus string
		if water < 5 {
			waterStatus = "Aman"
		} else if water >= 6 && water <= 8 {
			waterStatus = "Siaga"
		} else {
			waterStatus = "Bahaya"
		}

		var windStatus string
		if wind < 6 {
			windStatus = "Aman"
		} else if wind >= 7 && wind <= 15 {
			windStatus = "Siaga"
		} else {
			windStatus = "Bahaya"
		}

		err := updateData(water, wind)
		if err != nil {
			log.Println("Failed to update data:", err)
		} else {
			log.Println("Data updated: water =", water, "wind =", wind, "water status =", waterStatus, "wind status =", windStatus)
		}

		time.Sleep(15 * time.Second)
	}
}

func updateData(water, wind int) error {
	url := "https://example.com/api/update"
	payload := []byte(fmt.Sprintf(`{"water": %d, "wind": %d}`, water, wind))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response: %s", resp.Status)
	}

	return nil
}
