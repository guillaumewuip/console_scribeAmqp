# console_scribeAmqp
An hook for console.go that send log to an server over AMQP

Created for the [Scribe.js log aggregator example](https://medium.com/@guillaumewuip/amqp-scribe-js-for-a-lightweight-logs-management-ed632f057a2a)

```go

package main

import (
	"fmt"
	K "github.com/guillaumewuip/console.go"
	scribeHook "github.com/guillaumewuip/console_scribeAmqp.go"
)

func main() {

	console := K.NewConsole(K.ColorsOptions{})

    //Create AMQP hook
	h, err := scribeHook.AmqpHook(scribeHook.AmqpOptions{
		"server":       "amqp://localhost",
		"exchange":     "test-notif",
		"exchangeType": "fanout",
		"rountingKey":  "console",
	})

	if err != nil {
		fmt.Println(err)
	} else {
		console.AddHook(h)
	}

	console.Tag("First", "Test").Time().File().Log("This log will be send over %s", "AMQP")
}

```

**Godoc :** http://godoc.org/github.com/guillaumewuip/console_scribeAmqp.go
