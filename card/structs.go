package card

type Request struct {
	WalletID               int64                        `json:"walletId,omitempty"`
	Wallet                 *WalletRequest               `json:"wallet,omitempty"`
	Amount                 string                       `json:"amount,omitempty"`
	VaultedShopperID       int64                        `json:"vaultedShopperId,omitempty"`
	MerchantTransactionID  string                       `json:"merchantTransactionId,omitempty"`
	SoftDescriptor         string                       `json:"softDescriptor,omitempty"`
	DescriptorPhoneNumber  string                       `json:"descriptorPhoneNumber,omitempty"`
	TaxReference           string                       `json:"taxReference,omitempty"`
	VendorsInfo            *VendorsInfo                 `json:"vendorsInfo,omitempty"`
	CardHolderInfo         *CardHolderInfo              `json:"cardHolderInfo,omitempty"`
	Currency               string                       `json:"currency,omitempty"`
	TransactionFraudInfo   *TransactionFraudInfoRequest `json:"transactionFraudInfo,omitempty"`
	CreditCard             *CreditCardRequest           `json:"creditCard,omitempty"`
	CardTransactionType    string                       `json:"cardTransactionType,omitempty"`
	ThreeDSecure           *ThreeDSecureRequest         `json:"threeDSecure,omitempty"`
	TransactionMetaData    *TransactionMetadata         `json:"transactionMetaData,omitempty"`
	PFToken                string                       `json:"pfToken,omitempty"`
	Level3Data             *Level3DataRequest           `json:"level3Data,omitempty"`
	StoreCard              bool                         `json:"storeCard,omitempty"`
	NetworkTransactionInfo map[string]string            `json:"networkTransactionInfo,omitempty"` // TODO
	TransactionOrderSource string                       `json:"transactionOrderSource,omitempty"`
	TransactionInitiator   string                       `json:"transactionInitiator,omitempty"`
	RecurringTransaction   string                       `json:"recurringTransaction"`
	TransactionID          string                       `json:"transactionId"`
}

type Response struct {
	CardTransactionType     string                `json:"cardTransactionType"`
	TransactionID           string                `json:"transactionId"`
	SoftDescriptor          string                `json:"softDescriptor"`
	Amount                  float64               `json:"Amount"`
	USDAmount               float64               `json:"usdAmount"`
	Currency                string                `json:"currency"`
	TransactionApprovalDate string                `json:"transactionApprovalDate"`
	TransactionApprovalTime string                `json:"transactionApprovalTime"`
	CardHolderInfo          CardHolderInfo        `json:"cardHolderInfo"`
	VaultedShopperID        int64                 `json:"vaultedShopperId"`
	CreditCard              CreditCardResponse    `json:"creditCard"`
	Wallet                  WalletResponse        `json:"wallet"`
	ThreeDSecure            *ThreeDSecureResponse `json:"threeDSecure,omitempty"`
	// TODO add NetworkTransactionInfo
	ProcessingInfo        ProcessingInfo      `json:"processingInfo"`
	FraudResultInfo       FraudResultInfo     `json:"fraudResultInfo"`
	VendorInfo            VendorInfo          `json:"vendorInfo"`
	VendorsInfo           VendorsInfo         `json:"vendorsInfo"`
	TransactionMetadata   TransactionMetadata `json:"transactionMetaData"`
	MerchantTransactionId string              `json:"merchantTransactionId"`
	TaxReference          string              `json:"taxReference"`
	AVSResponseCode       string              `json:"avsResponseCode"`
}

type WalletRequest struct {
	WalletType          string `json:"walletType,omitempty"`
	EncodedPaymentToken string `json:"encodedPaymentToken,omitempty"`
}

type WalletResponse struct {
	WalletType         string              `json:"walletType"`
	BillingContactInfo *BillingContactInfo `json:"billingContactInfo,omitempty"`
	TokenizedCard      *TokenizedCard      `json:"tokenizedCard,omitempty"`
}

// BillingContactInfo request and response struct
type BillingContactInfo struct {
	FirstName                    string `json:"firstName,omitempty"`
	LastName                     string `json:"lastName,omitempty"`
	Address1                     string `json:"address1,omitempty"`
	Address2                     string `json:"address2,omitempty"`
	City                         string `json:"city,omitempty"`
	State                        string `json:"state,omitempty"`
	Zip                          string `json:"zip,omitempty"`
	Country                      string `json:"country,omitempty"`
	PersonalIdentificationNumber string `json:"personalIdentificationNumber,omitempty"`
}

// ShippingContactInfo request and response struct
type ShippingContactInfo struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Country   string `json:"country,omitempty"`
}

// TokenizedCard request and response struct
type TokenizedCard struct {
	CardLastFourDigits  string `json:"cardLastFourDigits,omitempty"`
	CardType            string `json:"cardType,omitempty"`
	CardSubType         string `json:"cardSubType,omitempty"`
	BinCategory         string `json:"binCategory,omitempty"`
	CardRegulated       string `json:"cardRegulated,omitempty"`
	IssuingCountryCode  string `json:"issuingCountryCode,omitempty"`
	DPANExpirationMonth string `json:"dpanExpirationMonth,omitempty"`
	DPANExpirationYear  string `json:"dpanExpirationYear,omitempty"`
	DPANLastFourDigits  string `json:"dpanLastFourDigits,omitempty"`
}

// VendorsInfo request and response struct
type VendorsInfo struct {
	VendorInfo []VendorInfo `json:"vendorInfo,omitempty"`
}

type VendorInfo struct {
	VendorId          int64   `json:"vendorId,omitempty"`
	CommissionPercent float64 `json:"commissionPercent,omitempty"`
	CommissionAmount  int64   `json:"commissionAmount,omitempty"`
}

// CardHolderInfo request and response struct
type CardHolderInfo struct {
	FirstName                    string `json:"firstName,omitempty"`
	LastName                     string `json:"lastName,omitempty"`
	Email                        string `json:"email,omitempty"`
	Country                      string `json:"country,omitempty"`
	State                        string `json:"state,omitempty"`
	Address                      string `json:"address,omitempty"`
	Address2                     string `json:"address_2,omitempty"`
	City                         string `json:"city,omitempty"`
	Zip                          string `json:"zip,omitempty"`
	Phone                        string `json:"phone,omitempty"`
	MerchantShopperID            string `json:"merchantShopperId,omitempty"`
	PersonalIdentificationNumber string `json:"personalIdentificationNumber,omitempty"`
	CompanyName                  string `json:"companyName,omitempty"`
}

type TransactionFraudInfoRequest struct {
	FraudSessionId      string               `json:"fraudSessionId,omitempty"`
	ShopperIpAddress    string               `json:"shopperIpAddress,omitempty"`
	Company             string               `json:"company,omitempty"`
	ShippingContactInfo *ShippingContactInfo `json:"shippingContactInfo,omitempty"`
	FraudProducts       []FraudProduct       `json:"fraudProducts,omitempty"`
	EnterpriseSiteID    string               `json:"enterpriseSiteId,omitempty"`
	EnterpriseUDFs      *EnterpriseUDFs      `json:"enterpriseUdfs,omitempty"`
}

type FraudProduct struct {
	FraudProductName     string  `json:"fraudProductName,omitempty"`
	FraudProductDesc     string  `json:"fraudProductDesc,omitempty"`
	FraudProductType     string  `json:"fraudProductType,omitempty"`
	FraudProductQuantity int64   `json:"fraudProductQuantity,omitempty"`
	FraudProductPrice    float64 `json:"fraudProductPrice,omitempty"`
}

type EnterpriseUDFs struct {
	UDFs []UDF `json:"udf,omitempty"`
}

type UDF struct {
	Name  string `json:"udfName,omitempty"`
	Value string `json:"udfValue,omitempty"`
}

type CreditCardRequest struct {
	CardNumber            string `json:"cardNumber,omitempty"`
	EncryptedCardNumber   string `json:"encryptedCardNumber,omitempty"`
	CardLastFourDigits    string `json:"cardLastFourDigits,omitempty"`
	CardType              string `json:"cardType,omitempty"`
	ExpirationMonth       string `json:"expirationMonth,omitempty"`
	ExpirationYear        string `json:"expirationYear,omitempty"`
	SecurityCode          string `json:"securityCode,omitempty"`
	EncryptedSecurityCode string `json:"encryptedSecurityCode,omitempty"`
	SecurityCodePfToken   string `json:"securityCodePfToken,omitempty"`
}

type CreditCardResponse struct {
	CardLastFourDigits string `json:"cardLastFourDigits,omitempty"`
	CardType           string `json:"cardType,omitempty"`
	CardSubType        string `json:"cardSubType,omitempty"`
	CardCategory       string `json:"cardCategory,omitempty"`
	BinCategory        string `json:"binCategory,omitempty"`
	BinNumber          string `json:"binNumber,omitempty"`
	CardRegulated      string `json:"cardRegulated,omitempty"`
	IssuingBank        string `json:"issuingBank,omitempty"`
	IssuingCountryCode string `json:"issuingCountryCode,omitempty"`
	ExpirationMonth    string `json:"expirationMonth,omitempty"`
	ExpirationYear     string `json:"expirationyear,omitempty"`
}

type ThreeDSecureRequest struct {
	ThreeDSecureResultToken string `json:"threeDSecureResultToken,omitempty"`
	ECI                     string `json:"eci,omitempty"`
	CAVV                    string `json:"cavv,omitempty"`
	XID                     string `json:"xid,omitempty"`
	DSTransactionID         string `json:"dsTransactionId,omitempty"`
	ThreeDSecureVersion     string `json:"threeDSecureVersion,omitempty"`
	ThreeDSecureReferenceID string `json:"threeDSecureReferenceId,omitempty"`
}

type ThreeDSecureResponse struct {
	AuthenticationResult string `json:"authenticationResult,omitempty"`
}

// TransactionMetadata request and response struct
type TransactionMetadata struct {
	Metadata []Metadata `json:"metaData,omitempty"`
}

// Metadata request and response struct
type Metadata struct {
	MetaKey         string `json:"metaKey,omitempty"`
	MetaValue       string `json:"metaValue,omitempty"`
	MetaDescription string `json:"metaDescription,omitempty"`
	IsVisible       string `json:"isVisible,omitempty"`
}

type Level3DataRequest struct {
	CustomerReferenceNumber string           `json:"customerReferenceNumber,omitempty"`
	SalesTaxAmount          int64            `json:"salesTaxAmount,omitempty"`
	FreightAmount           int64            `json:"freightAmount,omitempty"`
	DutyAmount              int64            `json:"dutyAmount,omitempty"`
	DestinationZipCode      string           `json:"destinationZipCode,omitempty"`
	DestinationCountryCode  string           `json:"destinationCountryCode,omitempty"`
	ShipFromZipCode         string           `json:"shipFromZipCode,omitempty"`
	DiscountAmount          int64            `json:"discountAmount,omitempty"`
	TaxAmount               int64            `json:"taxAmount,omitempty"`
	TaxRate                 int64            `json:"taxRate,omitempty"`
	Level3DataItems         []Level3DataItem `json:"level3DataItems,omitempty"`
}

type Level3DataResponse struct {
	Level3DataRequest
	TransactionProcessedWithL3dSupportedAcquirer bool `json:"transactionProcessedWithL3dSupportedAcquirer,omitempty"`
}

type Level3DataItem struct {
	LineItemTotal     int64  `json:"lineItemTotal,omitempty"`
	CommodityCode     string `json:"commodityCode,omitempty"`
	Description       string `json:"description,omitempty"`
	DiscountAmount    int64  `json:"discountAmount,omitempty"`
	DiscountIndicator string `json:"discountIndicator,omitempty"`
	GrossNetIndicator string `json:"grossNetIndicator,omitempty"`
	ProductCode       string `json:"productCode,omitempty"`
	ItemQuantity      int64  `json:"itemQuantity,omitempty"`
	TaxAmount         int64  `json:"taxAmount,omitempty"`
	TaxRate           int64  `json:"taxRate,omitempty"`
	TaxType           string `json:"taxType,omitempty"`
	UnitCost          int64  `json:"unitCost,omitempty"`
	UnitOfMeasure     string `json:"unitOfMeasure,omitempty"`
}

type ProcessingInfo struct {
	ProcessingStatus       string `json:"processingStatus"`
	CVVResponseCode        string `json:"cvvResponseCode"`
	AuthorizationCode      string `json:"authorizationCode"`
	AVSResponseCodeZip     string `json:"avsResponseCodeZip"`
	AVSResponseCodeAddress string `json:"avsResponseCodeAddress"`
	AVSResponseCodeName    string `json:"avsResponseCodeName"`
	NetworkTransactionId   string `json:"networkTransactionId"`
}

type FraudResultInfo struct {
	DeviceDataCollector string `json:"deviceDataCollector"`
}
