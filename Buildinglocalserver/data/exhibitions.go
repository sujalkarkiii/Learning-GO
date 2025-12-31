package data

type Exibition struct{
	Title string
	Description string
	Image string
}
func Getall() []Exibition{
	return  list
}

func Add(exhi Exibition){
	list=append(list,exhi)
}

var list=[]Exibition{
	{
		Title: "purwanchal campus",
		Description: "located in dhatan",
		Image: "ancient-greece.png",
	},	{
		Title: "purwanchal cgfdgsdsgampus",
		Description: "located in dhatan",
		Image: "ancient-greece.png",
	},	{
		Title: "purwanchal camsgdfsgdgdfpus",
		Description: "located in dhatan",
		Image: "ancient-greece.png",
	},

}