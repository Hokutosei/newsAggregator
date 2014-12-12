package news_getter

import(

)

type jsonNewsBody struct {
	By				string
	Id				int
	//Kids 			[]int
	Score			int
	Text			string
	Time			int
	Title			string
	Type			string
	Url				string
	ProviderName	string
	ProviderUrl		string
	CreatedAt		string
	RelatedStories	[]RelatedStories
}
