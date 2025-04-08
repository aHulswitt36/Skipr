package lineup

type Player struct {
	Id string
	Name string
	SkillTag string
}

type Position string 

const (
	Pitcher		Position="P"
	Catcher 	Position="C"
	FirstBase	Position="1B"
	SecondBase	Position="2B"
	Shortstop	Position="SS"
	ThirdBase	Position="3B"
	ShortField	Position="SF"
	LeftField	Position="LF"
	CenterField	Position="CF"
	RightField	Position="RF"
	LeftCenter	Position="LCF"
	RightCenter	Position="RCF"
)
