package productcoupon

import (
	"time"
)

type CreateStockRequest struct {
	OutRequestNo string          `json:"out_request_no,omitempty"`
	Stock        *StockForCreate `json:"stock,omitempty"`
}

type StockEntity struct {
	ProductCouponId      string                `json:"product_coupon_id,omitempty"`
	StockId              string                `json:"stock_id,omitempty"`
	Remark               string                `json:"remark,omitempty"`
	CouponCodeMode       string                `json:"coupon_code_mode,omitempty"`
	CouponCodeCountInfo  *CouponCodeCountInfo  `json:"coupon_code_count_info,omitempty"`
	StockSendRule        *StockSendRule        `json:"stock_send_rule,omitempty"`
	SingleUsageRule      *SingleUsageRule      `json:"single_usage_rule,omitempty"`
	SequentialUsageRule  *SequentialUsageRule  `json:"sequential_usage_rule,omitempty"`
	UsageRuleDisplayInfo *UsageRuleDisplayInfo `json:"usage_rule_display_info,omitempty"`
	CouponDisplayInfo    *CouponDisplayInfo    `json:"coupon_display_info,omitempty"`
	NotifyConfig         *NotifyConfig         `json:"notify_config,omitempty"`
	StoreScope           string                `json:"store_scope,omitempty"`
	SentCountInfo        *StockSentCountInfo   `json:"sent_count_info,omitempty"`
	State                *StockState           `json:"state,omitempty"`
	DeactivateRequestNo  string                `json:"deactivate_request_no,omitempty"`
	DeactivateTime       *time.Time            `json:"deactivate_time,omitempty"`
	DeactivateReason     string                `json:"deactivate_reason,omitempty"`
}

type StockForCreate struct {
	Remark               string                `json:"remark,omitempty"`
	CouponCodeMode       string                `json:"coupon_code_mode,omitempty"`
	StockSendRule        *StockSendRule        `json:"stock_send_rule,omitempty"`
	SingleUsageRule      *SingleUsageRule      `json:"single_usage_rule,omitempty"`
	SequentialUsageRule  *SequentialUsageRule  `json:"sequential_usage_rule,omitempty"`
	UsageRuleDisplayInfo *UsageRuleDisplayInfo `json:"usage_rule_display_info,omitempty"`
	CouponDisplayInfo    *CouponDisplayInfo    `json:"coupon_display_info,omitempty"`
	NotifyConfig         *NotifyConfig         `json:"notify_config,omitempty"`
	StoreScope           string                `json:"store_scope,omitempty"`
}

const (
	COUPONCODEMODE_WECHATPAY  string = "WECHATPAY"
	COUPONCODEMODE_UPLOAD     string = "UPLOAD"
	COUPONCODEMODE_API_ASSIGN string = "API_ASSIGN"
)

type CouponCodeCountInfo struct {
	TotalCount     int64 `json:"total_count,omitempty"`
	AvailableCount int64 `json:"available_count,omitempty"`
}

type StockSendRule struct {
	MaxCount        int64 `json:"max_count,omitempty"`
	MaxCountPerDay  int64 `json:"max_count_per_day,omitempty"`
	MaxCountPerUser int64 `json:"max_count_per_user,omitempty"`
}

type SingleUsageRule struct {
	CouponAvailablePeriod *SingleCouponAvailablePeriod `json:"coupon_available_period,omitempty"`
	NormalCoupon          *NormalCouponUsageRule       `json:"normal_coupon,omitempty"`
	DiscountCoupon        *DiscountCouponUsageRule     `json:"discount_coupon,omitempty"`
	ExchangeCoupon        *ExchangeCouponUsageRule     `json:"exchange_coupon,omitempty"`
}

type SequentialUsageRule struct {
	CouponAvailablePeriod *SequentialCouponAvailablePeriod `json:"coupon_available_period,omitempty"`
	NormalCouponList      []NormalCouponUsageRule          `json:"normal_coupon_list,omitempty"`
	DiscountCouponList    []DiscountCouponUsageRule        `json:"discount_coupon_list,omitempty"`
	ExchangeCouponList    []ExchangeCouponUsageRule        `json:"exchange_coupon_list,omitempty"`
	SpecialFirst          bool                             `json:"special_first,omitempty"`
}

type UsageRuleDisplayInfo struct {
	CouponUsageMethodList    []string                  `json:"coupon_usage_method_list,omitempty"`
	MiniProgramAppid         string                    `json:"mini_program_appid,omitempty"`
	MiniProgramPath          string                    `json:"mini_program_path,omitempty"`
	AppPath                  string                    `json:"app_path,omitempty"`
	UsageDescription         string                    `json:"usage_description,omitempty"`
	CouponAvailableStoreInfo *CouponAvailableStoreInfo `json:"coupon_available_store_info,omitempty"`
}

type CouponDisplayInfo struct {
	CodeDisplayMode         string                   `json:"code_display_mode,omitempty"`
	BackgroundColor         string                   `json:"background_color,omitempty"`
	EntranceMiniProgram     *EntranceMiniProgram     `json:"entrance_mini_program,omitempty"`
	EntranceOfficialAccount *EntranceOfficialAccount `json:"entrance_official_account,omitempty"`
	EntranceFinder          *EntranceFinder          `json:"entrance_finder,omitempty"`
}

type NotifyConfig struct {
	NotifyAppid string `json:"notify_appid,omitempty"`
}

const (
	STOCKSTORESCOPE_NONE     string = "NONE"
	STOCKSTORESCOPE_ALL      string = "ALL"
	STOCKSTORESCOPE_SPECIFIC string = "SPECIFIC"
)

type StockSentCountInfo struct {
	TotalCount int64 `json:"total_count,omitempty"`
	TodayCount int64 `json:"today_count,omitempty"`
}

type StockState string

func (e StockState) Ptr() *StockState {
	return &e
}

const (
	STOCKSTATE_AUDITING    StockState = "AUDITING"
	STOCKSTATE_SENDING     StockState = "SENDING"
	STOCKSTATE_PAUSED      StockState = "PAUSED"
	STOCKSTATE_STOPPED     StockState = "STOPPED"
	STOCKSTATE_DEACTIVATED StockState = "DEACTIVATED"
)

type SingleCouponAvailablePeriod struct {
	AvailableBeginTime           string           `json:"available_begin_time,omitempty"`
	AvailableEndTime             string           `json:"available_end_time,omitempty"`
	AvailableDays                int64            `json:"available_days,omitempty"`
	WaitDaysAfterReceive         int64            `json:"wait_days_after_receive,omitempty"`
	WeeklyAvailablePeriod        *FixedWeekPeriod `json:"weekly_available_period,omitempty"`
	IrregularAvailablePeriodList []TimePeriod     `json:"irregular_available_period_list,omitempty"`
}

type NormalCouponUsageRule struct {
	Threshold      int64 `json:"threshold,omitempty"`
	DiscountAmount int64 `json:"discount_amount,omitempty"`
}

type DiscountCouponUsageRule struct {
	Threshold  int64 `json:"threshold,omitempty"`
	PercentOff int64 `json:"percent_off,omitempty"`
}

type ExchangeCouponUsageRule struct {
	Threshold     int64 `json:"threshold"`
	ExchangePrice int64 `json:"exchange_price"`
}

type SequentialCouponAvailablePeriod struct {
	AvailableBeginTime           string           `json:"available_begin_time,omitempty"`
	AvailableEndTime             string           `json:"available_end_time,omitempty"`
	WaitDaysAfterReceive         int64            `json:"wait_days_after_receive,omitempty"`
	WeeklyAvailablePeriod        *FixedWeekPeriod `json:"weekly_available_period,omitempty"`
	IrregularAvailablePeriodList []TimePeriod     `json:"irregular_available_period_list,omitempty"`
}

const (
	COUPONUSAGEMETHOD_OFFLINE      string = "OFFLINE"
	COUPONUSAGEMETHOD_MINI_PROGRAM string = "MINI_PROGRAM"
	COUPONUSAGEMETHOD_APP          string = "APP"
	COUPONUSAGEMETHOD_PAYMENT_CODE string = "PAYMENT_CODE"
)

type CouponAvailableStoreInfo struct {
	Description      string `json:"description,omitempty"`
	MiniProgramAppid string `json:"mini_program_appid,omitempty"`
	MiniProgramPath  string `json:"mini_program_path,omitempty"`
}

const (
	COUPONCODEDISPLAYMODE_INVISIBLE string = "INVISIBLE"
	COUPONCODEDISPLAYMODE_BARCODE   string = "BARCODE"
	COUPONCODEDISPLAYMODE_QRCODE    string = "QRCODE"
)

type EntranceMiniProgram struct {
	Appid           string `json:"appid,omitempty"`
	Path            string `json:"path,omitempty"`
	EntranceWording string `json:"entrance_wording,omitempty"`
	GuidanceWording string `json:"guidance_wording,omitempty"`
}

type EntranceOfficialAccount struct {
	Appid string `json:"appid,omitempty"`
}

type EntranceFinder struct {
	FinderId                 string `json:"finder_id,omitempty"`
	FinderVideoId            string `json:"finder_video_id,omitempty"`
	FinderVideoCoverImageUrl string `json:"finder_video_cover_image_url,omitempty"`
}

type FixedWeekPeriod struct {
	DayList       []string         `json:"day_list,omitempty"`
	DayPeriodList []PeriodOfTheDay `json:"day_period_list,omitempty"`
}

type TimePeriod struct {
	BeginTime string `json:"begin_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

const (
	WEEKENUM_MONDAY    string = "MONDAY"
	WEEKENUM_TUESDAY   string = "TUESDAY"
	WEEKENUM_WEDNESDAY string = "WEDNESDAY"
	WEEKENUM_THURSDAY  string = "THURSDAY"
	WEEKENUM_FRIDAY    string = "FRIDAY"
	WEEKENUM_SATURDAY  string = "SATURDAY"
	WEEKENUM_SUNDAY    string = "SUNDAY"
)

type PeriodOfTheDay struct {
	BeginTime int64 `json:"begin_time,omitempty"`
	EndTime   int64 `json:"end_time,omitempty"`
}

type CreateProductCouponRequest struct {
	OutRequestNo        string                    `json:"out_request_no,omitempty"`
	Scope               string                    `json:"scope,omitempty"`
	Type                string                    `json:"type,omitempty"`
	UsageMode           string                    `json:"usage_mode,omitempty"`
	SingleUsageInfo     *SingleUsageInfo          `json:"single_usage_info,omitempty"`
	SequentialUsageInfo *SequentialUsageInfo      `json:"sequential_usage_info,omitempty"`
	DisplayInfo         *ProductCouponDisplayInfo `json:"display_info,omitempty"`
	OutProductNo        string                    `json:"out_product_no,omitempty"`
	Stock               *StockForCreate           `json:"stock,omitempty"`
	BrandId             string                    `json:"brand_id"`
}

type CreateProductCouponResponse struct {
	ProductCouponId     string                    `json:"product_coupon_id,omitempty"`
	Scope               string                    `json:"scope,omitempty"`
	Type                string                    `json:"type,omitempty"`
	UsageMode           string                    `json:"usage_mode,omitempty"`
	SingleUsageInfo     *SingleUsageInfo          `json:"single_usage_info,omitempty"`
	SequentialUsageInfo *SequentialUsageInfo      `json:"sequential_usage_info,omitempty"`
	DisplayInfo         *ProductCouponDisplayInfo `json:"display_info,omitempty"`
	OutProductNo        string                    `json:"out_product_no,omitempty"`
	State               string                    `json:"state,omitempty"`
	Stock               *StockEntity              `json:"stock,omitempty"`
}

type ProductCouponEntity struct {
	ProductCouponId     string                    `json:"product_coupon_id,omitempty"`
	Scope               string                    `json:"scope,omitempty"`
	Type                string                    `json:"type,omitempty"`
	UsageMode           string                    `json:"usage_mode,omitempty"`
	SingleUsageInfo     *SingleUsageInfo          `json:"single_usage_info,omitempty"`
	SequentialUsageInfo *SequentialUsageInfo      `json:"sequential_usage_info,omitempty"`
	DisplayInfo         *ProductCouponDisplayInfo `json:"display_info,omitempty"`
	OutProductNo        string                    `json:"out_product_no,omitempty"`
	State               string                    `json:"state,omitempty"`
	DeactivateRequestNo string                    `json:"deactivate_request_no,omitempty"`
	DeactivateTime      string                    `json:"deactivate_time,omitempty"`
	DeactivateReason    string                    `json:"deactivate_reason,omitempty"`
}

const (
	PRODUCTCOUPONSCOPE_ALL    string = "ALL"
	PRODUCTCOUPONSCOPE_SINGLE string = "SINGLE"
)

const (
	PRODUCTCOUPONTYPE_NORMAL   string = "NORMAL"
	PRODUCTCOUPONTYPE_DISCOUNT string = "DISCOUNT"
	PRODUCTCOUPONTYPE_EXCHANGE string = "EXCHANGE"
)

const (
	USAGEMODE_SINGLE     string = "SINGLE"
	USAGEMODE_SEQUENTIAL string = "SEQUENTIAL"
)

type SingleUsageInfo struct {
	NormalCoupon   *NormalCouponUsageRule   `json:"normal_coupon,omitempty"`
	DiscountCoupon *DiscountCouponUsageRule `json:"discount_coupon,omitempty"`
}

type SequentialUsageInfo struct {
	Type          string `json:"type,omitempty"`
	Count         int64  `json:"count,omitempty"`
	AvailableDays int64  `json:"available_days,omitempty"`
	IntervalDays  int64  `json:"interval_days,omitempty"`
}

type ProductCouponDisplayInfo struct {
	Name               string         `json:"name,omitempty"`
	ImageUrl           string         `json:"image_url,omitempty"`
	BackgroundUrl      string         `json:"background_url,omitempty"`
	DetailImageUrlList []string       `json:"detail_image_url_list,omitempty"`
	OriginalPrice      int64          `json:"original_price,omitempty"`
	ComboPackageList   []ComboPackage `json:"combo_package_list,omitempty"`
}

const (
	PRODUCTCOUPONSTATE_AUDITING    string = "AUDITING"
	PRODUCTCOUPONSTATE_EFFECTIVE   string = "EFFECTIVE"
	PRODUCTCOUPONSTATE_DEACTIVATED string = "DEACTIVATED"
)

const (
	SEQUENTIALUSAGETYPE_INCREMENTAL string = "INCREMENTAL"
	SEQUENTIALUSAGETYPE_EQUAL       string = "EQUAL"
)

type ComboPackage struct {
	Name       string               `json:"name,omitempty"`
	PickCount  int64                `json:"pick_count,omitempty"`
	ChoiceList []ComboPackageChoice `json:"choice_list,omitempty"`
}

type ComboPackageChoice struct {
	Name             string `json:"name,omitempty"`
	Price            int64  `json:"price,omitempty"`
	Count            int64  `json:"count,omitempty"`
	ImageUrl         string `json:"image_url,omitempty"`
	MiniProgramAppid string `json:"mini_program_appid,omitempty"`
	MiniProgramPath  string `json:"mini_program_path,omitempty"`
}

type SetNotifyConfigRequest struct {
	NotifyUrl string `json:"notify_url,omitempty"`
}

type SetNotifyConfigResponse struct {
	NotifyUrl  string    `json:"notify_url,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}

type UpdateProductCouponRequest struct {
	OutRequestNo string                    `json:"out_request_no,omitempty"`
	DisplayInfo  *ProductCouponDisplayInfo `json:"display_info,omitempty"`
}

type DeactivateRequest struct {
	OutRequestNo     string `json:"out_request_no,omitempty"`
	DeactivateReason string `json:"deactivate_reason"`
	BrandId          string `json:"brand_id"`
}

type ListStocksResponse struct {
	TotalCount    int64          `json:"total_count,omitempty"`
	StockList     []*StockEntity `json:"stock_list,omitempty"`
	NextPageToken string         `json:"next_page_token,omitempty"`
}

type UpdateStockRequest struct {
	OutRequestNo         string                `json:"out_request_no,omitempty"`
	Remark               string                `json:"remark,omitempty"`
	UsageRuleDisplayInfo *UsageRuleDisplayInfo `json:"usage_rule_display_info,omitempty"`
	CouponDisplayInfo    *CouponDisplayInfo    `json:"coupon_display_info,omitempty"`
	NotifyConfig         *NotifyConfig         `json:"notify_config,omitempty"`
	StoreScope           string                `json:"store_scope,omitempty"`
	BrandId              string                `json:"brand_id,omitempty"`
}

type UpdateStockBudgetRequest struct {
	OutRequestNo          string `json:"out_request_no,omitempty"`
	UpdateMode            string `json:"update_mode,omitempty"`
	CurrentMaxCount       *int64 `json:"current_max_count,omitempty"`
	TargetMaxCount        *int64 `json:"target_max_count,omitempty"`
	CurrentMaxCountPerDay *int64 `json:"current_max_count_per_day,omitempty"`
	TargetMaxCountPerDay  *int64 `json:"target_max_count_per_day,omitempty"`
	BrandId               string `json:"brand_id,omitempty"`
}
