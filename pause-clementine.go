package main

import (
	"code.google.com/p/gcfg"
	"flag"
	"github.com/brunoqc/go-clementine"
	"github.com/brunoqc/go-windows-session-notifications"
	"log"
)

var cfgFile = flag.String("config", "config.gcfg", "Config filename")

func HandleEvent(action string, clementine *clementine.Clementine) error {
	switch action {
	case "PLAY":
		return clementine.SimplePlay()
	case "PAUSE":
		return clementine.SimplePause()
	case "STOP":
		return clementine.SimpleStop()
	default:
		return nil
	}
}

func main() {
	flag.Parse()
	cfg := struct {
		Clementine struct {
			Host     string
			Port     int
			AuthCode int
			OnLock   string
			OnUnLock string
		}
	}{}

	err := gcfg.ReadFileInto(&cfg, *cfgFile)
	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}

	if cfg.Clementine.OnLock != "PLAY" &&
		cfg.Clementine.OnLock != "STOP" &&
		cfg.Clementine.OnLock != "NOTHING" &&
		cfg.Clementine.OnLock != "PAUSE" {
		log.Fatal(`OnLock must be "PLAY", "PAUSE", "STOP" or "NOTHING"`)
	}
	if cfg.Clementine.OnUnLock != "PLAY" &&
		cfg.Clementine.OnUnLock != "STOP" &&
		cfg.Clementine.OnUnLock != "NOTHING" &&
		cfg.Clementine.OnUnLock != "PAUSE" {
		log.Fatal(`OnUnLock must be "PLAY", "PAUSE" or "STOP" or "NOTHING"`)
	}

	clementine := clementine.Clementine{
		Host:     cfg.Clementine.Host,
		Port:     cfg.Clementine.Port,
		AuthCode: cfg.Clementine.AuthCode,
	}

	quit := make(chan int)

	changes := make(chan int, 100)
	closeChan := make(chan int)

	go func() {
		for {
			select {
			case c := <-changes:
				switch c {
				case session_notifications.WTS_SESSION_LOCK:
					log.Println("session locked")
					err := HandleEvent(cfg.Clementine.OnLock, &clementine)
					if err != nil {
						log.Println("Error: ", err.Error())
					}
				case session_notifications.WTS_SESSION_UNLOCK:
					log.Println("session unlocked")
					err := HandleEvent(cfg.Clementine.OnUnLock, &clementine)
					if err != nil {
						log.Println("Error: ", err.Error())
					}
				}
			}
		}
	}()

	session_notifications.Subscribe(changes, closeChan)

	// ctrl+c to quit
	<-quit
}
