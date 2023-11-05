package status

import "github.com/flydog-sdr/flycat-admin/config"

func GetController(pref *config.Preference) Controller {
	return Controller{
		Gain:    pref.Gain,
		Filter:  pref.Filter,
		Antenna: pref.Antenna,
	}
}
