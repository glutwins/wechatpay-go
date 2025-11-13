package productcoupon

import "time"

type SendUserProductCouponRequest struct {
	ProductCouponId string         `json:"product_coupon_id,omitempty"`
	StockId         string         `json:"stock_id,omitempty"`
	CouponCode      string         `json:"coupon_code,omitempty"`
	Appid           string         `json:"appid,omitempty"`
	SendRequestNo   string         `json:"send_request_no,omitempty"`
	Attach          string         `json:"attach,omitempty"`
	BrandId         string         `json:"brand_id,omitempty"`
	CouponTagInfo   *CouponTagInfo `json:"coupon_tag_info,omitempty"`
}

type UserProductCouponEntity struct {
	CouponCode            string                 `json:"coupon_code,omitempty"`
	CouponState           string                 `json:"coupon_state,omitempty"`
	ValidBeginTime        time.Time              `json:"valid_begin_time,omitempty"`
	ValidEndTime          time.Time              `json:"valid_end_time,omitempty"`
	ReceiveTime           time.Time              `json:"receive_time,omitempty"`
	SendRequestNo         string                 `json:"send_request_no,omitempty"`
	SendChannel           string                 `json:"send_channel,omitempty"`
	ConfirmRequestNo      string                 `json:"confirm_request_no,omitempty"`
	ConfirmTime           time.Time              `json:"confirm_time,omitempty"`
	DeactivateRequestNo   string                 `json:"deactivate_request_no,omitempty"`
	DeactivateTime        time.Time              `json:"deactivate_time,omitempty"`
	DeactivateReason      string                 `json:"deactivate_reason,omitempty"`
	SingleUsageDetail     *SingleUsageDetail     `json:"single_usage_detail,omitempty"`
	SequentialUsageDetail *SequentialUsageDetail `json:"sequential_usage_detail,omitempty"`
	ProductCoupon         *ProductCouponEntity   `json:"product_coupon,omitempty"`
	Stock                 *StockEntity           `json:"stock,omitempty"`
	Attach                string                 `json:"attach,omitempty"`
	ChannelCustomInfo     string                 `json:"channel_custom_info,omitempty"`
	BrandId               string                 `json:"brand_id,omitempty"`
	CouponTagInfo         *CouponTagInfo         `json:"coupon_tag_info,omitempty"`
}

type CouponTagInfo struct {
	CouponTagList []string       `json:"coupon_tag_list,omitempty"`
	MemberTagInfo *MemberTagInfo `json:"member_tag_info,omitempty"`
}

const (
	USERPRODUCTCOUPONSTATE_CONFIRMING  string = "CONFIRMING"
	USERPRODUCTCOUPONSTATE_PENDING     string = "PENDING"
	USERPRODUCTCOUPONSTATE_EFFECTIVE   string = "EFFECTIVE"
	USERPRODUCTCOUPONSTATE_USED        string = "USED"
	USERPRODUCTCOUPONSTATE_EXPIRED     string = "EXPIRED"
	USERPRODUCTCOUPONSTATE_DELETED     string = "DELETED"
	USERPRODUCTCOUPONSTATE_DEACTIVATED string = "DEACTIVATED"
)

const (
	USERPRODUCTCOUPONSENDCHANNEL_API               string = "API"
	USERPRODUCTCOUPONSENDCHANNEL_BRAND_MANAGE      string = "BRAND_MANAGE"
	USERPRODUCTCOUPONSENDCHANNEL_MERCHANT_CARD     string = "MERCHANT_CARD"
	USERPRODUCTCOUPONSENDCHANNEL_MEMBER            string = "MEMBER"
	USERPRODUCTCOUPONSENDCHANNEL_SMALL_ACTIVITY    string = "SMALL_ACTIVITY"
	USERPRODUCTCOUPONSENDCHANNEL_RECEIVE_COMPONENT string = "RECEIVE_COMPONENT"
)

type SingleUsageDetail struct {
	UseRequestNo        string                                `json:"use_request_no,omitempty"`
	UseTime             time.Time                             `json:"use_time,omitempty"`
	AssociatedOrderInfo *UserProductCouponAssociatedOrderInfo `json:"associated_order_info,omitempty"`
	ReturnRequestNo     string                                `json:"return_request_no,omitempty"`
	ReturnTime          time.Time                             `json:"return_time,omitempty"`
}

type SequentialUsageDetail struct {
	TotalCount     int64                       `json:"total_count,omitempty"`
	UsedCount      int64                       `json:"used_count,omitempty"`
	DetailItemList []SequentialUsageDetailItem `json:"detail_item_list,omitempty"`
}

const (
	USERPRODUCTCOUPONTAG_MEMBER string = "MEMBER"
)

type MemberTagInfo struct {
	MemberCardId string `json:"member_card_id,omitempty"`
}

type UserProductCouponAssociatedOrderInfo struct {
	TransactionId string `json:"transaction_id,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	Mchid         string `json:"mchid,omitempty"`
	SubMchid      string `json:"sub_mchid,omitempty"`
}

type SequentialUsageDetailItem struct {
	DetailState         string                                `json:"detail_state,omitempty"`
	ValidBeginTime      time.Time                             `json:"valid_begin_time,omitempty"`
	ValidEndTime        time.Time                             `json:"valid_end_time,omitempty"`
	UseRequestNo        string                                `json:"use_request_no,omitempty"`
	UseTime             time.Time                             `json:"use_time,omitempty"`
	AssociatedOrderInfo *UserProductCouponAssociatedOrderInfo `json:"associated_order_info,omitempty"`
	ReturnRequestNo     string                                `json:"return_request_no,omitempty"`
	ReturnTime          time.Time                             `json:"return_time,omitempty"`
	DeleteTime          time.Time                             `json:"delete_time,omitempty"`
}

const (
	USERPRODUCTCOUPONUSAGEDETAILITEMSTATE_PENDING     string = "PENDING"
	USERPRODUCTCOUPONUSAGEDETAILITEMSTATE_EFFECTIVE   string = "EFFECTIVE"
	USERPRODUCTCOUPONUSAGEDETAILITEMSTATE_USED        string = "USED"
	USERPRODUCTCOUPONUSAGEDETAILITEMSTATE_EXPIRED     string = "EXPIRED"
	USERPRODUCTCOUPONUSAGEDETAILITEMSTATE_DELETED     string = "DELETED"
	USERPRODUCTCOUPONUSAGEDETAILITEMSTATE_DEACTIVATED string = "DEACTIVATED"
)

type UserProductCouponRequest struct {
	ProductCouponId string `json:"product_coupon_id,omitempty"`
	StockId         string `json:"stock_id,omitempty"`
	Appid           string `json:"appid,omitempty"`
	OutRequestNo    string `json:"out_request_no,omitempty"`
	BrandId         string `json:"brand_id,omitempty"`
}

type UseUserProductCouponRequest struct {
	ProductCouponId       string                                `json:"product_coupon_id,omitempty"`
	StockId               string                                `json:"stock_id,omitempty"`
	Appid                 string                                `json:"appid,omitempty"`
	UseTime               time.Time                             `json:"use_time,omitempty"`
	AssociatedOrderInfo   *UserProductCouponAssociatedOrderInfo `json:"associated_order_info,omitempty"`
	OutRequestNo          string                                `json:"out_request_no,omitempty"`
	SequentialCouponIndex int64                                 `json:"sequential_coupon_index,omitempty"`
	BrandId               string                                `json:"brand_id,omitempty"`
}

type DeactivateUserProductCouponRequest struct {
	ProductCouponId  string `json:"product_coupon_id,omitempty"`
	StockId          string `json:"stock_id,omitempty"`
	Appid            string `json:"appid,omitempty"`
	OutRequestNo     string `json:"out_request_no,omitempty"`
	DeactivateReason string `json:"deactivate_reason,omitempty"`
	BrandId          string `json:"brand_id,omitempty"`
}

type ReturnUserProductCouponRequest struct {
	ProductCouponId       string `json:"product_coupon_id,omitempty"`
	StockId               string `json:"stock_id,omitempty"`
	Appid                 string `json:"appid,omitempty"`
	OutRequestNo          string `json:"out_request_no,omitempty"`
	SequentialCouponIndex int64  `json:"sequential_coupon_index,omitempty"`
	BrandId               string `json:"brand_id,omitempty"`
}
