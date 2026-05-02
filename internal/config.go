package internal

import (
	"log-generator/pkg"
	"strings"
)

type LogEntry struct {
	Level string
	Msg   string
}

type Config struct {
	Entries []LogEntry
	DelayMs int
	Service string
}

func LoadConfig() Config {
	cfg := pkg.GetEnv("LOG_CONFIG", "INFO:welcome tester,WARN:don't push Node modules")
	delay := pkg.GetEnvInt("LOG_DELAY_MS", 2000)
	service := pkg.GetEnv("SERVICE_NAME", "demo-log-service")

	return Config{
		Entries: parse(cfg),
		DelayMs: delay,
		Service: service,
	}
}

func parse(cfg string) []LogEntry {
	var out []LogEntry

	for item := range strings.SplitSeq(cfg, ",") {
    	parts := strings.SplitN(item, ":", 2)
    	if len(parts) != 2 {
        	continue
    	}

		out = append(out, LogEntry{
			Level: strings.TrimSpace(parts[0]),
			Msg:   strings.TrimSpace(parts[1]),
		})
	}
	return out
}