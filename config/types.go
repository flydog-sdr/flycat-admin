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
		Preference string `json:"preference"`
	} `json:"flycat_settings"`
}

type Preference struct {
	Gain    int `json:"gain"`
	Filter  int `json:"filter"`
	Antenna int `json:"antenna"`
}
