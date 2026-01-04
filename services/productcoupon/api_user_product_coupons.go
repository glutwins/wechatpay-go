package productcoupon

import (
	"context"
	"fmt"
	"net/url"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
)

func (a *ProductCouponService) SendProductCoupon(ctx context.Context, openId string, req *SendUserProductCouponRequest) (resp *UserProductCouponEntity, result *core.APIResult, err error) {
	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/partner/product-coupon/users/%s/coupons"

	result, err = a.Client.Post(ctx, fmt.Sprintf(localVarPath, openId), req)
	if err != nil {
		return nil, result, err
	}

	resp = new(UserProductCouponEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) ConfirmSendProductCoupon(ctx context.Context, openId string, couponCode string, req *UserProductCouponRequest) (resp *UserProductCouponEntity, result *core.APIResult, err error) {
	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/partner/product-coupon/users/%s/coupons/%s/confirm"

	result, err = a.Client.Post(ctx, fmt.Sprintf(localVarPath, openId, couponCode), req)
	if err != nil {
		return nil, result, err
	}

	resp = new(UserProductCouponEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) UserProductCoupon(ctx context.Context, openId string, couponCode string, req *UserProductCouponRequest) (resp *UserProductCouponEntity, result *core.APIResult, err error) {
	vals := make(url.Values)
	vals.Add("product_coupon_id", req.ProductCouponId)
	vals.Add("stock_id", req.StockId)
	vals.Add("appid", req.Appid)
	vals.Add("brand_id", req.BrandId)
	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/partner/product-coupon/users/%s/coupons/%s?" + vals.Encode()

	result, err = a.Client.Get(ctx, fmt.Sprintf(localVarPath, openId, couponCode))
	if err != nil {
		return nil, result, err
	}

	resp = new(UserProductCouponEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) ListUserProductCoupon(ctx context.Context, openId string, state string, req *UserProductCouponRequest) (resp *UserProductCouponEntity, result *core.APIResult, err error) {
	vals := make(url.Values)
	if state != "" {
		vals.Add("coupon_state", state)
	}
	if req.ProductCouponId != "" {
		vals.Add("product_coupon_id", req.ProductCouponId)
	}
	if req.StockId != "" {
		vals.Add("stock_id", req.StockId)
	}
	vals.Add("appid", req.Appid)
	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/partner/product-coupon/users/%s/coupons?" + vals.Encode()

	result, err = a.Client.Get(ctx, fmt.Sprintf(localVarPath, openId))
	if err != nil {
		return nil, result, err
	}

	resp = new(UserProductCouponEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) UseUserProductCoupon(ctx context.Context, openId string, couponCode string, req *UseUserProductCouponRequest) (resp *UserProductCouponEntity, result *core.APIResult, err error) {
	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/partner/product-coupon/users/%s/coupons/%s/use"

	result, err = a.Client.Post(ctx, fmt.Sprintf(localVarPath, openId, couponCode), req)
	if err != nil {
		return nil, result, err
	}

	resp = new(UserProductCouponEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) DeactivateUserProductCoupon(ctx context.Context, openId string, couponCode string, req *DeactivateUserProductCouponRequest) (resp *UserProductCouponEntity, result *core.APIResult, err error) {
	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/partner/product-coupon/users/%s/coupons/%s/deactivate"

	result, err = a.Client.Post(ctx, fmt.Sprintf(localVarPath, openId, couponCode), req)
	if err != nil {
		return nil, result, err
	}

	resp = new(UserProductCouponEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ProductCouponService) ReturnUserProductCoupon(ctx context.Context, openId string, couponCode string, req *ReturnUserProductCouponRequest) (resp *UserProductCouponEntity, result *core.APIResult, err error) {
	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/partner/product-coupon/users/%s/coupons/%s/return"

	result, err = a.Client.Post(ctx, fmt.Sprintf(localVarPath, openId, couponCode), req)
	if err != nil {
		return nil, result, err
	}

	resp = new(UserProductCouponEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
