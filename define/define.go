package define

const (
	HTML_INDEX string = "index.html"
)

type LogConfig struct {
	FileName string
}

type FrontEndServerConfig struct {
	LogCfg LogConfig
}
