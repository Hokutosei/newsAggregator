package realtimeServer

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Broker client/message handler
type Broker struct {
	clients map[chan string]bool

	newClients chan chan string

	defunctClients chan chan string

	messages chan string
}

// Start begin blocking sse events
func (b *Broker) Start() {
	fmt.Println("broker started..")

	go func() {

		for {
			select {
			case s := <-b.newClients:
				b.clients[s] = true
				log.Println("added new client!")

			case s := <-b.defunctClients:

				delete(b.clients, s)
				log.Println("removed client")

			case msg := <-b.messages:
				for s, _ := range b.clients {
					s <- msg
				}

				log.Printf("broadcast message to %d clients", len(b.clients))
			}
		}
	}()
}

func (b *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan string)

	// add client to list who will receive messages
	b.newClients <- messageChan

	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify

		b.defunctClients <- messageChan
		log.Println("HTTP connection just closed")
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		msg := <-messageChan

		fmt.Fprintf(w, "data: Message: %s\n\n", msg)

		f.Flush()
	}

	// log.Println("finished http request", r.URL.Path)
}

// Init will start and initialize all
func Init() {
	b := &Broker{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}

	b.Start()

	// test INIT
	http.Handle("/events/", b)

	// Generate a constant stream of events that get pushed
	// into the Broker's messages channel and are then broadcast
	// out to any clients that are attached.
	go func() {
		for i := 0; ; i++ {

			// Create a little message to send to clients,
			// including the current time.
			b.messages <- fmt.Sprintf("%d - the time is %v", i, time.Now())

			// Print a nice log message and sleep for 5s.
			log.Printf("Sent message %d ", i)
			time.Sleep(5 * 1e9)

		}
	}()

}
