package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("lesson #1")
	fmt.Println("starting Textio server")



	fmt.Println("lesson #2")
	
	messagesFromDoris := []string{
		"You doing anything later??",
		"Did you get my last message?",
		"Don't leave me hanging...",
		"Please respond I'm lonely!",
	}
	numMessages := float64(len(messagesFromDoris))
	costPerMessage := .02

	// don't touch above this line

	totalCost := costPerMessage * numMessages
	
	// don't touch above this line

	fmt.Printf("Doris spent %.2f on text messages today\n")
	fmt.Println(totalCost)
	


	fmt.Println("lesson #3")

	var username string = "wagslane"
	var password string = "20947382822"
	
	// don't touch above this line
	fmt.Println("Authorization: Basic", username+":"+password)

	
	fmt.Println("lesson #4")

	var smsSendingLimit int = 22
	var costPerSMS float32 = 0.2
	var hasPermission bool = true
	var username2 string = "user"

	// initalize variables here

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username2,
	)

	
	fmt.Println("lesson #5")
	congrats := "Happy birthday"
	fmt.Println(congrats)
	
	fmt.Println("lesson #6")

	penniesPerText := 2
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
	
	fmt.Println("lesson #7")
	averageOpenRate, displayMessage := 0.23, "is the average open rate of your messages"

	fmt.Println(averageOpenRate, displayMessage)

	fmt.Println("lesson #8")

	accountAge := 2.6
	accountAgeInt := int(accountAge)

	fmt.Println("Your accout has existed for", accountAgeInt, "years")

	fmt.Println("lesson #9")

	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
	
	fmt.Println("lesson #10")
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("number of seconds in a hour:", secondsInHour)
	
	fmt.Println("lesson #11")
	const name = "Soul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Println(msg)
	
	fmt.Println("lesson #12")
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and max length of:", maxMessageLen)

	if (messageLen <= maxMessageLen) {
		fmt.Println("Message sent")
	} else {
		fmt.Printf("Message not sent")
	}

	fmt.Println("lesson #13")
	

	fmt.Println(concat("Lane", " happy birthday!"))
	fmt.Println(concat("Elon,", " hope that Tesla thing works out"))
	fmt.Println(concat("Go", " is fantastic"))

	fmt.Println("lesson #14")

	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)

	fmt.Println("you've sent", sendsSoFar, "messages")

	fmt.Println("lesson #15")
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)

	fmt.Println("lesson #16")
	test(messageToSend{
		phoneNumber: 12345678,
		message: "Thanks for signing up",
	})

	fmt.Println("lesson #17")
	test2(sender{
		rateLimit: 10000,
		user: user{
			name: "name",
			number: 12345678,
		},
	})

	fmt.Println("lesson #18")
	test3(authenticationInfo{
		username: "username",
		password: "password",
	})

	fmt.Println("lesson #19")
	test4(sendingReport{
		reportName: "FirstReport",
		numberOfSends: 100,
	})
	test4(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime: time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
}

func test4(m message) {
	sendMessage(m)
	fmt.Println("==========================================")
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage()) 
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
	reportName string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v`, sr.reportName, sr.numberOfSends)
}

type authenticationInfo struct {
	username string
	password string
}

func (user authenticationInfo) getBasicAuth() string{
	
	return fmt.Sprintf("Authorization: Basic %s:%s", user.username, user.password)
	// return "Authorization: Basic " + user.username + ":" + user.password
}

func test3(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("======================================")
}

func test2(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("================================")
}

type sender struct {
	rateLimit int
	user
}

type user struct {
	name string
	number int
}

type messageToSend2 struct {
	message string
	sender user
	recipient user
}
func canSendMessage(mToSend messageToSend2) bool {
	if mToSend.message == "" {
		return false
	}
	if mToSend.recipient.name == "" || mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.number == 0 || mToSend.sender.number == 0 {
		return false
	}

	return true
}

type messageToSend struct {
	message string
	phoneNumber int
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("==================================================")
}
func concat(s1 string, s2 string) string {
	return s1 + s2
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd

	return sendsSoFar
}

func getNames() (string, string) {
	return "John", "Doe"
}

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}