package midtranspay

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func InitConnection() coreapi.Client {
	var c = coreapi.Client{}
	c.New("SB-Mid-server-BiEyXQZgG7tFbTYC-fgBCWFI", midtrans.Sandbox)
	return c
}

func CreateTransaction(core coreapi.Client) *coreapi.ChargeResponse {
	req := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeConvenienceStore,
		ConvStore: &coreapi.ConvStoreDetails{
			Store: "Indomaret",
		},
		BCAKlikPay: &coreapi.BCAKlikPayDetails{},
		// coreapi.PaymentTypeBCAKlikpay,
		// BCAKlikPay: &coreapi.BCAKlikPayDetails{
		// 	Desc: "Coba Pembayaran",
		// },
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "be6-004",
			GrossAmt: 150000,
		},
		Items: &[]midtrans.ItemDetails{
			{Name: "Item 1", Price: 10000, Qty: 5},
			{Name: "Item 2", Price: 20000, Qty: 5},
		},
	}

	apiRes, err := core.ChargeTransaction(req)

	if err != nil {
		log.Warn("Payment err :", err)
	}
	fmt.Println(apiRes)
	return apiRes
}
