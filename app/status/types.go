package status

type Status struct{}

type Memory struct {
	Total   uint64  `json:"total"`
	Free    uint64  `json:"free"`
	Used    uint64  `json:"used"`
	Percent float64 `json:"percent"`
}

type OS struct {
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Distro   string `json:"distro"`
	Hostname string `json:"hostname"`
}

type CPU struct {
	Model   string  `json:"model"`
	Percent float64 `json:"percent"`
}

type Controller struct {
	Gain    int `json:"gain"`
	Filter  int `json:"filter"`
	Antenna int `json:"antenna"`
}

type System struct {
	Uptime     int64      `json:"uptime"`
	Memory     Memory     `json:"memory"`
	OS         OS         `json:"os"`
	CPU        CPU        `json:"cpu"`
	Controller Controller `json:"controller"`
}
