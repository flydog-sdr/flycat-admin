package controller

type Controller struct{}

type Binding struct {
	Gain    int `form:"gain" json:"gain" xml:"gain"`
	Antenna int `form:"antenna" json:"antenna" xml:"antenna"`
	Filter  int `form:"filter" json:"filter" xml:"filter"`
}
