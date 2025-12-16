package productcoupon

import (
	"bytes"
	"context"
	"crypto/sha256" // 引用微信支付工具库，参考 https://pay.weixin.qq.com/doc/v3/partner/4015119446
	"fmt"
	"mime"
	"mime/multipart"
	"path/filepath"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
)

func (a *ProductCouponService) UploadImage(filename string, content []byte) (response *UploadImageResponse, err error) {
	const path = consts.WechatPayAPIServer + "/v3/marketing/partner/product-coupon/media/upload-image"

	contentType := mime.TypeByExtension(filepath.Ext(filename))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	meta := make(map[string]interface{})
	meta["filename"] = core.String(filename)
	meta["sha256"] = core.String(fmt.Sprintf("%x", sha256.Sum256(content)))

	metaStr, err := core.ParameterToJSON(meta)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err = core.CreateFormField(writer, "meta", "application/json", []byte(metaStr)); err != nil {
		return nil, err
	}

	if err = core.CreateFormFile(writer, filename, contentType, content); err != nil {
		return nil, err
	}

	if err = writer.Close(); err != nil {
		return nil, err
	}

	result, err := a.Client.Upload(context.Background(), path, metaStr, body.String(), writer.FormDataContentType())
	if err != nil {
		return nil, err
	}
	var resp = new(UploadImageResponse)
	if err = core.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, err
	}
	return resp, err
}

type UploadImageResponse struct {
	ImageUrl string `json:"image_url,omitempty"`
}
