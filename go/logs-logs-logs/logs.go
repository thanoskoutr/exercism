package logs

import (
	"fmt"
	"strings"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	lineSplice := strings.Split(line, ":")
	msgSplice := lineSplice[1:]
	msg := strings.TrimSpace(strings.Join(msgSplice, " "))
	return msg
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	msg := Message(line)
	msgRune := []rune(msg)
	return len(msgRune)
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	lineSplice := strings.Split(line, ":")
	levelBrackets := lineSplice[:1][0]
	levelSplice := strings.Split(levelBrackets, "[")
	level := levelSplice[1:][0]
	levelSplice = strings.Split(level, "]")
	level = levelSplice[:1][0]
	return strings.ToLower(level)
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	msg := Message(line)
	level := LogLevel(line)
	logLine := fmt.Sprintf("%s (%s)", msg, level)
	return logLine
}
