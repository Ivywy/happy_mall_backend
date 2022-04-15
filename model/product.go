package model

type Product struct {
	ProductId            string `json:"productId" gorm:"colum:product_id"`
	ProductName          string `json:"productName" gorm:"colum:product_name"`
	ProductIntro         string `json:"productIntro" gorm:"colum:product_intro"`
	CategoryId           string `json:"categoryId" gorm:"colum:category_id"`
	ProductCoverImg      string `json:"productCoverImg" gorm:"colum:product_cover_img"`
	ProductBanner        string `json:"productBanner" gorm:"colum:product_banner"`
	OriginalPrice        string `json:"originalPrice" gorm:"colum:original_price"`
	SellerPrice          string `json:"sellerPrice" gorm:"colum:seller_price"`
	StockNum             string `json:"stockNum" gorm:"colum:stock_num"`
	Tag                  string `json:"tag" gorm:"colum:tag"`
	SellerStatus         string `json:"sellerStatus" gorm:"colum:seller_status"`
	CreateUser           string `json:"createUser" gorm:"colum:create_user"`
	UpdateUser           string `json:"updateUser" gorm:"colum:update_user"`
	ProductDetailContent string `json:"productDetailContent" gorm:"colum:product_detail_content"`
	IsDeletes            bool   `json:"isDeleted" gorm:"colum:product_detal_content"`
}
