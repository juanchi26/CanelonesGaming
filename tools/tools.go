package tools

import (
	"fmt"
	"strings"
	"time"
)

func FechaMySQL() string { //formatea la hora para mysql
	t := time.Now()

	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
}

func EscapeString(t string) string {
	desc := strings.ReplaceAll(t, "'", "")
	desc = strings.ReplaceAll(desc, "\"", "")
	return desc
}
