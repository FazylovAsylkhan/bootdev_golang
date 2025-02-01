package main

import (
	"errors"
	"fmt"
	"strings"
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

	fmt.Println("lesson #32")
	send("Bob", 0)
	send("Alice", 1)
	send("Mangalam", 2)
	send("Ozgur", 3)

	fmt.Println("lesson #33")
	send2("Bob", 0, planFree)
	send2("Alice", 1, "some plan")
	send2("Mangalam", 2, planPro)
	send2("Ozgur", 3, planFree)

	fmt.Println("lesson #34")
	test16([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})

	fmt.Println("lesson #35")
	test17(0.2, 0.2, 0.6)

	fmt.Println("lesson #36")
	test18([]cost{
		{0, 0.21},
		{
			day: 2,
			value: 0.22,
		},
		{
			day: 2,
			value: 0.3,
		},
		{
			day: 5,
			value: 0.5,
		},
	})

	fmt.Println("lesson #37")
	test19(10, 10)

	fmt.Println("lesson #38")
	test20(
		[]string{
		"crap",
		"shoot",
		"dang",
		"frick",
		},
		[]string{
			"hey",
			"there",
			"john",
		},
	)
	test20(
		[]string{
		"crap",
		"shoot",
		"dang",
		"frick",
		},
		[]string{
			"hey",
			"shoot",
			"john",
		},
	)

	fmt.Println("lesson #38")
	test21(
		[]string {
			"John",
			"Bob",
			"Jill",
		},
		[]int {
			12345678,
			12345678,
			12345678,
		},
	)
	test21(
		[]string {
			"John",
			"Bob",
			"Jill",
		},
		[]int {
			12345678,
			12345678,
		},
	)

	fmt.Println("lesson #39")
	users3 := map[string]user3{
		"John": {
			name: "John",
			number: 221231,
			scheduledForDeletion: false,
		},
		"Bob": {
			name: "Bob",
			number: 221231,
			scheduledForDeletion: true,
		},
	}
	test22(users3, "John")
	test22(users3, "Bob")
	test22(users3, "somebody")

	fmt.Println("lesson #40")
	test23(
		[]string{"aa", "bb", "aa", "bb", "cc", "12", "12", "1"},
		[]string{"aa", "bb", "cc", "12", "1"}, 
	)

	fmt.Println("lesson #41")
	test24([]string{"Bob", "John"}, 'B', "Bob")

	fmt.Println("lesson #42")
	messages := []string{
		"Thanks for getting back to me.",
		"Great to see you again.",
	}
	test25(messages, addGreeting)
	test25(messages, addSignature)

	fmt.Println("lesson #43")
	dbErrors := []error {
		errors.New("Out of memory"),
		errors.New("cpu is pegged"),
	}
	test26("Errors on database server", dbErrors, colonDelimit)
	test26("Errors on database server", dbErrors, commaDelimit)

	fmt.Println("lesson #44")
	users := map[string]user4{
		"john": {
			name: "john",
			number: 34234234,
			admin: true,
		},
		"elon": {
			name: "elon",
			number: 32434265,
			admin: false,
		},
	}
	test27(users, "john")
	test27(users, "elon")
	test27(users, "somebody")

	fmt.Println("lesson #45")
	test28([]emailBill{
		{45},
		{32},
		{43},
		{12},
		{34},
		{54},
	})

	fmt.Println("lesson #46")
	test29([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no place like home",
	})

	fmt.Println("lesson #47")
	test30(Message{
		"Lane",
		"Textio is getting better everyday",
	})

	fmt.Println("lesson #48")
	test31([]string{
		"Allian is going to heck",
	})

	fmt.Println("lesson #49")
	email := email3{
		"This is my first draft",
		"sandra@mailio-test.com",
		"bullock@mailio-test.com",
	}
	test32(&email, "This is my second draft")
}

func test32(e *email3, newMessage string) {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage(newMessage)
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
}

func (e *email3) setMessage(newMessage string) {
	e.message = newMessage
}

func (e email3) print() {
	fmt.Printf("Message: %s\n", e.message)
	fmt.Printf("From address: %s\n", e.fromAddress)
	fmt.Printf("To address: %s\n", e.toAddress)
}

type email3 struct {
	message string
	fromAddress string
	toAddress string
}

func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
}

func test31(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

type Message struct {
	Recipient string
	Text string
}

func test30(msg Message) {
	defer fmt.Println("==================")
	sendMessage2(msg)
}

func sendMessage2(m Message) {
	fmt.Printf("To: %v\n", &m.Recipient)
	fmt.Printf("Message: %v\n", &m.Text)
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

func printReports(messages []string) {
	for _, message := range messages {
		printCostReport(func (msg string) int {
			return len(msg) * 2
		}, message)
	}
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

func test29(messages []string) {
	defer fmt.Println("===================================")
	printReports(messages)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int{
		sum += x

		return sum
	}
}

type emailBill struct {
	costInPennies int
}

func test28(bills []emailBill) {
	defer fmt.Println("============================")
	countAdder, costAdder := adder(), adder()

	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

const (
	logDeleted = "user deleted"
	logNotFound = "user not found"
	logAdmin = "adminDeleted"
)

func test27(users map[string]user4, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("===========================")
	log := logAndDelete(users, name)
	fmt.Println("Log:", log)
}

func logAndDelete(users map[string]user4, name string) (log string) {
	defer delete(users, name)
	user, ok := users[name]
	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	
	return logDeleted
}

type user4 struct {
	name string
	number int
	admin bool
}

func getLogger(formatter func(string, string) string) func(string, string) {
	return func (a, b string) {
		fmt.Println(formatter(a, b))
	}
}

func test26(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("===============================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}

func commaDelimit(first, second string) string {
	return first + ", " + second
}

func test25(messages []string, formatter func(string) string) {
	fmt.Println("===========================\n")
	for _, message := range messages {
		fmt.Printf("* %s -> %s\n", message, formatter(message))
	}
}

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}

	return formattedMessages
}

func addSignature(message string) string {
	return message + " King regards."
}

func addGreeting(message string) string {
	return "Hello! " + message
}

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstChar := rune(name[0])
		_, ok := nameCounts[firstChar]
		if !ok {
			nameCounts[firstChar] = make(map[string]int)
		}
		nameCounts[firstChar][name]++
	}

	return nameCounts
}

func test24(names []string, initial rune, name string) {
	fmt.Printf("Generating counts for %v names...\n", len(names))

	nameCounts := getNameCounts(names)
	count := nameCounts[initial][name]
	fmt.Printf("Count for [%c][%s]: %d\n", initial, name, count)
	fmt.Println("================================")
}

func getCounts(userIDs []string) map[string]int {
	counts := make(map[string]int)
	for _, userID := range userIDs {
		count := counts[userID]
		count++
		counts[userID] = count
	}

	return counts
}

func test23(userIDs, ids []string) {
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))

	counts := getCounts(userIDs)
	fmt.Println("Counts from select IDs:\n")
	for _, id := range ids {
		fmt.Printf("- %s: %d\n", id, counts[id])
	}
	fmt.Println("===================================================")
}

func test22(users map[string]user3, name string) {
	fmt.Printf("Attempting to delete %s\n", name)
	defer fmt.Println("=========================\n")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	if deleted {		
		fmt.Printf("Deleted: %v\n", name)
	} else {
		fmt.Printf("Did not delete: %s\n", name)
	}
}
func deleteIfNecessary(users map[string]user3, name string) (deleted bool, err error) {
	user, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if user.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}

	return false, nil
}

type user3 struct {
	name string
	number int
	scheduledForDeletion bool
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user2, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid sizes")
	}
	userMap := make(map[string]user2)
	for i, name := range names {
		userMap[name] = user2{
			name: name,
			phoneNumber: phoneNumbers[i],
		}
	}

	return userMap, nil
}

type user2 struct {
	name string
	phoneNumber int
}

func test21(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...\n")
	defer fmt.Println("========================================")
	users, err := getUserMap(names, phoneNumbers)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)

	}
}

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, badWord := range badWords {
		for _, message := range msg {
			if badWord == message {
				return i
			}
		}
	}

	return -1
}

func test20(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println("-", badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("====================================================")
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func test19(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

type cost struct {
	day int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}

	return costsByDay
}

func test18(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i :=0; i < len(costsByDay); i++ {
		fmt.Printf("- Day %v: %v\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")

}

func sum(nums ...float64) float64 {
	total := 0.0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	return total
}

func test17(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}

	return costs
}

func test16(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %v\n", costs[i])
	}
	fmt.Println("===== END REPORT =====")
}

const (
	planFree = "free"
	planPro = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if (plan == planPro) {
		return allMessages[:], nil
	}
	if (plan == planFree) {
		return allMessages[0:2], nil
	}

	return nil, errors.New("unsupported plan")
}


func send2(name string, doneAt int, plan string) {
	defer fmt.Println("====================================")
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages) - 1 {
			fmt.Println("complete failure")
		}
	}
}

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func send(name string, doneAt int) {
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages) - 1 {
			fmt.Println("complete failure")
		}
	}
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
