package env

var devConfig = config{Port: 8080}

type config struct {
	Port int
}

func GetConfig() config {
	return devConfig
}
