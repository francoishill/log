package log

import stdlog "log"

// assert interface compliance.
var _ Interface = (*Logger)(nil)

// Fielder is an interface for providing fields to custom types.
type Fielder interface {
	Fields() Fields
}

// Fields represents a map of entry level data used for structured logging.
type Fields map[string]interface{}

// Fields implements Fielder.
func (f Fields) Fields() Fields {
	return f
}

// The HandlerFunc type is an adapter to allow the use of ordinary functions as
// log handlers. If f is a function with the appropriate signature,
// HandlerFunc(f) is a Handler object that calls f.
type HandlerFunc func(*Entry) error

// HandleLog calls f(e).
func (f HandlerFunc) HandleLog(e *Entry) error {
	return f(e)
}

// Handler is used to handle log events, outputting them to
// stdio or sending them to remote services. See the "handlers"
// directory for implementations.
//
// It is left up to Handlers to implement thread-safety.
type Handler interface {
	HandleLog(*Entry) error
}

// Logger represents a logger with configurable Level and Handler.
type Logger struct {
	Handler Handler
	Level   Level
}

// WithFields returns a new entry with `fields` set.
func (l *Logger) WithFields(fields Fielder) *Entry {
	return NewEntry(l).WithFields(fields.Fields())
}

// WithField returns a new entry with the `key` and `value` set.
func (l *Logger) WithField(key string, value interface{}) *Entry {
	return NewEntry(l).WithField(key, value)
}

// WithError returns a new entry with the "error" set to `err`.
func (l *Logger) WithError(err error) *Entry {
	return NewEntry(l).WithError(err)
}

// Debug level message.
func (l *Logger) Debug(msg string) {
	NewEntry(l).Debug(msg)
}

// Info level message.
func (l *Logger) Info(msg string) {
	NewEntry(l).Info(msg)
}

// Notice level message.
func (l *Logger) Notice(msg string) {
	NewEntry(l).Notice(msg)
}

// Warn level message.
func (l *Logger) Warn(msg string) {
	NewEntry(l).Warn(msg)
}

// Error level message.
func (l *Logger) Error(msg string) {
	NewEntry(l).Error(msg)
}

// Critical level message, followed by an exit.
func (l *Logger) Critical(msg string) {
	NewEntry(l).Critical(msg)
}

// Alert level message, followed by an exit.
func (l *Logger) Alert(msg string) {
	NewEntry(l).Alert(msg)
}

// Emergency level message, followed by an exit.
func (l *Logger) Emergency(msg string) {
	NewEntry(l).Emergency(msg)
}

// Debugf level formatted message.
func (l *Logger) Debugf(msg string, v ...interface{}) {
	NewEntry(l).Debugf(msg, v...)
}

// Infof level formatted message.
func (l *Logger) Infof(msg string, v ...interface{}) {
	NewEntry(l).Infof(msg, v...)
}

// Noticef level formatted message.
func (l *Logger) Noticef(msg string, v ...interface{}) {
	NewEntry(l).Noticef(msg, v...)
}

// Warnf level formatted message.
func (l *Logger) Warnf(msg string, v ...interface{}) {
	NewEntry(l).Warnf(msg, v...)
}

// Errorf level formatted message.
func (l *Logger) Errorf(msg string, v ...interface{}) {
	NewEntry(l).Errorf(msg, v...)
}

// Criticalf level formatted message, followed by an exit.
func (l *Logger) Criticalf(msg string, v ...interface{}) {
	NewEntry(l).Criticalf(msg, v...)
}

// Alertf level formatted message, followed by an exit.
func (l *Logger) Alertf(msg string, v ...interface{}) {
	NewEntry(l).Alertf(msg, v...)
}

// Emergencyf level formatted message, followed by an exit.
func (l *Logger) Emergencyf(msg string, v ...interface{}) {
	NewEntry(l).Emergencyf(msg, v...)
}

// Trace returns a new entry with a Stop method to fire off
// a corresponding completion log, useful with defer.
func (l *Logger) Trace(msg string) *Entry {
	return NewEntry(l).Trace(msg)
}

// log the message, invoking the handler. We clone the entry here
// to bypass the overhead in Entry methods when the level is not
// met.
func (l *Logger) log(level Level, e *Entry, msg string) {
	if level < l.Level {
		return
	}

	if err := l.Handler.HandleLog(e.finalize(level, msg)); err != nil {
		stdlog.Printf("error logging: %s", err)
	}
}
