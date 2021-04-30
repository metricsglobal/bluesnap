package bluesnap

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/metricsglobal/bluesnap/card"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	testURL         = "https://sandbox.bluesnap.com"
	testCredentials = "QVBJXzE0NDQ2NTAyMDMxNDQ5NDA0MjIzNjU6QkxVRTEyMw=="
)

func TestCardAuth(t *testing.T) {
	scenarios := []struct {
		name    string
		input   Serializer
		output  Deserializer
		err     error
		wantErr bool
	}{
		{ // Basic
			name: "Basic",
			input: card.Request{
				CardTransactionType: "AUTH_CAPTURE",
				SoftDescriptor:      "DescTest",
				Amount:              "11.00",
				Currency:            "USD",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				CreditCard: &card.CreditCardRequest{
					CardNumber:      "4263982640269299",
					SecurityCode:    "837",
					ExpirationMonth: "02",
					ExpirationYear:  "2023",
				},
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				TransactionID:           "1035511869",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11,
				USDAmount:               11,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:08:47",
				CardHolderInfo: card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				VaultedShopperID: 28855193,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "9299",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					CardCategory:       "PLATINUM",
					BinCategory:        "CONSUMER",
					BinNumber:          "426398",
					ExpirationMonth:    "02",
					ExpirationYear:     "2023",
					CardRegulated:      "N",
					IssuingBank:        "ALLIED IRISH BANKS PLC",
					IssuingCountryCode: "ie",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					CVVResponseCode:        "MA",
					AuthorizationCode:      "123456",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		{ // Basic with vendors info
			name: "Basic with vendors info",
			input: card.Request{
				Amount:         "100",
				SoftDescriptor: "DescTest",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "01003",
					Country:   "US",
					State:     "MA",
				},
				Currency: "USD",
				CreditCard: &card.CreditCardRequest{
					CardNumber:      "4263982640269299",
					SecurityCode:    "111",
					ExpirationMonth: "02",
					ExpirationYear:  "2023",
				},
				CardTransactionType: "AUTH_CAPTURE",
				VendorsInfo: &card.VendorsInfo{
					VendorInfo: []card.VendorInfo{
						{
							VendorId: 10398032,
						},
					},
				},
			},
			output: &card.Response{
				CardTransactionType: "AUTH_CAPTURE",
				VendorsInfo: card.VendorsInfo{
					VendorInfo: []card.VendorInfo{
						{
							VendorId:          10398032,
							CommissionPercent: 90.2,
						},
					},
				},
				TransactionID:           "1035511531",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  100,
				USDAmount:               100,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:11:56",
				CardHolderInfo: card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "01003",
					Country:   "US",
					State:     "MA",
				},
				VaultedShopperID: 28855203,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "9299",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					CardCategory:       "PLATINUM",
					BinCategory:        "CONSUMER",
					CardRegulated:      "N",
					BinNumber:          "426398",
					ExpirationMonth:    "02",
					ExpirationYear:     "2023",
					IssuingBank:        "ALLIED IRISH BANKS PLC",
					IssuingCountryCode: "ie",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					CVVResponseCode:        "ND",
					AuthorizationCode:      "654321",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},

				VendorInfo: card.VendorInfo{
					VendorId:          10398032,
					CommissionPercent: 90.2,
				},
			},
		},
		{ // Basic with metadata
			name: "Basic with metadata",
			input: card.Request{
				SoftDescriptor: "DescTest",
				Amount:         "11",
				Currency:       "USD",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				CreditCard: &card.CreditCardRequest{
					CardNumber:      "4263982640269299",
					SecurityCode:    "837",
					ExpirationMonth: "02",
					ExpirationYear:  "2023",
				},
				CardTransactionType: "AUTH_CAPTURE",
				TransactionMetaData: &card.TransactionMetadata{
					Metadata: []card.Metadata{
						{
							MetaValue:       "20",
							MetaKey:         "stateTaxAmount",
							MetaDescription: "State Tax Amount",
						},
						{
							MetaValue:       "20",
							MetaKey:         "cityTaxAmount",
							MetaDescription: "City Tax Amount",
						},
						{
							MetaValue:       "10",
							MetaKey:         "shippingAmount",
							MetaDescription: "Shipping Amount",
						},
					},
				},
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				TransactionID:           "1035511881",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11,
				USDAmount:               11,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:14:14",
				CardHolderInfo: card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				VaultedShopperID: 28855211,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "9299",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					CardCategory:       "PLATINUM",
					BinCategory:        "CONSUMER",
					BinNumber:          "426398",
					ExpirationMonth:    "02",
					ExpirationYear:     "2023",
					CardRegulated:      "N",
					IssuingBank:        "ALLIED IRISH BANKS PLC",
					IssuingCountryCode: "ie",
				},
				TransactionMetadata: card.TransactionMetadata{
					Metadata: []card.Metadata{
						{
							MetaKey:         "stateTaxAmount",
							MetaValue:       "20",
							MetaDescription: "State Tax Amount",
						},
						{
							MetaKey:         "cityTaxAmount",
							MetaValue:       "20",
							MetaDescription: "City Tax Amount",
						},
						{
							MetaKey:         "shippingAmount",
							MetaValue:       "10",
							MetaDescription: "Shipping Amount",
						},
					},
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					CVVResponseCode:        "MA",
					AuthorizationCode:      "123456",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		{ // Basic with encrypted credit card
			name: "Basic with encrypted credit card",
			input: card.Request{
				SoftDescriptor: "DescTest",
				Amount:         "11",
				Currency:       "USD",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				CardTransactionType: "AUTH_CAPTURE",
				CreditCard: &card.CreditCardRequest{
					ExpirationYear:        "2023",
					ExpirationMonth:       "07",
					EncryptedCardNumber:   "$bsjs_1_0_3$B23uuxq8drUwOYZm3wZi+Qm69V5GPEt8PEio+Edwcm8akczQSK7odgLQH/Au+VqOCsGspW1Q9mPyQIzGLSZLVToAQVfq5C1ld+2ogIIsDL32Hd6IojboLyVlYT1FvPQoDyz19K6N0CUHh5uk0kCLuHSUyjvoJH38ojHZifbJSm/7S5vAtiuC3BJt2z8k9nauQaAXkbyoAYwrS1yDpqOt2k2lGhKcmdQ4ImDR0RL8m8xig6sFrki9oqo3Mju/M5r7wXXVTf7TMtWiQbzdfREOxKUnviXJZpncdHqVjj5GvPYun2qgopKVKr8F5+yd19TVW2gvA1kXBkXonFL9159Gxg==$zckJgo2i8jXDiAHwVVHBKypXFnWqF2e+6luBkmtQQRKniDXyXaalRVKLtYscBaGd$W7Ojqk1Q2iOJVeGL39RAsZTtfup3f1deSzvxrvC9rXA=",
					EncryptedSecurityCode: "$bsjs_1_0_3$MB1nBpok/YkuWPG1/7e6dyFFhDPHB8p8E9Yo+0YHHV+xkHuzFKr02wAnE8PJ8QCzWH+2ctXy5FN6wLKjwFrfTOgy0BJ9k9+NDEe8mhsu66wMlyc3lnwrbvMRCWN1O+5gUNCFExj7B0mDtf4gtxecXs74KZ5l5dbpGWdKUk5i7OewWyTqsONbn9taLfVBOwuIOy2Jgi4fx+yB8Q05KdZeHSNSBJh8H/47AUNAn5dM+d9iO6yGQB3obzEzzR3UtHlkGR52ZsgbbFh0JMm9lBM2ClgYM8jvmQjS9HX2ojt1fkbhuPEb1IY/M498a+1wDPpI4aMfDxO1lSpJneRSpY5k4g==$XaKq1NbPcS0iHy9N9jHekEIByHYS4G3wJXlC9EQjAGM=$BJn6X6mBYGUo8Eoq4RQz69gsi4Azl8jT973mNpG9Yuo=",
				},
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				TransactionID:           "1035511971",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11,
				USDAmount:               11,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:16:27",
				CardHolderInfo: card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				VaultedShopperID: 28855215,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "1111",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					BinCategory:        "CONSUMER",
					BinNumber:          "411111",
					ExpirationMonth:    "07",
					ExpirationYear:     "2023",
					CardRegulated:      "Y",
					IssuingBank:        "JPMORGAN CHASE BANK, N.A.",
					IssuingCountryCode: "us",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					CVVResponseCode:        "ND",
					AuthorizationCode:      "654321",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		{ // Basic with merchant shopper ID
			name: "Basic with merchant shopper ID",
			input: card.Request{
				SoftDescriptor: "DescTest",
				Amount:         "11",
				Currency:       "USD",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName:         "test first name",
					LastName:          "test last name",
					Zip:               "123456",
					MerchantShopperID: randStringRunes(20),
				},
				MerchantTransactionID: "31233",
				CreditCard: &card.CreditCardRequest{
					CardNumber:      "4263982640269299",
					SecurityCode:    "837",
					ExpirationMonth: "02",
					ExpirationYear:  "2023",
				},
				CardTransactionType: "AUTH_CAPTURE",
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				MerchantTransactionId:   "31233",
				TransactionID:           "1035511537",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11,
				USDAmount:               11,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:17:56",
				CardHolderInfo: card.CardHolderInfo{
					MerchantShopperID: "a432512",
					FirstName:         "test first name",
					LastName:          "test last name",
					Zip:               "123456",
				},
				VaultedShopperID: 28855219,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "9299",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					CardCategory:       "PLATINUM",
					BinCategory:        "CONSUMER",
					BinNumber:          "426398",
					ExpirationMonth:    "02",
					ExpirationYear:     "2023",
					CardRegulated:      "N",
					IssuingBank:        "ALLIED IRISH BANKS PLC",
					IssuingCountryCode: "ie",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					CVVResponseCode:        "MA",
					AuthorizationCode:      "123456",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		//{ // Basic with Hosted Payment Fields token
		//	name: "Basic with Hosted Payment Fields token",
		//	input: card.Request{
		//		SoftDescriptor: "DescTest",
		//		Amount:         "11",
		//		Currency:       "USD",
		//		CardHolderInfo: &card.CardHolderInfo{
		//			FirstName: "test first name",
		//			LastName:  "test last name",
		//			Zip:       "123456",
		//		},
		//		CardTransactionType: "AUTH_CAPTURE",
		//		// TODO MISSING VALID PF TOKEN
		//		PFToken:             "07ca959efe86d79f70919edc6b057e18b886096e72cb0a7f50707ff6c2e4a739_",
		//	},
		//	output: &card.Response{
		//		Amount:                  11,
		//		USDAmount:               11,
		//		Currency:                "USD",
		//		TransactionApprovalDate: "09/29/2020",
		//		TransactionApprovalTime: "13:16:27",
		//		VaultedShopperID:        1234,
		//		ProcessingInfo: card.ProcessingInfo{
		//			AVSResponseCodeAddress: "M",
		//			ProcessingStatus:       "success",
		//			CVVResponseCode:        "MA",
		//			AVSResponseCodeZip:     "U",
		//			AVSResponseCodeName:    "M",
		//		},
		//		SoftDescriptor: "BLS*DescTest",
		//		CardHolderInfo: card.CardHolderInfo{
		//			FirstName: "test first name",
		//			LastName:  "test last name",
		//			Zip:       "123456",
		//		},
		//		CreditCard: card.CreditCardResponse{
		//			CardLastFourDigits: "9299",
		//			CardSubType:        "CREDIT",
		//			CardType:           "VISA",
		//			CardCategory:       "CLASSIC",
		//		},
		//		CardTransactionType: "AUTH_CAPTURE",
		//		TransactionID:       "38486450",
		//	},
		//},
		{ // Basic with merchant transaction ID
			name: "Basic with merchant transaction ID",
			input: card.Request{
				SoftDescriptor: "DescTest",
				Amount:         "11",
				Currency:       "USD",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				MerchantTransactionID: "3",
				CreditCard: &card.CreditCardRequest{
					CardNumber:      "4263982640269299",
					SecurityCode:    "837",
					ExpirationMonth: "02",
					ExpirationYear:  "2023",
				},
				CardTransactionType: "AUTH_CAPTURE",
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				MerchantTransactionId:   "3",
				TransactionID:           "1035511807",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11,
				USDAmount:               11,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:23:19",
				CardHolderInfo: card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				VaultedShopperID: 28855261,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "9299",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					CardCategory:       "PLATINUM",
					BinCategory:        "CONSUMER",
					BinNumber:          "426398",
					ExpirationMonth:    "02",
					ExpirationYear:     "2023",
					CardRegulated:      "N",
					IssuingBank:        "ALLIED IRISH BANKS PLC",
					IssuingCountryCode: "ie",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					CVVResponseCode:        "MA",
					AuthorizationCode:      "123456",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		//{ // Basic with Embedded checkout token
		//	input: card.Request{
		//		Amount:              "11",
		//		SoftDescriptor:      "DescTest",
		//		Currency:            "USD",
		//		CardTransactionType: "AUTH_ONLY",
		//		PFToken:             "abcde12345**********", // TODO missing valid PF token
		//	},
		//	output: &card.Response{
		//		Amount:                  11,
		//		USDAmount:               11,
		//		Currency:                "USD",
		//		TransactionApprovalDate: "09/29/2020",
		//		TransactionApprovalTime: "13:23:19",
		//		VaultedShopperID:        1234,
		//		ProcessingInfo: card.ProcessingInfo{
		//			AVSResponseCodeAddress: "M",
		//			ProcessingStatus:       "success",
		//			CVVResponseCode:        "MA",
		//			AuthorizationCode:      "123456",
		//			AVSResponseCodeName:    "U",
		//			AVSResponseCodeZip:     "M",
		//		},
		//		SoftDescriptor: "BLS*DescTest",
		//		CardHolderInfo: card.CardHolderInfo{
		//			FirstName: "test first name",
		//			LastName:  "test last name",
		//			Zip:       "123456",
		//		},
		//		CreditCard: card.CreditCardResponse{
		//			CardLastFourDigits: "9299",
		//			CardSubType:        "CREDIT",
		//			CardType:           "VISA",
		//			CardCategory:       "CLASSIC",
		//		},
		//		CardTransactionType: "AUTH_CAPTURE",
		//		TransactionID:       "38486450",
		//	},
		//},
		{ // Basic with store card approval
			name: "Basic with store card approval",
			input: card.Request{
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
						FirstName: "YY",
						LastName:  "LL",
						Address1:  "Address1",
						Address2:  "Address2",
						City:      "Juneau",
						State:     "AL",
						Zip:       "12345",
						Country:   "US",
					},
				},
				CreditCard: &card.CreditCardRequest{
					ExpirationYear:  "2023",
					SecurityCode:    "837",
					ExpirationMonth: "02",
					CardNumber:      "4263982640269299",
				},
				CardTransactionType: "AUTH_CAPTURE",
				StoreCard:           true,
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				TransactionID:           "1035511921",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11,
				USDAmount:               11,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:23:19",
				CardHolderInfo: card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				VaultedShopperID: 28855295,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "9299",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					CardCategory:       "PLATINUM",
					BinCategory:        "CONSUMER",
					BinNumber:          "426398",
					ExpirationMonth:    "02",
					ExpirationYear:     "2023",
					CardRegulated:      "N",
					IssuingBank:        "ALLIED IRISH BANKS PLC",
					IssuingCountryCode: "ie",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					CVVResponseCode:        "MA",
					AuthorizationCode:      "123456",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		{ // Fraud request Basic with fraud info
			name: "Fraud request Basic with fraud info",
			input: card.Request{
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
					FraudProducts: []card.FraudProduct{
						{
							FraudProductName:     "123RRC",
							FraudProductDesc:     "my product",
							FraudProductType:     "Online game",
							FraudProductQuantity: 1,
							FraudProductPrice:    14.5,
						},
						{
							FraudProductName:     "345RRC",
							FraudProductDesc:     "my product2",
							FraudProductType:     "Video game",
							FraudProductQuantity: 2,
							FraudProductPrice:    18,
						},
					},
					ShippingContactInfo: &card.ShippingContactInfo{
						FirstName: "YY",
						LastName:  "LL",
						Address1:  "Address1",
						Address2:  "Address2",
						City:      "Juneau",
						State:     "AL",
						Zip:       "12345",
						Country:   "US",
					},
				},
				CreditCard: &card.CreditCardRequest{
					ExpirationYear:  "2023",
					SecurityCode:    "837",
					ExpirationMonth: "02",
					CardNumber:      "4263982640269299",
				},
				CardTransactionType: "AUTH_CAPTURE",
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				TransactionID:           "1035511921",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11.0,
				USDAmount:               11.0,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:23:19",
				CardHolderInfo: card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				VaultedShopperID: 28855295,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "9299",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					CardCategory:       "PLATINUM",
					BinCategory:        "CONSUMER",
					BinNumber:          "426398",
					CardRegulated:      "N",
					IssuingBank:        "ALLIED IRISH BANKS PLC",
					ExpirationMonth:    "02",
					ExpirationYear:     "2023",
					IssuingCountryCode: "ie",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					CVVResponseCode:        "MA",
					AuthorizationCode:      "123456",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		{ // Fraud request with Complete-level (formerly Enterprise) fraud info
			name: "Fraud request with Complete-level (formerly Enterprise) fraud info",
			input: card.Request{
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
						UDFs: []card.UDF{
							{
								Value: "aaa",
								Name:  "ENT_UDF1",
							},
							{
								Value: "bbb",
								Name:  "ENT_UDF2",
							},
						},
					},
					Company: "BBBBB",
					ShippingContactInfo: &card.ShippingContactInfo{
						FirstName: "YY",
						LastName:  "LL",
						Address1:  "Address1",
						Address2:  "Address2",
						City:      "Juneau",
						State:     "AL",
						Zip:       "12345",
						Country:   "US",
					},
				},
				CreditCard: &card.CreditCardRequest{
					ExpirationYear:  "2023",
					SecurityCode:    "837",
					ExpirationMonth: "02",
					CardNumber:      "4263982640269299",
				},
				CardTransactionType: "AUTH_CAPTURE",
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				TransactionID:           "1035511921",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11.0,
				USDAmount:               11.0,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:23:19",
				CardHolderInfo: card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				VaultedShopperID: 28855295,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "9299",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					CardCategory:       "PLATINUM",
					BinCategory:        "CONSUMER",
					BinNumber:          "426398",
					CardRegulated:      "N",
					IssuingBank:        "ALLIED IRISH BANKS PLC",
					ExpirationMonth:    "02",
					ExpirationYear:     "2023",
					IssuingCountryCode: "ie",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					CVVResponseCode:        "MA",
					AuthorizationCode:      "123456",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		{ // Vault Shopper request with vaulted shopper and credit card specified
			name: "Vault Shopper request with vaulted shopper and credit card specified",
			input: card.Request{
				Amount:           "11",
				VaultedShopperID: 20769005,
				SoftDescriptor:   "DescTest",
				Currency:         "USD",
				CreditCard: &card.CreditCardRequest{
					CardLastFourDigits: "1111",
					CardType:           "VISA",
				},
				CardTransactionType: "AUTH_CAPTURE",
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				TransactionID:           "1035511921",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11,
				USDAmount:               11,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:23:19",
				VaultedShopperID:        20769005,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "1111",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					BinCategory:        "CONSUMER",
					BinNumber:          "411111",
					ExpirationMonth:    "7",
					ExpirationYear:     "2022",
					CardRegulated:      "Y",
					IssuingBank:        "JPMORGAN CHASE BANK, N.A.",
					IssuingCountryCode: "us",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					AuthorizationCode:      "654321",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		{ // Vault Shopper request with vaulted shopper who has one card
			name: "Vault Shopper request with vaulted shopper who has one card",
			input: card.Request{
				Amount:              "11",
				VaultedShopperID:    20781033,
				SoftDescriptor:      "DescTest",
				Currency:            "USD",
				CardTransactionType: "AUTH_CAPTURE",
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				TransactionID:           "1035511921",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11,
				USDAmount:               11,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:23:19",
				VaultedShopperID:        20769005,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "0026",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					BinCategory:        "CONSUMER",
					CardCategory:       "CLASSIC",
					BinNumber:          "401200",
					ExpirationMonth:    "7",
					ExpirationYear:     "2021",
					CardRegulated:      "N",
					//IssuingBank:        "JPMORGAN CHASE BANK, N.A.", // TODO missing from actual resp
					IssuingCountryCode: "ru",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					AuthorizationCode:      "654321",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		{ // Vault Shopper request with vaulted shopper and CVV
			name: "Vault Shopper request with vaulted shopper and CVV",
			input: card.Request{
				Amount:           "11",
				VaultedShopperID: 20769005,
				SoftDescriptor:   "DescTest",
				Currency:         "USD",
				CreditCard: &card.CreditCardRequest{
					CardLastFourDigits: "1111",
					CardType:           "VISA",
					SecurityCode:       "837",
				},
				CardTransactionType: "AUTH_CAPTURE",
			},
			output: &card.Response{
				CardTransactionType:     "AUTH_CAPTURE",
				TransactionID:           "1035512361",
				SoftDescriptor:          "BLS*DescTest",
				Amount:                  11,
				USDAmount:               11,
				Currency:                "USD",
				TransactionApprovalDate: "09/29/2020",
				TransactionApprovalTime: "13:23:19",
				VaultedShopperID:        20769005,
				CreditCard: card.CreditCardResponse{
					CardLastFourDigits: "1111",
					CardType:           "VISA",
					CardSubType:        "CREDIT",
					BinCategory:        "CONSUMER",
					BinNumber:          "411111",
					ExpirationMonth:    "7",
					ExpirationYear:     "2022",
					CardRegulated:      "Y",
					IssuingBank:        "JPMORGAN CHASE BANK, N.A.",
					IssuingCountryCode: "us",
				},
				ProcessingInfo: card.ProcessingInfo{
					ProcessingStatus:       "success",
					CVVResponseCode:        "ND",
					AuthorizationCode:      "654321",
					AVSResponseCodeZip:     "U",
					AVSResponseCodeAddress: "U",
					AVSResponseCodeName:    "U",
				},
			},
		},
		//// TODO FIX ME
		////{ // Wallet request with wallet id
		////	input: card.Request{
		////		WalletID:            21,
		////		Amount:              "11",
		////		Currency:            "USD",
		////		CardTransactionType: "AUTH_CAPTURE",
		////	},
		////},
		// TODO INVALID_WALLET_PAYMENT_DATA
		//{ // Wallet request with Apple pay
		//	name: "Wallet request with Apple pay",
		//	input: card.Request{
		//		CardTransactionType: "AUTH_CAPTURE",
		//		SoftDescriptor:      "DescTest",
		//		Amount:              "11.0",
		//		Currency:            "USD",
		//		Wallet: &card.WalletRequest{
		//			WalletType:          "APPLE_PAY",
		//			EncodedPaymentToken: "ImRhdGEiOiJuY1AvRitIUy8zeG5bXhCMFd",
		//		},
		//	},
		//	output: &card.Response{
		//		CardTransactionType:     "AUTH_CAPTURE",
		//		MerchantTransactionId:   "112233",
		//		TransactionID:           "38602972",
		//		SoftDescriptor:          "BLS*DescTest",
		//		Amount:                  11,
		//		USDAmount:               11,
		//		Currency:                "USD",
		//		TransactionApprovalDate: "09/29/2020",
		//		TransactionApprovalTime: "13:23:19",
		//		CardHolderInfo: card.CardHolderInfo{
		//			FirstName: "John",
		//			LastName:  "Doe",
		//			Email:     "john@bluesnap.com",
		//			Country:   "us",
		//			State:     "NY",
		//			Address:   "61 Main St.",
		//			City:      "New York",
		//			Zip:       "12345",
		//			Phone:     "222513654654",
		//		},
		//		VaultedShopperID: 20769005,
		//		Wallet: card.WalletResponse{
		//			WalletType: "APPLE_PAY",
		//			BillingContactInfo: &card.BillingContactInfo{
		//				FirstName: "John",
		//				LastName:  "Doe",
		//				Address1:  "61 Main St.",
		//				City:      "New York",
		//				State:     "NY",
		//				Zip:       "12345",
		//				Country:   "us",
		//			},
		//			TokenizedCard: &card.TokenizedCard{
		//				DPANExpirationMonth: "7",
		//				DPANExpirationYear:  "2019",
		//				DPANLastFourDigits:  "0010",
		//				CardLastFourDigits:  "1471",
		//				CardType:            "MASTERCARD",
		//				CardSubType:         "DEBIT",
		//				BinCategory:         "CONSUMER",
		//				CardRegulated:       "N",
		//				IssuingCountryCode:  "us",
		//			},
		//		},
		//		ProcessingInfo: card.ProcessingInfo{
		//			ProcessingStatus:       "success",
		//			CVVResponseCode:        "ND",
		//			AuthorizationCode:      "654321",
		//			AVSResponseCodeZip:     "U",
		//			AVSResponseCodeAddress: "U",
		//			AVSResponseCodeName:    "U",
		//		},
		//	},
		//},
		//{ // Wallet request with Google pay
		//	input: card.Request{
		//		CardTransactionType: "AUTH_CAPTURE",
		//		SoftDescriptor:      "DescTest",
		//		Amount:              "11.0",
		//		Currency:            "USD",
		//		Wallet: &card.WalletRequest{
		//			WalletType:          "GOOGLE_PAY",
		//			EncodedPaymentToken: "ImRhdGEiOiJuY1AvRitIUy8zeG5bXhCMFd",
		//		},
		//	},
		//	output: &card.Response{
		//		CardTransactionType:     "AUTH_CAPTURE",
		//		TransactionID:           "38602972",
		//		SoftDescriptor:          "BLS&#x2a;DescTest",
		//		Amount:                  11,
		//		USDAmount:               11,
		//		Currency:                "USD",
		//		TransactionApprovalDate: "09/29/2020",
		//		TransactionApprovalTime: "13:23:19",
		//		CardHolderInfo: card.CardHolderInfo{
		//			FirstName: "test first name",
		//			LastName:  "test last name",
		//			Zip:       "12345",
		//			Country:   "us",
		//		},
		//		VaultedShopperID: 21289371,
		//		Wallet: card.WalletResponse{
		//			WalletType: "GOOGLE_PAY",
		//			BillingContactInfo: &card.BillingContactInfo{
		//				FirstName: "test first name",
		//				LastName:  "test last name",
		//				Zip:       "12345",
		//				Country:   "us",
		//			},
		//			TokenizedCard: &card.TokenizedCard{
		//				CardLastFourDigits:  "1111",
		//				CardType:            "VISA",
		//				CardSubType:         "CREDIT",
		//				DPANExpirationMonth: "9",
		//				DPANExpirationYear:  "2025",
		//				DPANLastFourDigits:  "2222",
		//			},
		//		},
		//		ProcessingInfo: card.ProcessingInfo{
		//			ProcessingStatus: "success",
		//		},
		//	},
		//},
		//{ // 3-D Secure request with shopper-initiated transaction
		//	input: card.Request{
		//		Amount:         "10.0",
		//		SoftDescriptor: "DescTest",
		//		Currency:       "USD",
		//		CardHolderInfo: &card.CardHolderInfo{
		//			FirstName: "Will",
		//			LastName:  "Smith",
		//		},
		//		PFToken:              "124ed4af4a7530e1a5f6359c3120cde7b05260d792a3aeed38acd098046846af_",
		//		CardTransactionType:  "AUTH_CAPTURE",
		//		TransactionInitiator: "SHOPPER",
		//	},
		//	output: &card.Response{
		//		CardTransactionType:     "AUTH_CAPTURE",
		//		TransactionID:           "38510976",
		//		SoftDescriptor:          "BLS*DescTest",
		//		Amount:                  10,
		//		USDAmount:               10,
		//		Currency:                "USD",
		//		TransactionApprovalDate: "09/29/2020",
		//		TransactionApprovalTime: "13:23:19",
		//		AVSResponseCode:         "G",
		//		CardHolderInfo: card.CardHolderInfo{
		//			FirstName: "Will",
		//			LastName:  "Smith",
		//		},
		//		VaultedShopperID: 19574632,
		//		CreditCard: card.CreditCardResponse{
		//			CardLastFourDigits: "1000",
		//			CardType:           "VISA",
		//			CardSubType:        "CREDIT",
		//			BinCategory:        "CONSUMER",
		//			IssuingCountryCode: "us",
		//		},
		//		ThreeDSecure: &card.ThreeDSecureResponse{
		//			AuthenticationResult: "AUTHENTICATION_SUCCEEDED",
		//		},
		//		// TODO
		//		ProcessingInfo: card.ProcessingInfo{
		//			ProcessingStatus:       "success",
		//			CVVResponseCode:        "NR",
		//			AVSResponseCodeZip:     "U",
		//			AVSResponseCodeAddress: "U",
		//			AVSResponseCodeName:    "U",
		//			NetworkTransactionId:   "019072416113666",
		//		},
		//	},
		//},
		//{ // 3-D Secure request with MIT
		//	input: card.Request{
		//		Amount:               "25.0",
		//		SoftDescriptor:       "DescTest",
		//		Currency:             "USD",
		//		VaultedShopperID:     19574632,
		//		CardTransactionType:  "AUTH_CAPTURE",
		//		TransactionInitiator: "MERCHANT",
		//		NetworkTransactionInfo: map[string]string{
		//			"originalNetworkTransactionId": "019072416113666",
		//			"networkTransactionId":         "019072416144666",
		//		},
		//	},
		//},
		//{ // 3-D Secure request returning shopper with 3DS
		//	input: card.Request{
		//		CardTransactionType: "AUTH_CAPTURE",
		//		SoftDescriptor:      "DescTest",
		//		Amount:              "10",
		//		Currency:            "USD",
		//		VaultedShopperID:    19574268,
		//		CreditCard: &card.CreditCardRequest{
		//			CardLastFourDigits: "1111",
		//			CardType:           "VISA",
		//		},
		//		ThreeDSecure: &card.ThreeDSecureRequest{
		//			ThreeDSecureReferenceID: "5903",
		//		},
		//	},
		//},
		//{ // 3-D Secure request with full card details with 3DS id
		//	input: card.Request{
		//		CardTransactionType:  "AUTH_CAPTURE",
		//		RecurringTransaction: "ECOMMERCE",
		//		SoftDescriptor:       "DescTest",
		//		Amount:               "10",
		//		Currency:             "USD",
		//		CreditCard: &card.CreditCardRequest{
		//			CardNumber:      "4012000033330026",
		//			SecurityCode:    "123",
		//			ExpirationMonth: "07",
		//			ExpirationYear:  "2023",
		//		},
		//		ThreeDSecure: &card.ThreeDSecureRequest{
		//			ThreeDSecureReferenceID: "4759",
		//		},
		//	},
		//},
	}

	c := New(testURL, testCredentials)
	for _, scenario := range scenarios {
		resp := card.Response{}
		if err := c.do("POST", "/services/2/transactions", scenario.input, &resp); err != nil {
			t.Errorf(err.Error())
		}
		if scenario.output == nil {
			continue
		}

		compareResponses(t, *(scenario.output.(*card.Response)), resp)
	}
}

func compareResponses(t *testing.T, expected, actual card.Response) {
	equalsString(t, "cardTransactionType", expected.CardTransactionType, actual.CardTransactionType)
	nonEmptyString(t, "transactionId", actual.TransactionID)
	equalsString(t, "softDescriptor", expected.SoftDescriptor, actual.SoftDescriptor)
	equalsFloat64(t, "Amount", expected.Amount, actual.Amount)
	equalsFloat64(t, "usdAmount", expected.USDAmount, actual.USDAmount)
	equalsString(t, "currency", expected.Currency, actual.Currency)
	nonEmptyString(t, "transactionApprovalDate", actual.TransactionApprovalDate)
	nonEmptyString(t, "transactionApprovalTime", actual.TransactionApprovalTime)
	compareCardholderInfo(t, expected.CardHolderInfo, actual.CardHolderInfo)
	nonEmptyInt64(t, "vaultedShopperId", actual.VaultedShopperID)
	compareCreditCard(t, expected.CreditCard, actual.CreditCard)
	compareWalletResponse(t, expected.Wallet, actual.Wallet)
	compareThreeDSecure(t, expected.ThreeDSecure, actual.ThreeDSecure)
	compareProcessingInfo(t, expected.ProcessingInfo, actual.ProcessingInfo)
	compareFraudResultInfo(t, expected.FraudResultInfo, actual.FraudResultInfo)
	compareVendorInfo(t, expected.VendorInfo, actual.VendorInfo)
	compareVendorsInfo(t, expected.VendorsInfo, actual.VendorsInfo)
	compareTransactionMetadata(t, expected.TransactionMetadata, actual.TransactionMetadata)
	equalsString(t, "merchantTransactionId", expected.MerchantTransactionId, actual.MerchantTransactionId)
	equalsString(t, "taxReference", expected.TaxReference, actual.TaxReference)
	equalsString(t, "avsResponseCode", expected.AVSResponseCode, actual.AVSResponseCode)
}

func compareCardholderInfo(t *testing.T, expected, actual card.CardHolderInfo) {
	equalsString(t, "firstName", expected.FirstName, actual.FirstName)
	equalsString(t, "lastName", expected.LastName, actual.LastName)
	equalsString(t, "email", expected.Email, actual.Email)
	equalsString(t, "country", expected.Country, actual.Country)
	equalsString(t, "state", expected.State, actual.State)
	equalsString(t, "address", expected.Address, actual.Address)
	equalsString(t, "address_2", expected.Address2, actual.Address2)
	equalsString(t, "city", expected.City, actual.City)
	equalsString(t, "zip", expected.Zip, actual.Zip)
	equalsString(t, "phone", expected.Phone, actual.Phone)
	if expected.MerchantShopperID != "" {
		nonEmptyString(t, "merchantShopperId", actual.MerchantShopperID)
	}
	equalsString(t, "personalIdentificationNumber", expected.PersonalIdentificationNumber, actual.PersonalIdentificationNumber)
	equalsString(t, "companyName", expected.CompanyName, actual.CompanyName)
}

func compareCreditCard(t *testing.T, expected, actual card.CreditCardResponse) {
	equalsString(t, "cardLastFourDigits", expected.CardLastFourDigits, actual.CardLastFourDigits)
	equalsString(t, "cardType", expected.CardType, actual.CardType)
	equalsString(t, "cardSubType", expected.CardSubType, actual.CardSubType)
	equalsString(t, "cardCategory", expected.CardCategory, actual.CardCategory)
	equalsString(t, "binCategory", expected.BinCategory, actual.BinCategory)
	equalsString(t, "binNumber", expected.BinNumber, actual.BinNumber)
	equalsString(t, "cardRegulated", expected.CardRegulated, actual.CardRegulated)
	equalsString(t, "issuingBank", expected.IssuingBank, actual.IssuingBank)
	equalsString(t, "issuingCountryCode", expected.IssuingCountryCode, actual.IssuingCountryCode)
	equalsString(t, "expirationMonth", expected.ExpirationMonth, actual.ExpirationMonth)
	equalsString(t, "expirationyear", expected.ExpirationYear, actual.ExpirationYear)
}

func compareWalletResponse(t *testing.T, expected, actual card.WalletResponse) {
	equalsString(t, "walletType", expected.WalletType, actual.WalletType)
	compareWalletBillingContactInfo(t, expected.BillingContactInfo, actual.BillingContactInfo)
	compareWalletTokenizedCard(t, expected.TokenizedCard, actual.TokenizedCard)
}

func compareWalletBillingContactInfo(t *testing.T, expected, actual *card.BillingContactInfo) {
	if expected == nil {
		nill(t, "walletBillingContactInfo", actual)
		return
	}
	notNil(t, "walletBillingInfo", actual)
	if actual == nil {
		return
	}
	equalsString(t, "firstName", expected.FirstName, actual.FirstName)
	equalsString(t, "lastName", expected.LastName, actual.LastName)
	equalsString(t, "address1", expected.Address1, actual.Address1)
	equalsString(t, "address2", expected.Address2, actual.Address2)
	equalsString(t, "city", expected.City, actual.City)
	equalsString(t, "state", expected.State, actual.State)
	equalsString(t, "zip", expected.Zip, actual.Zip)
	equalsString(t, "country", expected.Country, actual.Country)
	equalsString(t, "personalIdentificationNumber", expected.PersonalIdentificationNumber, actual.PersonalIdentificationNumber)
}

func compareWalletTokenizedCard(t *testing.T, expected, actual *card.TokenizedCard) {
	if expected == nil {
		nill(t, "walletTokenizedCard", actual)
		return
	}
	notNil(t, "walletCard", actual)
	if actual == nil {
		return
	}
	equalsString(t, "cardLastFourDigits", expected.CardLastFourDigits, actual.CardLastFourDigits)
	equalsString(t, "cardType", expected.CardType, actual.CardType)
	equalsString(t, "cardSubType", expected.CardSubType, actual.CardSubType)
	equalsString(t, "binCategory", expected.BinCategory, actual.BinCategory)
	equalsString(t, "cardRegulated", expected.CardRegulated, actual.CardRegulated)
	equalsString(t, "issuingCountryCode", expected.IssuingCountryCode, actual.IssuingCountryCode)
	equalsString(t, "dpanExpirationMonth", expected.DPANExpirationMonth, actual.DPANExpirationMonth)
	equalsString(t, "dpanExpirationYear", expected.DPANExpirationYear, actual.DPANExpirationYear)
	equalsString(t, "dpanLastFourDigits", expected.DPANLastFourDigits, actual.DPANLastFourDigits)
}

func compareTransactionMetadata(t *testing.T, expected, actual card.TransactionMetadata) {
	if expected.Metadata != nil {
		notNil(t, "transaction metadata", actual.Metadata)
	}
	for i, _ := range expected.Metadata {
		compareMetadata(t, expected.Metadata[i], actual.Metadata[i])
	}
	for i, _ := range actual.Metadata {
		compareMetadata(t, expected.Metadata[i], actual.Metadata[i])
	}
}

func compareMetadata(t *testing.T, expected, actual card.Metadata) {
	equalsString(t, "metaKey", expected.MetaKey, actual.MetaKey)
	equalsString(t, "metaValue", expected.MetaValue, actual.MetaValue)
	equalsString(t, "metaDescription", expected.MetaDescription, actual.MetaDescription)
	equalsString(t, "isVisible", expected.IsVisible, actual.IsVisible)
}

func compareVendorsInfo(t *testing.T, expected, actual card.VendorsInfo) {
	if expected.VendorInfo != nil {
		notNil(t, "vendors info", actual.VendorInfo)
	}
	for i, _ := range expected.VendorInfo {
		compareVendorInfo(t, expected.VendorInfo[i], actual.VendorInfo[i])
	}
	for i, _ := range actual.VendorInfo {
		compareVendorInfo(t, expected.VendorInfo[i], actual.VendorInfo[i])
	}
}

func compareVendorInfo(t *testing.T, expected, actual card.VendorInfo) {
	equalsInt64(t, "vendorId", expected.VendorId, actual.VendorId)
	equalsFloat64(t, "commissionPercent", expected.CommissionPercent, actual.CommissionPercent)
	equalsInt64(t, "commissionAmount", expected.CommissionAmount, actual.CommissionAmount)
}

func compareFraudResultInfo(t *testing.T, expected, actual card.FraudResultInfo) {
	equalsString(t, "deviceDataCollector", expected.DeviceDataCollector, actual.DeviceDataCollector)
}

func compareProcessingInfo(t *testing.T, expected, actual card.ProcessingInfo) {
	equalsString(t, "processingStatus", expected.ProcessingStatus, actual.ProcessingStatus)
	equalsString(t, "cvvResponseCode", expected.CVVResponseCode, actual.CVVResponseCode)
	equalsString(t, "authorizationCode", expected.AuthorizationCode, actual.AuthorizationCode)
	equalsString(t, "avsResponseCodeZip", expected.AVSResponseCodeZip, actual.AVSResponseCodeZip)
	equalsString(t, "avsResponseCodeAddress", expected.AVSResponseCodeAddress, actual.AVSResponseCodeAddress)
	equalsString(t, "avsResponseCodeName", expected.AVSResponseCodeName, actual.AVSResponseCodeName)
	equalsString(t, "networkTransactionId", expected.NetworkTransactionId, actual.NetworkTransactionId)
}

func compareThreeDSecure(t *testing.T, expected, actual *card.ThreeDSecureResponse) {
	if expected == nil {
		nill(t, "threeDSecure", actual)
		return
	}
	notNil(t, "3DauthenticationResult", actual)
	if actual == nil {
		return
	}
	equalsString(t, "authenticationResult", expected.AuthenticationResult, actual.AuthenticationResult)
}

func equalsString(t *testing.T, name, expected, actual string) {
	if actual != expected {
		t.Errorf("%s should be %s, instead of %s", name, expected, actual)
	}
}

func nonEmptyString(t *testing.T, name, val string) {
	if val == "" {
		t.Errorf("%s shouldn't be empty", name)
	}
}

func nonEmptyInt64(t *testing.T, name string, val int64) {
	if val == 0 {
		t.Errorf("%s shouldn't be empty", name)
	}
}

func equalsFloat64(t *testing.T, name string, expected, actual float64) {
	if actual != expected {
		t.Errorf("%s should be %f, instead of %f", name, expected, actual)
	}
}

func equalsInt64(t *testing.T, name string, expected, actual int64) {
	if actual != expected {
		t.Errorf("%s should be %d, instead of %d", name, expected, actual)
	}
}

func notNil(t *testing.T, name string, value interface{}) {
	if value == nil {
		t.Errorf("%s value should not be nil", name)
	}
}

func nill(t *testing.T, name string, value interface{}) {
	if value != nil && !(reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		t.Errorf("%s value should be nil %v", name, value)
	}
}

func randStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
