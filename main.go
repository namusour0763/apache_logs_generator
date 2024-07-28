package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type WeightedItem struct {
	item   string
	weight int
}

type LogConfig struct {
	ipAddresses      []WeightedItem
	endpointStatuses []WeightedItem
	userAgents       []WeightedItem
}

var logConfig = LogConfig{
	ipAddresses: []WeightedItem{
		{item: "192.168.1.1", weight: 50},
		{item: "10.0.0.1", weight: 30},
		{item: "172.16.0.1", weight: 10},
		{item: "8.8.8.8", weight: 5},
		{item: "1.1.1.1", weight: 5},
		{item: "192.168.0.10", weight: 20},
		{item: "10.0.0.2", weight: 15},
		{item: "172.16.0.2", weight: 8},
		{item: "8.8.4.4", weight: 3},
		{item: "9.9.9.9", weight: 2},
		{item: "192.168.1.100", weight: 25},
		{item: "10.0.0.100", weight: 12},
		{item: "172.16.0.100", weight: 6},
		{item: "208.67.222.222", weight: 4},
		{item: "208.67.220.220", weight: 3},
	},
	endpointStatuses: []WeightedItem{
		{item: "/,200", weight: 40},
		{item: "/about,200", weight: 20},
		{item: "/contact,200", weight: 15},
		{item: "/products,200", weight: 10},
		{item: "/products,404", weight: 5},
		{item: "/services,200", weight: 8},
		{item: "/services,500", weight: 2},
		{item: "/blog,200", weight: 18},
		{item: "/blog/post1,200", weight: 12},
		{item: "/blog/post2,404", weight: 3},
		{item: "/faq,200", weight: 10},
		{item: "/support,200", weight: 7},
		{item: "/support,503", weight: 1},
		{item: "/login,200", weight: 25},
		{item: "/login,401", weight: 5},
		{item: "/register,200", weight: 15},
		{item: "/profile,200", weight: 10},
		{item: "/profile,403", weight: 2},
		{item: "/settings,200", weight: 8},
		{item: "/logout,302", weight: 6},
		{item: "/api/v1/users,200", weight: 15},
		{item: "/api/v1/products,200", weight: 12},
		{item: "/api/v1/orders,200", weight: 10},
	},
	userAgents: []WeightedItem{
		{item: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36", weight: 50},
		{item: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15", weight: 30},
		{item: "Mozilla/5.0 (X11; Linux x86_64; rv:89.0) Gecko/20100101 Firefox/89.0", weight: 20},
		{item: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36", weight: 40},
		{item: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36", weight: 25},
		{item: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0", weight: 15},
		{item: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Mobile/15E148 Safari/604.1", weight: 35},
		{item: "Mozilla/5.0 (iPad; CPU OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/91.0.4472.80 Mobile/15E148 Safari/604.1", weight: 20},
		{item: "Mozilla/5.0 (Android 11; Mobile; rv:68.0) Gecko/68.0 Firefox/88.0", weight: 15},
	},
}

func weightedRandomChoice(items []WeightedItem) string {
	totalWeight := 0
	for _, item := range items {
		totalWeight += item.weight
	}

	randomNumber := rand.Intn(totalWeight)
	for _, item := range items {
		randomNumber -= item.weight
		if randomNumber < 0 {
			return item.item
		}
	}

	return items[0].item // This should never happen, but it's here as a fallback
}

func generateLogEntry() string {
	ipAddress := weightedRandomChoice(logConfig.ipAddresses)
	endpointStatus := weightedRandomChoice(logConfig.endpointStatuses)
	userAgent := weightedRandomChoice(logConfig.userAgents)

	parts := strings.Split(endpointStatus, ",")
	endpoint := parts[0]
	statusCode := parts[1]

	timestamp := time.Now().Format("02/Jan/2006:15:04:05 -0700")
	method := "GET"
	protocol := "HTTP/1.1"
	bytesSent := rand.Intn(5000) + 500
	referer := "-"

	return fmt.Sprintf("%s - - [%s] \"%s %s %s\" %s %d \"%s\" \"%s\"",
		ipAddress, timestamp, method, endpoint, protocol, statusCode, bytesSent, referer, userAgent)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <lines_per_second>")
		os.Exit(1)
	}

	linesPerSecond, err := strconv.Atoi(os.Args[1])
	if err != nil || linesPerSecond <= 0 {
		fmt.Println("Please provide a valid positive integer for lines per second")
		os.Exit(1)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for i := 0; i < linesPerSecond; i++ {
			fmt.Println(generateLogEntry())
		}
	}
}
