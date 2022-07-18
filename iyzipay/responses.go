package iyzipay

type (
	Resource struct {
		Status         string `json:"status"`
		ErrorCode      string `json:"errorCode"`
		ErrorMessage   string `json:"errorMessage"`
		ErrorGroup     string `json:"errorGroup"`
		Locale         string `json:"locale"`
		SystemTime     int64  `json:"systemTime"`
		ConversationId string `json:"conversationId"`
	}

	SubMerchantResponse struct {
		Resource
		Name                  string `json:"name"`
		Email                 string `json:"email"`
		GsmNumber             string `json:"gsmNumber"`
		Address               string `json:"address"`
		Iban                  string `json:"iban"`
		SwiftCode             string `json:"swiftCode"`
		Currency              string `json:"currency"`
		TaxOffice             string `json:"taxOffice"`
		ContactName           string `json:"contactName"`
		ContactSurname        string `json:"contactSurname"`
		LegalCompanyTitle     string `json:"legalCompanyTitle"`
		SubMerchantExternalId string `json:"subMerchantExternalId"`
		IdentityNumber        string `json:"identityNumber"`
		TaxNumber             string `json:"taxNumber"`
		SubMerchantType       string `json:"subMerchantType"`
		SubMerchantKey        string `json:"subMerchantKey"`
	}
)
