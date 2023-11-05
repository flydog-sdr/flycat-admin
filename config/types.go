package config

type Args struct {
	Path    string
	Version bool
}

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server_settings"`
	FlyCat struct {
		Volume     string `json:"volume"`
		Preference string `json:"preference"`
	} `json:"flycat_settings"`
	Docker struct {
		Socket   string `json:"socket"`
		Registry string `json:"registry"`
	} `json:"docker_settings"`
	Script struct {
		Update   string `json:"update"`
		Reset    string `json:"reset"`
		Reboot   string `json:"reboot"`
		Shutdown string `json:"shutdown"`
	} `json:"script_settings"`
}

type Preference struct {
	Gain    int `json:"gain"`
	Filter  int `json:"filter"`
	Antenna int `json:"antenna"`
}
