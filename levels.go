package log

import (
	"errors"
)

// Level of severity.
type Level int

// Log levels.
const (
	DebugLevel Level = iota
	InfoLevel
	NoticeLevel
	WarnLevel
	ErrorLevel
	CriticalLevel
	AlertLevel
	EmergencyLevel
)

var levelNames = [...]string{
	DebugLevel:     "debug",
	InfoLevel:      "info",
	NoticeLevel:    "notice",
	WarnLevel:      "warn",
	ErrorLevel:     "error",
	CriticalLevel:  "critical",
	AlertLevel:     "alert",
	EmergencyLevel: "emergency",
}

// String implements io.Stringer.
func (l Level) String() string {
	return levelNames[l]
}

// MarshalJSON returns the level string.
func (l Level) MarshalJSON() ([]byte, error) {
	return []byte(`"` + l.String() + `"`), nil
}

// ParseLevel parses level string.
func ParseLevel(s string) (Level, error) {
	switch s {
	case "debug":
		return DebugLevel, nil
	case "info":
		return InfoLevel, nil
	case "notice":
		return NoticeLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "error":
		return ErrorLevel, nil
	case "critical":
		return CriticalLevel, nil
	case "alert":
		return AlertLevel, nil
	case "emergency":
		return EmergencyLevel, nil
	default:
		return -1, errors.New("invalid level")
	}
}
