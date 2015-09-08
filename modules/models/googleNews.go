package models

// GoogleNewsResults google news result struct
type GoogleNewsResults struct {
	GsearchResultClass string `json:"GsearchResultClass"`
	ClusterURL         string `json:"clusterUrl"`
	Content            string `json:"content"`
	Image              struct {
		OriginalContextURL string `json:"originalContextUrl"`
		Publisher          string `json:"publisher"`
		TbHeight           int    `json:"tbHeight"`
		TbURL              string `json:"tbUrl"`
		TbWidth            int    `json:"tbWidth"`
		URL                string `json:"url"`
	} `json:"image"`
	Language          string           `json:"language"`
	Location          string           `json:"location"`
	PublishedDate     string           `json:"publishedDate"`
	Publisher         string           `json:"publisher"`
	RelatedStories    []RelatedStories `json:"relatedStories"`
	SignedRedirectURL string           `json:"signedRedirectUrl"`
	Title             string           `json:"title"`
	TitleNoFormatting string           `json:"titleNoFormatting"`
	UnescapedURL      string           `json:"unescapedUrl"`
	URL               string           `json:"url"`
	Category          TopicIdentity
}
