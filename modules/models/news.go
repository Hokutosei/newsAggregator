package models

import "time"

// "_id" : ObjectId("55ee44fae897998f09c4cd24"),
// 	"by" : "GoogleNews",
// 	"id" : 0,
// 	"score" : 0,
// 	"text" : "",
// 	"time" : 1441678586,
// 	"title" : "イチローは代打で三振 外角球を見逃し",
// 	"secondary_title" : "イチローは代打で三振 外角球を見逃し",
// 	"encoded_title" : "ã¤ãã­ã¼ã¯ä»£æã§ä¸æ¯ å¤è§çãè¦éã",
// 	"type" : "",
// 	"url" : "http%3A%2F%2Fwww.sankei.com%2Fsports%2Fnews%2F150908%2Fspo1509080018-n1.html",
// 	"providername" : "GoogleNews",
// 	"providerurl" : "",
// 	"publisher" : "産経ニュース",
// 	"created_at" : ISODate("2015-09-08T02:16:26.435Z"),
// 	"relatedstories" : [ ],
// 	"category" : {
// 		"initial" : "s",
// 		"name" : "スポーツ"
// 	},
// 	"image" : {
// 		"originalcontexturl" : "",
// 		"publisher" : "",
// 		"tbheight" : 0,
// 		"tburl" : "",
// 		"tbwidth" : 0,
// 		"url" : ""
// 	},
// 	"image_url" : "",
// 	"news_page_view" : 0,
// 	"content" : "マーリンズのイチローは１－９の八回無死一、二塁から代打で出場し、見逃し三振だった。１ボール２ストライクからの外角球に球審の右手が上がり、チームと同様にフラストレーションのたまる結果となった。 試合は七回攻撃中にジェニングズ監督がストライク、ボールの判定への&nbsp;...",
// 	"language" : "jp"

// NewsMaster main data struct from mongodb
type NewsMaster struct {
	By             string    `bson:"by" json:"by"`
	ID             int       `bson:"_id" json:"_id"`
	Score          int       `bson:"score" json:"score"`
	Text           string    `bson:"text" json:"text"`
	Time           int       `bson:"time" json:"time"`
	Title          string    `bson:"title" json:"title"`
	Type           string    `bson:"type" json:"type"`
	URL            string    `bson:"url"`
	ProviderName   string    `bson:"providername" json:"providername"`
	ProviderURL    string    `bson:"providerurl"`
	CreatedAt      time.Time `bson:"created_at" json:"created_at"`
	RelatedStories []RelatedStories
	Category       TopicIdentity `json:"category"`
	Image          interface{}   `json:"image"`
	Content        string        `bson:"content" json:"content"`
	NewsPageView   int           `bson:"news_page_view"`
	ImageURL       string        `bson:"image_url" json:"image_url"`
	Language       string        `bson:"language" json:"language"`
}

// RelatedStories google related stories
type RelatedStories struct {
	Language          string `json:"language"`
	Location          string `json:"location"`
	PublishedDate     string `json:"publishedDate"`
	Publisher         string `json:"publisher"`
	SignedRedirectURL string `json:"signedRedirectUrl"`
	Title             string `json:"title"`
	TitleNoFormatting string `json:"titleNoFormatting"`
	UnescapedURL      string `json:"unescapedUrl"`
	URL               string `json:"url"`
}

//TopicIdentity topic identifier
type TopicIdentity struct {
	Initial string `json:"initial"`
	Name    string `json:"name"`
}

// Topics topics list map holder
type Topics map[string]TopicIdentity

// Image google news item top image
type Image struct {
	Publisher string `json:"publisher"`
	URL       string `json:"url"`
}
