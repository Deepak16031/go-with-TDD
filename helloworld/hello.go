package main

const (
	spanish            = "Spanish"
	french             = "French"
	helloEnglishPrefix = "Hello, "
	helloSpanishPrefix = "Hola, "
	helloFrenchPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if len(name) == 0 {
		name = "World"
	}
	return getGreeting(language) + name
}

func getGreeting(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloSpanishPrefix
	case french:
		prefix = helloFrenchPrefix
	default:
		prefix = helloEnglishPrefix
	}
	return
}

func main() {
	//fmt.Println(Hello())
}
