package main

import "fmt"

type Notifier interface {
	Send(message string)
}
type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Println("Sending message: #{message} (Sender: Email)")
}

type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) {
	fmt.Println("Sending message: #{message} (Sender: SMS)")
}

type NotificationService struct {
	notifier Notifier
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}

func CreateNotifier(t string) Notifier {
	if t == "sms" {
		return SMSNotifier{}
	}
	return EmailNotifier{}
}

func main() {
	s := NotificationService{
		notifier: CreateNotifier("email"),
	}
	s.SendNotification("Hello world!")
}

//Single responsibility
