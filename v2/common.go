package applog

// Logging message
type Logger interface {
	Always(v interface{})
	Info(v interface{})
	Debug(v interface{})
}

// Logging to real world
type Output interface {
	Println(v ...interface{})
}

type config struct {
	output Output
}

func (config config) Println(v interface{}) {
	config.output.Println(v)
}
