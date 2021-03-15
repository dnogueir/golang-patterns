package main

import "fmt"

type baseNotifier interface {
	sendNotification(message string)
}

type notifier struct {
}

func (n *notifier) sendNotification(message string) {
	fmt.Println("Notification message:", message)
}

type smsDecorator struct {
	baseNotifier baseNotifier
}

func (sms *smsDecorator) sendNotification(message string) {
	sms.baseNotifier.sendNotification(message)
	fmt.Println("Sending sms notification ", message)
}

type facebookDecorator struct {
	baseNotifier baseNotifier
}

func (f *facebookDecorator) sendNotification(message string) {
	f.baseNotifier.sendNotification(message)
	fmt.Println("Sending facebook notification ", message)
}

type slackDecorator struct {
	baseNotifier baseNotifier
}

func (sl *slackDecorator) sendNotification(message string) {
	sl.baseNotifier.sendNotification(message)
	fmt.Println("Sending slack notification ", message)
}

func main() {

	message := "Your house is burning"

	notif := &notifier{}

	notif.sendNotification(message)

	facebookNotifier := &facebookDecorator{
		baseNotifier: notif,
	}
	fmt.Println("")
	fmt.Println("Facebook: ")

	facebookNotifier.sendNotification(message)

	smsNotifier := &smsDecorator{
		baseNotifier: notif,
	}

	fmt.Println("")
	fmt.Println("SMS: ")

	smsNotifier.sendNotification(message)

	slackNotifier := &slackDecorator{
		baseNotifier: notif,
	}

	fmt.Println("")
	fmt.Println("Slack: ")

	slackNotifier.sendNotification(message)

	//-------------------------------------------
	//------------combining them-----------------

	fmt.Println("")
	fmt.Println("Slack + facebook")

	slackAndFace := &slackDecorator{
		baseNotifier: facebookNotifier,
	}

	slackAndFace.sendNotification(message)

	fmt.Println("")
	fmt.Println("Slack + SMS")

	slackAndSms := &smsDecorator{
		baseNotifier: slackNotifier,
	}

	slackAndSms.sendNotification(message)

	fmt.Println("")
	fmt.Println("Facebook + SMS")

	facebookAndSms := &smsDecorator{
		baseNotifier: facebookNotifier,
	}

	facebookAndSms.sendNotification(message)

	fmt.Println("")
	fmt.Println("Facebook + SMS + slack")

	slackAndFacebookAndSms := slackDecorator{
		baseNotifier: facebookAndSms,
	}

	slackAndFacebookAndSms.sendNotification(message)
}
