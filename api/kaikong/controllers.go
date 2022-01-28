package kaikong

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kjk/betterguid"
)

type Live struct {
	ID          string `json:"id"`
	CreatedTime string `json:"created_time"`
	Channel     string `json:"channel"`
	URL         string `json:"url"`
	Status      string `json:"status"`
}

type Setting struct {
	TrackInventory bool `json:"track_inventory"`
	UseVariant     bool `json:"use_variant"`
}

type CoverURL struct {
	URL     string `json:"url"`
	Arrange int    `json:"arrange"`
}

type Inventory struct {
	SKU      string `json:"sku"`
	Status   string `json:"status"`
	Quantity int    `json:"quantity"`
	SoldOut  int    `json:"sold_out"`
}

type Variants struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Price      int       `json:"price"`
	LiveStatus string    `json:"live_status"`
	Inventory  Inventory `json:"inventory"`
}

type Product struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	LiveStatus string     `json:"live_status"`
	Status     string     `json:"status"`
	Setting    Setting    `json:"setting"`
	CoverURL   CoverURL   `json:"cover_url"`
	Variants   []Variants `json:"variants"`
}

type CreateLivePayload struct {
	ProductVariants []string `json:"product_variants"`
	Lives           []string `json:"lives"`
}

type SocialLive struct {
	ID           string `json:"id"`
	Channel      string `json:"channel"`
	SocialLiveID string `json:"social_lives_id"`
	URL          string `json:"url"`
	Status       string `json:"status"`
	CommentCount int    `json:"comment_count"`
	OrderCount   int    `json:"order_count"`
	StartedDate  string `json:"started_date"`
	StopedDate   string `json:"stoped_date"`
}

type ChannelLive struct {
	ID          string       `json:"id"`
	SocialLives []SocialLive `json:"social_lives"`
	CreatedDate string       `json:"created_date"`
}

type LiveCommentMember struct {
	ID         string `json:"id"`
	Social     string `json:"social"`
	SocialID   string `json:"social_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	PictureURL string `json:"picture_url"`
}
type LiveCommentReplyPayload struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
	Likes     int    `json:"likes"`
}
type LiveCommentPayloadReply struct {
	ID         string                  `json:"id"`
	Name       string                  `json:"name"`
	PictureURL string                  `json:"picture_url"`
	Payload    LiveCommentReplyPayload `json:"payload"`
}
type LiveCommentPayload struct {
	ID        string                    `json:"id"`
	Message   string                    `json:"message"`
	AttachURL string                    `json:"attach_url"`
	Reply     []LiveCommentPayloadReply `json:"reply"`
	Likes     int                       `json:"likes"`
	CreatedAt string                    `json:"created_at"`
}
type SendAllInvoicesPayload struct {
	RequestDate string `json:"request_date"`
}
type SendAllInvoicesCustomers struct {
	ID     string `json:"id"`
	Status bool   `json:"status"`
}

func GetAvailableLives(c *gin.Context) {
	lives := []Live{
		{
			ID:          "106263458498148",
			CreatedTime: "2021-10-06T10:49:02+0000",
			Channel:     "facebook",
			URL:         "https://www.facebook.com/106217388502755/videos/425246279254265/",
			Status:      "LIVE",
		},
		{
			ID:          "106263458498149",
			CreatedTime: "2021-10-07T10:49:02+0000",
			Channel:     "facebook",
			URL:         "https://www.facebook.com/106217388502755/videos/425246279254265/",
			Status:      "LIVE",
		},
	}
	fmt.Println("[debug]lives", lives)
	c.JSON(http.StatusOK, gin.H{
		"lives": lives,
	})
}

func SearchProducts(c *gin.Context) {
	products := []Product{
		{
			ID:         "kkp-a9f9409f-2b8b-4f4e-84f8-aa8ba7c653d1",
			Name:       "กางเกง",
			LiveStatus: "stopped",
			Status:     "insotck",
			Setting: Setting{
				TrackInventory: false,
				UseVariant:     false,
			},
			CoverURL: CoverURL{
				URL:     "https://assets.adidas.com/images/h_840,f_auto,q_auto:sensitive,fl_lossy,c_fill,g_auto/465a7f734f554e6ea870ab69000b03d7_9366/UB_Cargo_Pants_Beige_GL0395_21_model.jpg",
				Arrange: 0,
			},
			Variants: []Variants{
				{
					ID:         "kkpv-6c18908c-51a7-4da7-81d7-008a2d143aee",
					Name:       "กางเกง",
					Code:       "a001",
					Price:      500,
					LiveStatus: "stopped",
					Inventory: Inventory{
						SKU:      "",
						Status:   "insotck",
						Quantity: 0,
						SoldOut:  0,
					},
				},
			},
		},
		{
			ID:         "kkp-a9f9409f-2b8b-4f4e-84f8-aa8ba7c653d2",
			Name:       "เสื้อ",
			LiveStatus: "stopped",
			Status:     "insotck",
			Setting: Setting{
				TrackInventory: false,
				UseVariant:     false,
			},
			CoverURL: CoverURL{
				URL:     "https://m.media-amazon.com/images/I/319Zx40AhUL.jpg",
				Arrange: 0,
			},
			Variants: []Variants{
				{
					ID:         "kkpv-6c18908c-51a7-4da7-81d7-008a2d143aef",
					Name:       "M",
					Code:       "s001",
					Price:      799,
					LiveStatus: "stopped",
					Inventory: Inventory{
						SKU:      "",
						Status:   "insotck",
						Quantity: 0,
						SoldOut:  0,
					},
				},
				{
					ID:         "kkpv-6c18908c-51a7-4da7-81d7-008a2d143aeg",
					Name:       "XL",
					Code:       "s002",
					Price:      899,
					LiveStatus: "stopped",
					Inventory: Inventory{
						SKU:      "",
						Status:   "insotck",
						Quantity: 0,
						SoldOut:  0,
					},
				},
			},
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"currentPage":  1,
		"pages":        1,
		"currentCount": len(products),
		"totalCount":   len(products),
		"products":     products,
	})
}

func CreateLive(c *gin.Context) {
	var input CreateLivePayload
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       "kkl-301b9c38-0245-4a4f-97c6-d280150ea936",
		"products": input.ProductVariants,
		"lives":    input.Lives,
	})
}

func GetLiveList(c *gin.Context) {
	socialLive := []SocialLive{
		{
			ID:           "kklc-301b9c38-0245-4a4f-97c6-d280150ea937",
			Channel:      "facebook",
			SocialLiveID: "106263458498148",
			URL:          "https://www.google.co.th/",
			Status:       "live",
			CommentCount: 1102,
			OrderCount:   201,
			StartedDate:  "2021-10-11T05:25:00Z",
			StopedDate:   "2021-10-11T11:25:00Z",
		},
		{
			ID:           "kklc-301b9c38-0245-4a4f-97c6-d280150ea938",
			Channel:      "instagram",
			SocialLiveID: "106263458498148",
			URL:          "https://www.google.co.th/",
			Status:       "live",
			CommentCount: 1102,
			OrderCount:   201,
			StartedDate:  "2021-10-11T05:25:00Z",
			StopedDate:   "2021-10-11T11:25:00Z",
		},
	}
	socialLiveIG := []SocialLive{
		{
			ID:           "kklc-301b9c38-0245-4a4f-97c6-d280150ea922",
			Channel:      "instagram",
			SocialLiveID: "106263458498122",
			URL:          "https://www.google.co.th/",
			Status:       "live",
			CommentCount: 1102,
			OrderCount:   201,
			StartedDate:  "2021-10-10T05:25:00Z",
			StopedDate:   "2021-10-10T11:25:00Z",
		},
	}
	socialLiveFB := []SocialLive{
		{
			ID:           "kklc-301b9c38-0245-4a4f-97c6-d280150ea912",
			Channel:      "facebook",
			SocialLiveID: "106263458498112",
			URL:          "https://www.google.co.th/",
			Status:       "live",
			CommentCount: 1102,
			OrderCount:   201,
			StartedDate:  "2021-10-10T05:25:00Z",
			StopedDate:   "2021-10-10T11:25:00Z",
		},
	}
	channelLive := []ChannelLive{
		{
			ID:          "kklc-301b9c38-0245-4a4f-97c6-d280150ea936",
			SocialLives: socialLive,
			CreatedDate: "2021-10-11T05:25:00Z",
		},
		{
			ID:          "kklc-301b9c38-0245-4a4f-97c6-d280150ea921",
			SocialLives: socialLiveIG,
			CreatedDate: "2021-10-10T05:25:00Z",
		},
		{
			ID:          "kklc-301b9c38-0245-4a4f-97c6-d280150ea911",
			SocialLives: socialLiveFB,
			CreatedDate: "2021-10-10T05:25:00Z",
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"lives": channelLive,
	})
}

func GetLiveInfo(c *gin.Context) {
	liveId := c.Param("live_id")
	socialLive := []SocialLive{
		{
			ID:           "kklc-301b9c38-0245-4a4f-97c6-d280150ea912",
			Channel:      "facebook",
			SocialLiveID: "106263458498112",
			URL:          "https://www.google.co.th/",
			Status:       "live",
			CommentCount: 1102,
			OrderCount:   201,
			StartedDate:  "2021-10-10T05:25:00Z",
			StopedDate:   "2021-10-10T11:25:00Z",
		},
		{
			ID:           "kklc-301b9c38-0245-4a4f-97c6-d280150ea922",
			Channel:      "instagram",
			SocialLiveID: "206263458498112",
			URL:          "https://www.google.co.th/",
			Status:       "live",
			CommentCount: 1102,
			OrderCount:   201,
			StartedDate:  "2021-10-10T05:25:00Z",
			StopedDate:   "2021-10-10T11:25:00Z",
		},
	}
	liveProduct := []Variants{
		{
			ID:         "kkpv-6c18908c-51a7-4da7-81d7-008a2d143aee",
			Name:       "กางเกง",
			Code:       "a001",
			Price:      500,
			LiveStatus: "stopped",
			Inventory: Inventory{
				SKU:      "",
				Status:   "insotck",
				Quantity: 0,
				SoldOut:  0,
			},
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"lives": gin.H{
			"id": liveId,
			"setting": map[string]bool{
				"auto_cf": true,
			},
			"social_lives":  socialLive,
			"live_products": liveProduct,
			"created_date":  "2021-10-10T05:25:00Z",
		},
	})
}

func GetLiveComments(c *gin.Context) {
	fmt.Println("[debug]", c.Request.Header)
	chanStream := make(chan int, 2)
	go func() {
		defer close(chanStream)
		for i := 0; i < 5; i++ {
			min := 1
			max := 5
			sec := rand.Intn(max-min) + min
			chanStream <- i
			time.Sleep(time.Duration(sec) * time.Second)
		}
	}()
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			fmt.Println("[debug]msg", msg)
			// if msg < 4 {
			// 	reply := []LiveCommentPayloadReply{
			// 		{
			// 			ID:         "user_id",
			// 			Name:       "name",
			// 			PictureURL: "picture_url",
			// 			Payload: LiveCommentReplyPayload{
			// 				ID:        "comment_id",
			// 				Message:   "99",
			// 				CreatedAt: "",
			// 				Likes:     0,
			// 			},
			// 		},
			// 	}
			// 	c.SSEvent("message", gin.H{
			// 		"member": LiveCommentMember{
			// 			ID:         "kkm-301b9c38-0245-4a4f-97c6-d280150ea936",
			// 			Social:     "facebook",
			// 			SocialID:   "BeNzfunnii",
			// 			Name:       "Somchai Kaidee",
			// 			Email:      "test@test.com",
			// 			Phone:      "66--0651234567",
			// 			ProfileURL: "https://i.pravatar.cc/150?img=52",
			// 		},
			// 		"payload": LiveCommentPayload{
			// 			ID:          "comment_id",
			// 			Message:     "test 🎉",
			// 			AttachURL:   "https://scontent-sin6-4.xx.fbcdn.net/v/t1.6435-9/s720x720/242127503_10216409718793867_8372515284152829181_n.jpg?_nc_cat=101&ccb=1-5&_nc_sid=1480c5&_nc_eui2=AeEjxPDc7uv0z3b3_RT-BMmEQhVSAjlnEjhCFVICOWcSONyU8rPf1Lw9cP8YW5slK6o&_nc_ohc=kTaCsLnMqlkAX8HoK7Z&_nc_ht=scontent-sin6-4.xx&edm=AIjQbYoEAAAA&oh=186ca431403cd311dec6f01c184ddea4&oe=616F5D6E",
			// 			Reply:       reply,
			// 			CreatedDate: "yyy-MM-dd'T'HH:mm:ssZ",
			// 		},
			// 	})
			// } else {
			// 	c.SSEvent("message", "STOPPED")
			// }
			reply := []LiveCommentPayloadReply{
				{
					ID:         "user_id",
					Name:       "name",
					PictureURL: "picture_url",
					Payload: LiveCommentReplyPayload{
						ID:        "comment_id",
						Message:   "99",
						CreatedAt: "",
						Likes:     0,
					},
				},
			}
			// c.SSEvent("message", gin.H{
			// 	"member": LiveCommentMember{
			// 		ID:         "kkm-301b9c38-0245-4a4f-97c6-d280150ea936",
			// 		Social:     memberSocial,
			// 		SocialID:   "BeNzfunnii",
			// 		Name:       "Somchai Kaidee",
			// 		Email:      "test@test.com",
			// 		Phone:      "66--0651234567",
			// 		ProfileURL: "https://i.pravatar.cc/150?img=52",
			// 	},
			// 	"payload": LiveCommentPayload{
			// 		ID:          betterguid.New(),
			// 		Message:     "test 🎉",
			// 		AttachURL:   "https://scontent-sin6-4.xx.fbcdn.net/v/t1.6435-9/s720x720/242127503_10216409718793867_8372515284152829181_n.jpg?_nc_cat=101&ccb=1-5&_nc_sid=1480c5&_nc_eui2=AeEjxPDc7uv0z3b3_RT-BMmEQhVSAjlnEjhCFVICOWcSONyU8rPf1Lw9cP8YW5slK6o&_nc_ohc=kTaCsLnMqlkAX8HoK7Z&_nc_ht=scontent-sin6-4.xx&edm=AIjQbYoEAAAA&oh=186ca431403cd311dec6f01c184ddea4&oe=616F5D6E",
			// 		Reply:       reply,
			// 		CreatedDate: "yyy-MM-dd'T'HH:mm:ssZ",
			// 	},
			// })
			c.SSEvent("message", gin.H{
				"member": LiveCommentMember{
					ID:         "kkm-301b9c38-0245-4a4f-97c6-d280150ea936",
					SocialID:   "106217388502755",
					Name:       "Shop test kaikong",
					Email:      "test@test.com",
					Phone:      "66--0651234567",
					PictureURL: "https://scontent.fcjb7-1.fna.fbcdn.net/v/t1.6435-1/cp0/c12.0.50.50a/p50x50/244684106_106217791836048_2067742612478830836_n.jpg?_nc_cat=101&ccb=1-5&_nc_sid=dbb9e7&_nc_ohc=NIceFdcOidsAX8hkNFa&_nc_ht=scontent.fcjb7-1.fna&edm=AJ8zDikEAAAA&oh=f36b0e3b52fb082ee7c9bd54302b5ad1&oe=618CF82B",
				},
				"payload": LiveCommentPayload{
					ID:        betterguid.New(),
					Message:   "ดีจ๊ะ 🎉",
					AttachURL: "https://scontent-sin6-4.xx.fbcdn.net/v/t1.6435-9/s720x720/242127503_10216409718793867_8372515284152829181_n.jpg?_nc_cat=101&ccb=1-5&_nc_sid=1480c5&_nc_eui2=AeEjxPDc7uv0z3b3_RT-BMmEQhVSAjlnEjhCFVICOWcSONyU8rPf1Lw9cP8YW5slK6o&_nc_ohc=kTaCsLnMqlkAX8HoK7Z&_nc_ht=scontent-sin6-4.xx&edm=AIjQbYoEAAAA&oh=186ca431403cd311dec6f01c184ddea4&oe=616F5D6E",
					Reply:     reply,
					// Reply:       [],
					CreatedAt: "2021-10-14T09:24:34+0000",
					Likes:     0,
				},
			})
			return true
		}
		return false
	})
}

func SendLiveMessage(c *gin.Context) {
	liveId := c.Param("live_id")
	id := c.PostForm("id")
	channel := c.PostFormArray("channel")
	message := c.PostForm("message")

	c.JSON(http.StatusOK, gin.H{
		"id":       liveId,
		"channels": channel,
		"payloads": map[string]string{
			"id":           id,
			"message":      message,
			"created_date": time.Now().Format(time.RFC3339),
		},
	})
}

func SendAllInvoices(c *gin.Context) {
	var input SendAllInvoicesPayload
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customers := []SendAllInvoicesCustomers{
		{
			ID:     "customers_1",
			Status: true,
		},
		{
			ID:     "customers_2",
			Status: true,
		},
	}

	time.Sleep(2 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"customers":    customers,
		"request_date": input.RequestDate,
	})
}
