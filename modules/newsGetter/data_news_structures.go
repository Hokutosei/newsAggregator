package news_getter

type jsonNewsBody struct {
	By string
	Id int
	//Kids 			[]int
	Score          int
	Text           string
	Time           int
	Title          string
	Type           string
	Url            string
	ProviderName   string
	ProviderUrl    string
	CreatedAt      string
	RelatedStories []RelatedStories
	Category       TopicIdentity
}

// Topics topics list map holder
type Topics map[string]TopicIdentity

//TopicIdentity topic identifier
type TopicIdentity struct {
	Initial string
	Name    string
}
