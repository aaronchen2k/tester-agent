package managerModel

type Config struct {
	AgentVersion string

	Interval int64
	Language string
}

func NewConfig() Config {
	return Config{
		Interval: 6,
		Language: "en",
	}
}
