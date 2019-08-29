package web

// Defaults
var (
	defaultDBURL         = "redis://localhost"
	defaultDBPort        = "6379"
	defaultInputChannel  = "message"
	defaultOutputChannel = "values"
)

// Config type
type Config struct {
	DBURL         string
	DBPort        string
	InputChannel  string
	OutputChannel string
}

// Create config
func NewConfig(dbURL, dbPort, inChan, outChan string) *Config {
	return &Config{
		DBURL:         dbURL,
		DBPort:        dbPort,
		InputChannel:  inChan,
		OutputChannel: outChan,
	}
}

// Default config
func DefaultConfig() *Config {
	return NewConfig(defaultDBURL, defaultDBPort, defaultInputChannel, defaultOutputChannel)
}
