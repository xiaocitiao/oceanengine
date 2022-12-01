package enum

// ProductSetting 商品库设置
type ProductSetting string

const (
	// ProductSetting_SINGLE 启用SDPA
	ProductSetting_SINGLE ProductSetting = "SINGLE"
	// ProductSetting_NO_MAP 不启用
	ProductSetting_NO_MAP ProductSetting = "NO_MAP"
)
