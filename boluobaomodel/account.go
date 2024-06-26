package boluobaomodel

type Users struct {
	Status `json:"status"`
	Data   AccountData `json:"data"`
}

type AccountData struct {
	UserName     string `json:"userName"`
	NickName     string `json:"nickName"`
	Email        string `json:"email"`
	AccountID    int    `json:"accountId"`
	RoleName     string `json:"roleName"`
	FireCoin     int    `json:"fireCoin"`
	Avatar       string `json:"avatar"`
	IsAuthor     bool   `json:"isAuthor"`
	PhoneNum     string `json:"phoneNum"`
	RegisterDate string `json:"registerDate"`
}
type Account struct {
	Status `json:"status"`
	Data   AccountData `json:"data"`
}
type AccountIp struct {
	Status `json:"status"`
	Data   AccountIpData `json:"data"`
}
type AuthorInfo struct {
	Status `json:"status"`
	Data   []BookInfoData `json:"data"`
}
type AccountIpData struct {
	IP          string `json:"ip"`
	Location    string `json:"location"`
	CountryCode int    `json:"countryCode"`
}
type AccountComment struct {
	Status `json:"status"`
	Data   []struct {
		PostDate string `json:"postDate"`
		Post     struct {
			Post struct {
				PostID      int    `json:"postId"`
				AccountID   int    `json:"accountId"`
				RepostNum   int    `json:"repostNum"`
				ReplyNum    int    `json:"replyNum"`
				FavNum      int    `json:"favNum"`
				NickName    string `json:"nickName"`
				Avatar      string `json:"avatar"`
				Content     string `json:"content"`
				PostDate    string `json:"postDate"`
				IsCanDelete bool   `json:"isCanDelete"`
				Images      []struct {
					SourceImageURL    string `json:"sourceImageUrl"`
					ThumbnailImageURL string `json:"thumbnailImageUrl"`
					SourceWidth       int    `json:"sourceWidth"`
					SourceHeight      int    `json:"sourceHeight"`
					ThumbnailWidth    int    `json:"thumbnailWidth"`
					ThumbnailHeight   int    `json:"thumbnailHeight"`
					ImageID           int    `json:"imageId"`
				} `json:"images"`
				From struct {
					Name interface{} `json:"name"`
					Type interface{} `json:"type"`
					Li   interface{} `json:"link"`
				} `json:"from"`
				IsDelete bool        `json:"isDelete"`
				Expand   interface{} `json:"expand"`
			} `json:"post"`
			Source interface{} `json:"source"`
		} `json:"post"`
		Comment     interface{} `json:"comment"`
		PostComment interface{} `json:"postComment"`
		Expand      interface{} `json:"expand"`
		Entity      interface{} `json:"entity"`
	} `json:"data"`
}
type Money struct {
	Status `json:"status"`
	Data   MoneyData `json:"data"`
}

type MoneyData struct {
	RmbCost         int `json:"rmbCost"`
	FireMoneyUsed   int `json:"fireMoneyUsed"`
	FireMoneyRemain int `json:"fireMoneyRemain"`
	VipLevel        int `json:"vipLevel"`
	CouponsRemain   int `json:"couponsRemain"`
}
