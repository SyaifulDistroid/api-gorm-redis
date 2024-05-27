package logger

type LoggerRequestData struct {
	Method string            `json:"method"`
	Path   string            `json:"path"`
	Header map[string]string `json:"header"`
	Body   string            `json:"body"`
}

type LogModel struct {
	AppName    string      `json:"app_name"`
	AppVersion string      `json:"app_version"`
	XID        string      `json:"xid"`
	Request    interface{} `json:"request"`
	Response   interface{} `json:"response"`
	LogType    string      `json:"log_type"`
}
