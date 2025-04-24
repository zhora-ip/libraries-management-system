package kafkaservice

type Config struct {
	Topic       string `yaml:"topic"`
	Server      string `yaml:"server"`
	OffsetReset string `yaml:"offset_reset"`
	LogLevel    int    `yaml:"log_level"`
	GroupID     string `yaml:"group_id"`
}
