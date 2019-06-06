package talk

import "rsc.io/quote"

func SayHello() string {
	quote.Hello()
	return "Hello, world."
}
