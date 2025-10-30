package productcoupon

import (
	"context"
	"fmt"
	"net/url"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

type ProductCouponService services.Service

func (a *ProductCouponService) CreateProductCoupon(ctx context.Context, req *CreateProductCouponRequest) (resp *CreateProductCouponResponse, result *core.APIResult, err error) {
	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/partner/product-coupon/product-coupons"

	result, err = a.Client.Post(ctx, localVarPath, req)
	if err != nil {
		return nil, result, err
	}

	resp = new(CreateProductCouponResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) UpdateProductCoupon(ctx context.Context, productCouponId string, req *UpdateProductCouponRequest) (resp *ProductCouponEntity, result *core.APIResult, err error) {
	result, err = a.Client.Patch(ctx, fmt.Sprintf(consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/product-coupons/%s", productCouponId), req)
	if err != nil {
		return nil, result, err
	}

	resp = new(ProductCouponEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) GetProductCoupon(ctx context.Context, productCouponId string, brandId string) (resp *ProductCouponEntity, result *core.APIResult, err error) {
	result, err = a.Client.Get(ctx, fmt.Sprintf(consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/product-coupons/%s?brand_id=%s", productCouponId, brandId))
	if err != nil {
		return nil, result, err
	}

	resp = new(ProductCouponEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) DeactivateProductCoupon(ctx context.Context, productCouponId string, req *DeactivateRequest) (resp *ProductCouponEntity, result *core.APIResult, err error) {
	result, err = a.Client.Post(ctx, fmt.Sprintf(consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/product-coupons/%s/deactivate", productCouponId), req)
	if err != nil {
		return nil, result, err
	}

	resp = new(ProductCouponEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) CreateStock(ctx context.Context, productCouponId string, req CreateStockRequest) (resp *StockEntity, result *core.APIResult, err error) {
	localVarPath := fmt.Sprintf(consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/product-coupons/%s/stocks", productCouponId)

	result, err = a.Client.Post(ctx, localVarPath, req)
	if err != nil {
		return nil, result, err
	}

	resp = new(StockEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) ListStocks(ctx context.Context, productCouponId string, state string, pageSize int, pageToken string, brandId string) (resp *ListStocksResponse, result *core.APIResult, err error) {
	vals := make(url.Values)
	if state != "" {
		vals.Add("state", state)
	}
	if pageSize != 0 {
		vals.Add("page_size", fmt.Sprintf("%d", pageSize))
	}

	if pageToken != "" {
		vals.Add("page_token", pageToken)
	}

	vals.Add("brand_id", brandId)
	localVarPath := fmt.Sprintf(consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/product-coupons/%s/stocks?"+vals.Encode(), productCouponId)

	result, err = a.Client.Get(ctx, localVarPath)
	if err != nil {
		return nil, result, err
	}

	resp = new(ListStocksResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) GetStock(ctx context.Context, productCouponId string, stockId string, brandId string) (resp *StockEntity, result *core.APIResult, err error) {
	localVarPath := fmt.Sprintf(consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/product-coupons/%s/stocks/%s?brand_id=%s", productCouponId, stockId, brandId)

	result, err = a.Client.Get(ctx, localVarPath)
	if err != nil {
		return nil, result, err
	}

	resp = new(StockEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) UpdateStock(ctx context.Context, productCouponId string, stockId string, req *UpdateStockRequest) (resp *StockEntity, result *core.APIResult, err error) {
	localVarPath := fmt.Sprintf(consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/product-coupons/%s/stocks/%s", productCouponId, stockId)

	result, err = a.Client.Patch(ctx, localVarPath, req)
	if err != nil {
		return nil, result, err
	}

	resp = new(StockEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) UpdateStockBudget(ctx context.Context, productCouponId string, stockId string, req *UpdateStockBudgetRequest) (resp *StockEntity, result *core.APIResult, err error) {
	localVarPath := fmt.Sprintf(consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/product-coupons/%s/stocks/%s/update-budget", productCouponId, stockId)

	result, err = a.Client.Post(ctx, localVarPath, req)
	if err != nil {
		return nil, result, err
	}

	resp = new(StockEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) DeactivateStock(ctx context.Context, productCouponId string, stockId string, req *DeactivateRequest) (resp *StockEntity, result *core.APIResult, err error) {
	localVarPath := fmt.Sprintf(consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/product-coupons/%s/stocks/%s/deactivate", productCouponId, stockId)

	result, err = a.Client.Post(ctx, localVarPath, req)
	if err != nil {
		return nil, result, err
	}

	resp = new(StockEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) GetNotifyConfig(ctx context.Context) (resp *SetNotifyConfigResponse, result *core.APIResult, err error) {
	result, err = a.Client.Get(ctx, consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/notify-configs")
	if err != nil {
		return nil, result, err
	}

	resp = new(SetNotifyConfigResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) SetNotifyConfig(ctx context.Context, notifyURL string) (resp *SetNotifyConfigResponse, result *core.APIResult, err error) {
	result, err = a.Client.Post(ctx, consts.WechatPayAPIServer+"/v3/marketing/partner/product-coupon/notify-configs", &SetNotifyConfigRequest{NotifyUrl: notifyURL})
	if err != nil {
		return nil, result, err
	}

	resp = new(SetNotifyConfigResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
