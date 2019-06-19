package processors

// Define these types in a separate package that both my main and my plugins can import and reference

// Every plugin must be able to give me something that meets this interface
type Processor interface {
	Init(map[string]interface{}) error
	Process(buf []byte) []byte
}
