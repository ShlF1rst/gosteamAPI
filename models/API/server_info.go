package API

import "fmt"

type ServerInfo struct {
	ServerTime       uint   `json:"servertime"`
	ServerTimeString string `json:"servertimestring"`
}

func (serverInfo ServerInfo) String() string {
	return fmt.Sprintf("Server info\n\tTime unix: %d\n\tTime string: %s", serverInfo.ServerTime, serverInfo.ServerTimeString)
}
