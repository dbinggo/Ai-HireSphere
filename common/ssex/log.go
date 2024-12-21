package ssex

type Log interface {
	Info(...any)
	Infof(string, ...any)
	Debug(...any)
	Debugf(string, ...any)
	Warn(...any)
	Warnf(string, ...any)
	Error(...any)
	Errorf(string, ...any)
}
