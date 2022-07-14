package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type User struct {
    OrderNo             int           	`db:"order_no" json:"order_no"`
    InvoiceNo           string          `db:"invoice_no" json:"invoice_no"`
    BillingCode         null.String     `db:"billing_code" json:"billing_code"`
    ClientID            int             `db:"client_id" json:"client_id"`
    ClientIdentityNo    string          `db:"client_identity_no" json:"client_identity_no"`
}

func main() {
    var jsonString = `{
        "id": 76,
        "client_cif_code": "20022",
        "product_code": "IDN000199809",
        "order_no": 18010900001,
        "type": "redemption",
        "amount": 2048355.00,
        "fee": null,
        "unit": 45519.00,
        "all_unit": 1,
        "order_date": "2018-01-09T03:13:49",
        "effective_date": "2018-01-09",
        "currency": "IDR",
        "status": 3,
        "confirmation_letter": null,
        "promo_code": null,
        "message": null,
        "created_at": "2018-01-09T03:13:49",
        "updated_at": "2018-01-09T03:14:01",
        "tokopedia_order_id": "20367",
        "current_unit_value": 48770.3571,
        "real_amount": null,
        "batch_no": null
	}`

    var jsonData = []byte(jsonString)
    var data bytes.Buffer
    var err = json.Unmarshal(jsonData, &data)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
	json.Indent(&data, jsonData, "", "")
	var mapping User;
	json.Unmarshal(data.Bytes(), &mapping)
    fmt.Println("order No :", mapping.OrderNo)
}