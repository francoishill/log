package log

// Interface represents the API of both Logger and Entry as per RFC5424
type Interface interface {
	WithFields(fields Fielder) *Entry
	WithField(key string, value interface{}) *Entry
	WithError(err error) *Entry

	Debug(msg string)
	Info(msg string)
	Notice(msg string)
	Warn(msg string)
	Error(msg string)
	Critical(msg string)
	Alert(msg string)
	Emergency(msg string)

	Debugf(msg string, v ...interface{})
	Infof(msg string, v ...interface{})
	Noticef(msg string, v ...interface{})
	Warnf(msg string, v ...interface{})
	Errorf(msg string, v ...interface{})
	Criticalf(msg string, v ...interface{})
	Alertf(msg string, v ...interface{})
	Emergencyf(msg string, v ...interface{})
	Trace(msg string) *Entry
}
