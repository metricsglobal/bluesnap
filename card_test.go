package bluesnap

import (
	"testing"

	"github.com/metricsglobal/bluesnap/card"
)

func TestAuthOnly(t *testing.T) {
	c := New(testURL, testCredentials)
	tests := []struct {
		name    string
		input   Serializer
		output  Deserializer
		err     error
		wantErr bool
	}{
		{
			name: "Auth Only Request: basic with fraud info",
			input: &card.Request{
				Amount:         "11",
				SoftDescriptor: "DescTest",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				Currency: "USD",
				TransactionFraudInfo: &card.TransactionFraudInfoRequest{
					ShopperIpAddress: "123.12.134.1",
					FraudSessionId:   "1234",
					Company:          "BBBBB",
					ShippingContactInfo: &card.ShippingContactInfo{
						Zip:       "12345",
						Country:   "US",
						FirstName: "YY",
						LastName:  "LL",
						City:      "Juneau",
						Address2:  "Address2",
						Address1:  "Address1",
						State:     "AL",
					},
				},
				CreditCard: &card.CreditCardRequest{
					ExpirationYear:  "2023",
					SecurityCode:    "837",
					ExpirationMonth: "02",
					CardNumber:      "4263982640269299",
				},
				CardTransactionType: "AUTH_ONLY",
			},
			output: &card.Response{
				Amount:    11.00,
				USDAmount: 11.00,
				ProcessingInfo: card.ProcessingInfo{
					AVSResponseCodeAddress: "U",
					ProcessingStatus:       "success",
					CVVResponseCode:        "MA",
					AVSResponseCodeName:    "U",
					AVSResponseCodeZip:     "U",
					AuthorizationCode:      "123456",
				},
				SoftDescriptor: "BLS*DescTest",
				CardHolderInfo: card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				Currency: "USD",
				FraudResultInfo: card.FraudResultInfo{
					DeviceDataCollector: "",
				},
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "9299",
					CardSubType:        "CREDIT",
					CardType:           "VISA",
					CardCategory:       "PLATINUM",
					BinCategory:        "CONSUMER",
					BinNumber:          "426398",
					CardRegulated:      "N",
					IssuingBank:        "ALLIED IRISH BANKS PLC",
					IssuingCountryCode: "ie",
					ExpirationMonth:    "02",
					ExpirationYear:     "2023",
				},
				CardTransactionType: "AUTH_ONLY",
				TransactionID:       "38488222",
			},
			wantErr: false,
		},
		{
			name: "with fraud info",
			input: &card.Request{
				Amount:         "11",
				SoftDescriptor: "DescTest",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				Currency: "USD",
				TransactionFraudInfo: &card.TransactionFraudInfoRequest{
					EnterpriseSiteID: "DEFAULT",
					ShopperIpAddress: "123.12.134.1",
					FraudSessionId:   "1234",
					EnterpriseUDFs: &card.EnterpriseUDFs{
						UDFs: []card.UDF{{
							Name:  "ENT_UDF1",
							Value: "aaa",
						}, {
							Name:  "ENT_UDF2",
							Value: "bbb",
						}},
					},
					Company: "BBBBB",
					ShippingContactInfo: &card.ShippingContactInfo{
						Zip:       "12345",
						Country:   "US",
						FirstName: "YY",
						LastName:  "LL",
						City:      "Juneau",
						Address2:  "Address2",
						Address1:  "Address1",
						State:     "AL",
					},
				},
				CreditCard: &card.CreditCardRequest{
					ExpirationYear:  "2023",
					SecurityCode:    "837",
					ExpirationMonth: "02",
					CardNumber:      "4263982640269299",
				},
				CardTransactionType: "AUTH_ONLY",
			},
			output: &card.Response{
				Amount:    11.00,
				USDAmount: 11.00,
				ProcessingInfo: card.ProcessingInfo{
					AVSResponseCodeAddress: "U",
					ProcessingStatus:       "success",
					CVVResponseCode:        "MA",
					AVSResponseCodeName:    "U",
					AVSResponseCodeZip:     "U",
					AuthorizationCode:      "123456",
				},
				SoftDescriptor: "BLS*DescTest",
				CardHolderInfo: card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				Currency: "USD",
				FraudResultInfo: card.FraudResultInfo{
					DeviceDataCollector: "",
				},
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "9299",
					CardSubType:        "CREDIT",
					CardType:           "VISA",
					CardCategory:       "PLATINUM",
					BinCategory:        "CONSUMER",
					BinNumber:          "426398",
					CardRegulated:      "N",
					IssuingBank:        "ALLIED IRISH BANKS PLC",
					IssuingCountryCode: "ie",
					ExpirationMonth:    "02",
					ExpirationYear:     "2023",
				},
				CardTransactionType: "AUTH_ONLY",
				TransactionID:       "38488222",
			},
			wantErr: false,
		},
		{
			name: "with vaulted shopper who has one card",
			input: &card.Request{
				Amount:              "11",
				VaultedShopperID:    20781033,
				SoftDescriptor:      "DescTest",
				Currency:            "USD",
				CardTransactionType: "AUTH_ONLY",
			},
			output: &card.Response{
				Amount:           11.00,
				USDAmount:        11.00,
				VaultedShopperID: 20781033,
				ProcessingInfo: card.ProcessingInfo{
					AVSResponseCodeAddress: "U",
					ProcessingStatus:       "success",
					CVVResponseCode:        "",
					AVSResponseCodeName:    "U",
					AVSResponseCodeZip:     "U",
					AuthorizationCode:      "654321",
				},
				SoftDescriptor: "BLS*DescTest",
				Currency:       "USD",
				FraudResultInfo: card.FraudResultInfo{
					DeviceDataCollector: "",
				},
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "0026",
					CardSubType:        "CREDIT",
					CardType:           "VISA",
					CardCategory:       "CLASSIC",
					BinCategory:        "CONSUMER",
					BinNumber:          "401200",
					CardRegulated:      "N",
					IssuingBank:        "",
					IssuingCountryCode: "ru",
					ExpirationMonth:    "7",
					ExpirationYear:     "2021",
				},
				CardTransactionType: "AUTH_ONLY",
				TransactionID:       "38488222",
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := card.Response{}
			if err := c.CardAuth(test.input, &o); (err != nil) != test.wantErr {
				t.Error(err)
			}
			if !CompareResponse(test.output.(*card.Response), &o) {
				t.Errorf("Expected output: \n%#v, got: \n%#v", test.output, &o)
			}
		})
	}
}

func CompareResponse(exp, got *card.Response) bool {
	return exp.Amount == got.Amount &&
		exp.ProcessingInfo.ProcessingStatus == got.ProcessingInfo.ProcessingStatus &&
		exp.ProcessingInfo.CVVResponseCode == got.ProcessingInfo.CVVResponseCode &&
		exp.ProcessingInfo.AuthorizationCode == got.ProcessingInfo.AuthorizationCode &&
		exp.ProcessingInfo.AVSResponseCodeZip == got.ProcessingInfo.AVSResponseCodeZip &&
		exp.ProcessingInfo.AVSResponseCodeAddress == got.ProcessingInfo.AVSResponseCodeAddress &&
		exp.ProcessingInfo.AVSResponseCodeName == got.ProcessingInfo.AVSResponseCodeName &&
		exp.ProcessingInfo.NetworkTransactionId == got.ProcessingInfo.NetworkTransactionId &&
		exp.SoftDescriptor == got.SoftDescriptor &&
		exp.CardHolderInfo.FirstName == got.CardHolderInfo.FirstName &&
		exp.CardHolderInfo.LastName == got.CardHolderInfo.LastName &&
		exp.CardHolderInfo.Zip == got.CardHolderInfo.Zip &&
		exp.Currency == got.Currency &&
		exp.FraudResultInfo.DeviceDataCollector == got.FraudResultInfo.DeviceDataCollector &&
		exp.CreditCard.CardRegulated == got.CreditCard.CardRegulated &&
		exp.CreditCard.CardLastFourDigits == got.CreditCard.CardLastFourDigits &&
		exp.CreditCard.CardType == got.CreditCard.CardType &&
		exp.CreditCard.CardSubType == got.CreditCard.CardSubType &&
		exp.CreditCard.CardCategory == got.CreditCard.CardCategory &&
		exp.CreditCard.BinCategory == got.CreditCard.BinCategory &&
		exp.CreditCard.BinNumber == got.CreditCard.BinNumber &&
		exp.CreditCard.IssuingBank == got.CreditCard.IssuingBank &&
		exp.CreditCard.IssuingCountryCode == got.CreditCard.IssuingCountryCode &&
		exp.CreditCard.ExpirationMonth == got.CreditCard.ExpirationMonth &&
		exp.CreditCard.ExpirationYear == got.CreditCard.ExpirationYear &&
		exp.CardTransactionType == got.CardTransactionType
}
