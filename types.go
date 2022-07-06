package main

type search struct {
	Data struct {
		Posts struct {
			Items []struct {
				ID             int      `json:"id"`
				Status         bool     `json:"status"`
				AuthorUsername string   `json:"authorUsername"`
				Title          string   `json:"title"`
				City           string   `json:"city"`
				PostDate       int      `json:"postDate"`
				UpdateDate     int      `json:"updateDate"`
				HasImage       bool     `json:"hasImage"`
				ThumbURL       string   `json:"thumbURL"`
				AuthorID       int      `json:"authorId"`
				BodyTEXT       string   `json:"bodyTEXT"`
				Tags           []string `json:"tags"`
				ImagesList     []string `json:"imagesList"`
				CommentStatus  int      `json:"commentStatus"`
				CommentCount   int      `json:"commentCount"`
				UpRank         int      `json:"upRank"`
				DownRank       int      `json:"downRank"`
				GeoHash        string   `json:"geoHash"`
			} `json:"items"`
			PageInfo struct {
				HasNextPage bool `json:"hasNextPage"`
			} `json:"pageInfo"`
		} `json:"posts"`
	} `json:"data"`
}

type numText struct {
	Data struct {
		PostContact struct {
			ContactText   string `json:"contactText"`
			ContactMobile string `json:"contactMobile"`
		} `json:"postContact"`
	} `json:"data"`
}
