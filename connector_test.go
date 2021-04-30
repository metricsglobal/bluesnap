package bluesnap

import (
	"testing"

	"github.com/metricsglobal/bluesnap/card"
)

const (
	testURL         = "https://sandbox.bluesnap.com"
	testCredentials = "QVBJXzE0NDQ2NTAyMDMxNDQ5NDA0MjIzNjU6QkxVRTEyMw=="
)

func TestCardAuth(t *testing.T) {
	scenarios := []struct {
		input   Serializer
		output  Deserializer
		err     error
		wantErr bool
		testFn func(output Deserializer) error
	}{
		{ // Basic
			input: card.Request{
				CardTransactionType: "AUTH_CAPTURE",
				SoftDescriptor:      "DescTest",
				Amount:              "11.00",
				Currency:            "USD",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "01003",
					Country:   "US",
					State:     "MA",
				},
				CreditCard: &card.CreditCardRequest{
					CardNumber:      "4263982640269299",
					SecurityCode:    "837",
					ExpirationMonth: "02",
					ExpirationYear:  "2023",
				},
			},
			output: &card.Response{},
		},
		{ // Basic with vendors info
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
		},
		{ // Basic with metadata
			input: card.Request{
				SoftDescriptor:      "DescTest",
				Amount:              "11",
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
				CardTransactionType: "AUTH_CAPTURE",
				TransactionMetaData: &card.TransactionMetadata{
					Metadata: []card.Metadata{
						{
							MetaValue: "20",
							MetaKey: "stateTaxAmount",
							MetaDescription: "State Tax Amount",
						},
						{
							MetaValue: "20",
							MetaKey: "cityTaxAmount",
							MetaDescription: "City Tax Amount",
						},
						{
							MetaValue: "10",
							MetaKey: "shippingAmount",
							MetaDescription: "Shipping Amount",
						},
					},
				},
			},
			output: nil,
		},
		{ // Basic with encrypted credit card
			input: card.Request{
				SoftDescriptor:      "DescTest",
				Amount:              "11",
				Currency:            "USD",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				CardTransactionType: "AUTH_CAPTURE",
				CreditCard: &card.CreditCardRequest{
					ExpirationYear: "2023",
					ExpirationMonth: "07",
					EncryptedCardNumber: "$bsjs_1_0_3$B23uuxq8drUwOYZm3wZi+Qm69V5GPEt8PEio+Edwcm8akczQSK7odgLQH/Au+VqOCsGspW1Q9mPyQIzGLSZLVToAQVfq5C1ld+2ogIIsDL32Hd6IojboLyVlYT1FvPQoDyz19K6N0CUHh5uk0kCLuHSUyjvoJH38ojHZifbJSm/7S5vAtiuC3BJt2z8k9nauQaAXkbyoAYwrS1yDpqOt2k2lGhKcmdQ4ImDR0RL8m8xig6sFrki9oqo3Mju/M5r7wXXVTf7TMtWiQbzdfREOxKUnviXJZpncdHqVjj5GvPYun2qgopKVKr8F5+yd19TVW2gvA1kXBkXonFL9159Gxg==$zckJgo2i8jXDiAHwVVHBKypXFnWqF2e+6luBkmtQQRKniDXyXaalRVKLtYscBaGd$W7Ojqk1Q2iOJVeGL39RAsZTtfup3f1deSzvxrvC9rXA=",
					EncryptedSecurityCode: "$bsjs_1_0_3$MB1nBpok/YkuWPG1/7e6dyFFhDPHB8p8E9Yo+0YHHV+xkHuzFKr02wAnE8PJ8QCzWH+2ctXy5FN6wLKjwFrfTOgy0BJ9k9+NDEe8mhsu66wMlyc3lnwrbvMRCWN1O+5gUNCFExj7B0mDtf4gtxecXs74KZ5l5dbpGWdKUk5i7OewWyTqsONbn9taLfVBOwuIOy2Jgi4fx+yB8Q05KdZeHSNSBJh8H/47AUNAn5dM+d9iO6yGQB3obzEzzR3UtHlkGR52ZsgbbFh0JMm9lBM2ClgYM8jvmQjS9HX2ojt1fkbhuPEb1IY/M498a+1wDPpI4aMfDxO1lSpJneRSpY5k4g==$XaKq1NbPcS0iHy9N9jHekEIByHYS4G3wJXlC9EQjAGM=$BJn6X6mBYGUo8Eoq4RQz69gsi4Azl8jT973mNpG9Yuo=",
				},
			},
			output: nil,
		},
		{ // Basic with merchant shopper ID
			input: card.Request{
				SoftDescriptor:      "DescTest",
				Amount:              "11",
				Currency:            "USD",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
					MerchantShopperID: "a432512",
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
			output: nil,
		},
		{ // Basic with Hosted Payment Fields token
			input: card.Request{
				SoftDescriptor:      "DescTest",
				Amount:              "11",
				Currency:            "USD",
				CardHolderInfo: &card.CardHolderInfo{
					FirstName: "test first name",
					LastName:  "test last name",
					Zip:       "123456",
				},
				CardTransactionType: "AUTH_CAPTURE",
				PFToken: "abcde12345**********",
			},
			output: nil,
		},
		{ // Basic with merchant transaction ID
			input: card.Request{
				SoftDescriptor:      "DescTest",
				Amount:              "11",
				Currency:            "USD",
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
			output: nil,
		},
	}

	c := New(testURL, testCredentials)
	for _, scenario := range scenarios {
		if err := c.do("POST", "/services/2/transactions", scenario.input, scenario.output); err != nil {
			t.Fatal(err)
		}
		//fmt.Printf("%+v\n", scenario.output)
	}
}
