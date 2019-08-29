package work

import (
	"log"

	"github.com/dnilosek/fib-overkill/lib/database"
	"github.com/go-redis/redis"
)

type Listener struct {
	database *database.RedisDB
}

func NewListener(db *database.RedisDB) *Listener {
	// Create the Listener
	return &Listener{
		database: db,
	}
}

// Listen to a category in redis and return messages to a channel
func (listener *Listener) Listen(channel string) chan string {

	// Get channel subscription
	pubSub := listener.database.Client.Subscribe(channel)
	msgChan := make(chan string)
	// Listen
	go func() {
		for {
			iface, err := pubSub.Receive()
			if err != nil {
				log.Println("Error recieved:", err)
				close(msgChan)
				return
			}
			switch iface.(type) {
			case *redis.Subscription:
				log.Printf("Subscription to channel [%v] Succeeded\n", channel)
			case *redis.Message:
				msgChan <- iface.(*redis.Message).Payload
			}
		}
	}()
	return msgChan
}
