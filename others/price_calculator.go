package others

import "fmt"

const (
	WebServer       = 48.99
	Domain          = 17.79
	GiaSpecialPrice = 46.8
	FakePrice       = 113.99

	ExchangeRate = 7.08
	HumanCount   = 7
)

func main() {
	avg := FakePrice / HumanCount
	fmt.Printf("每人平均支付：%.2f\n", avg*ExchangeRate)
	fmt.Printf("我的实际支付：%.2f\n", (avg-(FakePrice-WebServer-Domain-GiaSpecialPrice))*ExchangeRate)
}
