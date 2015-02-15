package console_scribeAmqp_test

func ExampleAmqpOptions() {
	AmqpOptions{
		"server":       "amqp://localhost",
		"rountingKey":  "config",
		"exchange":     "test-notif",
		"exchangeType": "fanout",
	}
}

func ExampleAmqpHook() {

	myConsole = console.NewConsole(console.ColorsOptions{})

	h, err := scribeHook.AmqpHook(scribeHook.AmqpOptions{
		"server":       "amqp://localhost",
		"exchange":     "test-notif",
		"exchangeType": "fanout",
		"rountingKey":  "console",
	})

	if err != nil {
		fmt.Println(err)
	} else {
		myConsole.AddHook(h)
	}
}
