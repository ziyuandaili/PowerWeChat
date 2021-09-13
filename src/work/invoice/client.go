package invoice

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/invoice/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90284
func (comp *Client) Get(cardID string, encryptCode string) (*response.ResponseInvoiceGetInfo, error) {

	result := &response.ResponseInvoiceGetInfo{}

	data := &object.HashMap{
		"card_id":      cardID,
		"encrypt_code": encryptCode,
	}
	_, err := comp.HttpPostJson("cgi-bin/card/invoice/reimburse/getinvoiceinfo", data, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90287
func (comp *Client) Select(invoices *power.HashMap) (*response.ResponseInvoiceGetInfoBatch, error) {

	result := &response.ResponseInvoiceGetInfoBatch{}

	data := &object.HashMap{
		"item_list": invoices.ToHashMap(),
	}
	_, err := comp.HttpPostJson("cgi-bin/card/invoice/reimburse/getinvoiceinfobatch", data, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90285
func (comp *Client) Update(cardId string, encryptCode string, status string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	data := &object.HashMap{
		"card_id":          cardId,
		"encrypt_code":     encryptCode,
		"reimburse_status": status,
	}
	_, err := comp.HttpPostJson("cgi-bin/card/invoice/reimburse/updateinvoicestatus", data, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90286
func (comp *Client) BatchUpdate(invoices *power.HashMap, openid string, status string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	data := &object.HashMap{
		"openid":           invoices.ToHashMap(),
		"reimburse_status": openid,
		"invoice_list":     status,
	}
	_, err := comp.HttpPostJson("cgi-bin/card/invoice/reimburse/updatestatusbatch", data, nil, nil, result)

	return result, err
}