package console_scribeAmqp

import (
	"encoding/json"
	"github.com/guillaumewuip/console.go"
	"github.com/streadway/amqp"
)

//Stores filename and file of a log
type Location struct {
	Filename string `json:"filename"`
	Line     int    `json:"line"`
}

//A Log
//will be converted to json
type Log struct {
	Type     string   `json:"type"`
	Tags     []string `json:"tags"`
	Location `json:"location"`
	Time     int64  `json:"time"`
	Message  string `json:"message"`
}

// AmqpOptions{
//	"server"          : "amqp://localhost",
//  "rountingKey"     : "",
//  "exchange"        : "",
//  "exchangeType"    : "",
// }
type AmqpOptions map[string]string

func AmqpHook(options AmqpOptions) (func(logger console.Logger) error, error) {

	var err error

	connection, err := amqp.Dial(options["server"])
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return func(logger console.Logger) error {

		log := Log{
			Type:    logger.Level,
			Tags:    logger.Tags,
			Time:    logger.Timestamp,
			Message: logger.Message,
			Location: Location{
				Filename: logger.Location.Filename,
				Line:     logger.Location.Line,
			},
		}

		//convert json
		m, _ := json.Marshal(log)

		channel.Publish(
			options["exchange"],
			options["routingKey"],
			false,
			false,
			amqp.Publishing{
				Headers:         amqp.Table{},
				ContentType:     "application/json",
				ContentEncoding: "utf8",
				Body:            []byte(string(m)),
				DeliveryMode:    amqp.Transient,
				Priority:        0,
			},
		)

		return nil
	}, nil

}
