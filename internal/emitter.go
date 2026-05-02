package internal

import (
	"fmt"
	"time"
)

func StartEmitter(cfg Config) {
	for _, entry := range cfg.Entries {
		go emit(entry, cfg)
	}
}

func emit(e LogEntry, cfg Config) {
	for {
		fmt.Printf(`%s  %s  %s : %s`+"\n",
			e.Level,
			time.Now().Format(time.RFC3339),
			cfg.Service,
			e.Msg,
		)

		time.Sleep(time.Duration(cfg.DelayMs) * time.Millisecond)
	}
}