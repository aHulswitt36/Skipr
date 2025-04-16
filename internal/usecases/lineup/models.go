package lineup

type Player struct {
	Id int
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
	Bench 		Position="Bench"
)

type Assignment struct {
    PlayerName string
    PlayerId int
    Position Position
    Inning int
} 

type Lineup struct {
    Innings int
    Players []Player
    Defense map[int][]Assignment
    BattingOrder []Player
}
