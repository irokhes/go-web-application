package viewmodel

type Home struct {
	Title  string
	Active string
}

func NewHome() Home {
	return Home{
		Title:  "Lemonade Standard Supply",
		Active: "home",
	}
}
