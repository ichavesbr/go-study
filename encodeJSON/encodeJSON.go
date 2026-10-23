package encodeJSON

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Message struct {
	Name string
	Body string
	Time time.Time
}

// Convert GO data into JSON
func EncodeJsonData() error {
	m := Message{"Alice", "Hello", time.Date(2011, 1, 25, 0, 0, 0, 0, time.UTC)}

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("JSON encoding error: %v", err)
	}

	fmt.Println(string(b))
	return nil
}
