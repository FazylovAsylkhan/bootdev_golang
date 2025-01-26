package main

import (
	"errors"
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

	if messageLen <= maxMessageLen {
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
		message:     "Thanks for signing up",
	})

	fmt.Println("lesson #17")
	test2(sender{
		rateLimit: 10000,
		user: user{
			name:   "name",
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
		reportName:    "FirstReport",
		numberOfSends: 100,
	})
	test4(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})

	fmt.Println("lesson #20")
	test5(contractor{
		name:         "John",
		hourlyPay:    150,
		hoursPerYear: 240,
	})
	test5(fullTime{
		name:   "Arman",
		salary: 40000,
	})

	fmt.Println("lesson #21")
	e := email{
		subscribed: true,
		body:       "Hello world\n",
	}
	test6(e, e)

	fmt.Println("lesson #22")
	test7(email2{
		toAddress:    "adsad@sd.s",
		isSubscribed: true,
		body:         "Hello",
	})
	test7(sms{
		toPhoneNumber: "231232323",
		isSubscribed:  false,
		body:          "Hi",
	})
	test7(invalid{})
	
	fmt.Println("lesson #23")
	test8("Thanks for coming in to our flower shop.", "We hope you enjoyed your gift.")
	test8("Thanks for joinging us.", "Have a good day.")
	
	fmt.Println("lesson #24")
	test9(1.4, "+1 (435) 555 0923")
	test9(2.1, "+2 (702) 555 3452")
	test9(32.10, "+1 (801) 555 7456")
	test9(14.40, "+1 (435) 555 6445")

	fmt.Println("lesson #25")
	test10(10.00, 0.00)
	test10(10.00, 2.00)

	fmt.Println("lesson #26")
	test11(10.00, 0.00)
	test11(10.00, 2.00)

	fmt.Println("lesson #27")
	test12(12);
	test12(5)
	test12(1)
	
	fmt.Println("lesson #28")
	test13(9)
	test13(50)

	fmt.Println("lesson #29")
	test14(1.1, 5)
	test14(1.3, 10)

	fmt.Println("lesson #30")
	fizzbuzz()

	fmt.Println("lesson #31")
	test15(15)
}

func test15(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("========================================")
}

func printPrimes(max int) {
	for n := 2; n < max + 1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		} 
		if n % 2 == 0 {
			continue
		}
		isPrime := true 
		for i := 3; i * i < n + 1; i++ {
			if n % i == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("=========================")
}

func test14(costMultiplier float64, maxCostInPennies int) {
	fmt.Printf("Multiplier: %.1f\n", costMultiplier)
	fmt.Printf("Max cost: %v\n", maxCostInPennies)
	maxMessages := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
	fmt.Printf("Max messages you can send: %v\n", maxMessages)
	fmt.Println("================================================")
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	for actualCostInPennies <= float64(maxCostInPennies) {
		maxMessagesToSend++
		actualCostInPennies *=costMultiplier
	}
	return maxCostInPennies
}

func test13(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
}

func maxMessages(thresh float64) int {
	totalCost := 0.00
	for i:= 0; ; i++ {
		totalCost += 1.00 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}

}

func bulkSend(numMessage int) float64 {
	cost := 0.00
	for i := 0; i < numMessage; i++ {
		cost += 1.00 + (0.01 * float64(i))
	}
	return cost
}

func test12(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost =%.2f\n", cost)
	fmt.Println("===========================================")
}

func test11(x, y float64) {
	defer fmt.Println("==========================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", x,y)
	quotient, err := divide2(x,y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func divide2(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("No dividing by 0")
	}
	return x / y, nil
}

func test10(dividend, divisor float64) {
	fmt.Printf("Dividing %v by %v ...\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		fmt.Println("=================================")
		return
	}
	fmt.Printf("Quotient: %v\n", quotient)
	fmt.Println("=================================")
}

type divideError struct {
	dividend float64
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0, divideError {
			dividend: dividend,
		}
	}
	 
	return dividend/divisor, nil
}

func (de divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", de.dividend)
}

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%s' can not be sent", cost, recipient) 
}

func test9(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("===================================")
}

func test8(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("========")
	fmt.Println("Message for Customer:", msgToCustomer)
	fmt.Println("Message for Spouse:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if (err != nil) {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costForCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costForSpose, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return costForCustomer + costForSpose, nil
}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

type invalid struct {
}

// cost implements expense.
func (i invalid) cost() float64 {
	panic("unimplemented")
}

func test7(e expense) {
	address, cost := getExpenseReport(e)
	switch v := e.(type) {
	case email2:
		fmt.Printf("Cost: %v\n", cost)
		fmt.Printf("Email Address(%T): %s\n", e, address)
		fmt.Printf("Message: %s\n", v.body)
		fmt.Println("=============================")
	case sms:
		fmt.Printf("Cost: %v\n", cost)
		fmt.Printf("Phone number(%T): %s\n", e, address)
		fmt.Printf("Message: %s\n", v.body)
		fmt.Println("=============================")
	default:
		fmt.Printf("Invalid Report (%T)\n", e)
		fmt.Println("=============================")
	}
}

func getExpenseReport(e expense) (string, float64) {
	email, ok := e.(email2)
	if ok {
		return email.toAddress, email.cost()
	}
	sms, ok := e.(sms)
	if ok {
		return sms.toPhoneNumber, sms.cost()
	}

	return "", 0.0
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.01
	}

	return float64(len(s.body)) * 0.03
}

func (e email2) cost() float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * 0.01
	}

	return float64(len(e.body)) * 0.05
}

type sms struct {
	toPhoneNumber string
	isSubscribed  bool
	body          string
}
type email2 struct {
	toAddress    string
	isSubscribed bool
	body         string
}

func test6(e expense, p print) {
	fmt.Printf("Cost's email: %v\n", e.cost())
	p.print()
	fmt.Println("=================================================")
}

func (e email) cost() float64 {
	if !e.subscribed {
		return float64(len(e.body)) * 0.05
	}
	return float64(len(e.body)) * 0.01
}

func (e email) print() {
	fmt.Printf(e.body)
}

type expense interface {
	cost() float64
}

type print interface {
	print()
}

type email struct {
	subscribed bool
	body       string
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

type fullTime struct {
	name   string
	salary int
}

func test5(c employee) {
	fmt.Printf("%s's salary equals %d\n", c.getName(), c.getSalary())
	fmt.Println("================================================")
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return (c.hourlyPay * c.hoursPerYear) / 12
}

func (ft fullTime) getName() string {
	return ft.name
}

func (ft fullTime) getSalary() int {
	return ft.salary
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
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v`, sr.reportName, sr.numberOfSends)
}

type authenticationInfo struct {
	username string
	password string
}

func (user authenticationInfo) getBasicAuth() string {

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
	name   string
	number int
}

type messageToSend2 struct {
	message   string
	sender    user
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
	message     string
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
