// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateMercariCrawlSettingExcludeKeywordResult interface {
	IsCreateMercariCrawlSettingExcludeKeywordResult()
}

type CreateMercariCrawlSettingExcludeKeywordResultErrors interface {
	IsCreateMercariCrawlSettingExcludeKeywordResultErrors()
}

type CreateProductResult interface {
	IsCreateProductResult()
}

type CreateProductResultErrors interface {
	IsCreateProductResultErrors()
}

type CreateYahooAuctionCrawlSettingExcludeKeywordResult interface {
	IsCreateYahooAuctionCrawlSettingExcludeKeywordResult()
}

type CreateYahooAuctionCrawlSettingExcludeKeywordResultErrors interface {
	IsCreateYahooAuctionCrawlSettingExcludeKeywordResultErrors()
}

type DeleteMercariCrawlSettingExcludeKeywordResult interface {
	IsDeleteMercariCrawlSettingExcludeKeywordResult()
}

type DeleteMercariCrawlSettingExcludeKeywordResultErrors interface {
	IsDeleteMercariCrawlSettingExcludeKeywordResultErrors()
}

type DeleteProductResult interface {
	IsDeleteProductResult()
}

type DeleteProductResultErrors interface {
	IsDeleteProductResultErrors()
}

type DeleteYahooAuctionCrawlSettingExcludeKeywordResult interface {
	IsDeleteYahooAuctionCrawlSettingExcludeKeywordResult()
}

type DeleteYahooAuctionCrawlSettingExcludeKeywordResultErrors interface {
	IsDeleteYahooAuctionCrawlSettingExcludeKeywordResultErrors()
}

type Node interface {
	IsNode()
	GetID() string
}

type ResultBase interface {
	IsResultBase()
	GetOk() bool
}

type UpdateMercariCrawlSettingExcludeKeywordResult interface {
	IsUpdateMercariCrawlSettingExcludeKeywordResult()
}

type UpdateMercariCrawlSettingExcludeKeywordResultErrors interface {
	IsUpdateMercariCrawlSettingExcludeKeywordResultErrors()
}

type UpdateProductResult interface {
	IsUpdateProductResult()
}

type UpdateProductResultErrors interface {
	IsUpdateProductResultErrors()
}

type UpdateYahooAuctionCrawlSettingExcludeKeywordResult interface {
	IsUpdateYahooAuctionCrawlSettingExcludeKeywordResult()
}

type UpdateYahooAuctionCrawlSettingExcludeKeywordResultErrors interface {
	IsUpdateYahooAuctionCrawlSettingExcludeKeywordResultErrors()
}

type UserError interface {
	IsUserError()
	GetCode() string
	GetMessage() string
}

type CreateMercariCrawlSettingExcludeKeywordInput struct {
	ProductID string  `json:"productId"`
	Keyword   *string `json:"keyword,omitempty"`
}

type CreateMercariCrawlSettingExcludeKeywordResultError struct {
	Ok    bool                                                `json:"ok"`
	Error CreateMercariCrawlSettingExcludeKeywordResultErrors `json:"error"`
}

func (CreateMercariCrawlSettingExcludeKeywordResultError) IsCreateMercariCrawlSettingExcludeKeywordResult() {
}

func (CreateMercariCrawlSettingExcludeKeywordResultError) IsResultBase()    {}
func (this CreateMercariCrawlSettingExcludeKeywordResultError) GetOk() bool { return this.Ok }

type CreateMercariCrawlSettingExcludeKeywordResultSuccess struct {
	Ok                                bool                               `json:"ok"`
	MercariCrawlSettingExcludeKeyword *MercariCrawlSettingExcludeKeyword `json:"mercariCrawlSettingExcludeKeyword"`
}

func (CreateMercariCrawlSettingExcludeKeywordResultSuccess) IsCreateMercariCrawlSettingExcludeKeywordResult() {
}

func (CreateMercariCrawlSettingExcludeKeywordResultSuccess) IsResultBase()    {}
func (this CreateMercariCrawlSettingExcludeKeywordResultSuccess) GetOk() bool { return this.Ok }

type CreateMercariCrawlSettingExcludeKeywordResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (CreateMercariCrawlSettingExcludeKeywordResultValidationFailed) IsCreateMercariCrawlSettingExcludeKeywordResultErrors() {
}

func (CreateMercariCrawlSettingExcludeKeywordResultValidationFailed) IsUserError() {}
func (this CreateMercariCrawlSettingExcludeKeywordResultValidationFailed) GetCode() string {
	return this.Code
}
func (this CreateMercariCrawlSettingExcludeKeywordResultValidationFailed) GetMessage() string {
	return this.Message
}

type CreateMercariCrawlSettingInput struct {
	Keyword    string `json:"keyword"`
	CategoryID *int   `json:"category_id,omitempty"`
	MinPrice   int    `json:"min_price"`
	MaxPrice   int    `json:"max_price"`
	Enabled    bool   `json:"enabled"`
}

type CreateProductInput struct {
	Name                     string                               `json:"name"`
	YahooAuctionCrawlSetting *CreateYahooAuctionCrawlSettingInput `json:"yahoo_auction_crawl_setting"`
	MercariCrawlSetting      *CreateMercariCrawlSettingInput      `json:"mercari_crawl_setting"`
}

type CreateProductResultError struct {
	Ok    bool                      `json:"ok"`
	Error CreateProductResultErrors `json:"error"`
}

func (CreateProductResultError) IsCreateProductResult() {}

func (CreateProductResultError) IsResultBase()    {}
func (this CreateProductResultError) GetOk() bool { return this.Ok }

type CreateProductResultSuccess struct {
	Ok      bool     `json:"ok"`
	Product *Product `json:"product"`
}

func (CreateProductResultSuccess) IsCreateProductResult() {}

func (CreateProductResultSuccess) IsResultBase()    {}
func (this CreateProductResultSuccess) GetOk() bool { return this.Ok }

type CreateProductResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (CreateProductResultValidationFailed) IsCreateProductResultErrors() {}

func (CreateProductResultValidationFailed) IsUserError()            {}
func (this CreateProductResultValidationFailed) GetCode() string    { return this.Code }
func (this CreateProductResultValidationFailed) GetMessage() string { return this.Message }

type CreateYahooAuctionCrawlSettingExcludeKeywordInput struct {
	ProductID string  `json:"productId"`
	Keyword   *string `json:"keyword,omitempty"`
}

type CreateYahooAuctionCrawlSettingExcludeKeywordResultError struct {
	Ok    bool                                                     `json:"ok"`
	Error CreateYahooAuctionCrawlSettingExcludeKeywordResultErrors `json:"error"`
}

func (CreateYahooAuctionCrawlSettingExcludeKeywordResultError) IsCreateYahooAuctionCrawlSettingExcludeKeywordResult() {
}

func (CreateYahooAuctionCrawlSettingExcludeKeywordResultError) IsResultBase()    {}
func (this CreateYahooAuctionCrawlSettingExcludeKeywordResultError) GetOk() bool { return this.Ok }

type CreateYahooAuctionCrawlSettingExcludeKeywordResultSuccess struct {
	Ok                                     bool                                    `json:"ok"`
	YahooAuctionCrawlSettingExcludeKeyword *YahooAuctionCrawlSettingExcludeKeyword `json:"yahooAuctionCrawlSettingExcludeKeyword"`
}

func (CreateYahooAuctionCrawlSettingExcludeKeywordResultSuccess) IsCreateYahooAuctionCrawlSettingExcludeKeywordResult() {
}

func (CreateYahooAuctionCrawlSettingExcludeKeywordResultSuccess) IsResultBase()    {}
func (this CreateYahooAuctionCrawlSettingExcludeKeywordResultSuccess) GetOk() bool { return this.Ok }

type CreateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (CreateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) IsCreateYahooAuctionCrawlSettingExcludeKeywordResultErrors() {
}

func (CreateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) IsUserError() {}
func (this CreateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) GetCode() string {
	return this.Code
}
func (this CreateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) GetMessage() string {
	return this.Message
}

type CreateYahooAuctionCrawlSettingInput struct {
	Keyword    string `json:"keyword"`
	CategoryID *int   `json:"category_id,omitempty"`
	MinPrice   int    `json:"min_price"`
	MaxPrice   int    `json:"max_price"`
	Enabled    bool   `json:"enabled"`
}

type DeleteMercariCrawlSettingExcludeKeywordResultError struct {
	Ok    bool                                                `json:"ok"`
	Error DeleteMercariCrawlSettingExcludeKeywordResultErrors `json:"error"`
}

func (DeleteMercariCrawlSettingExcludeKeywordResultError) IsDeleteMercariCrawlSettingExcludeKeywordResult() {
}

func (DeleteMercariCrawlSettingExcludeKeywordResultError) IsResultBase()    {}
func (this DeleteMercariCrawlSettingExcludeKeywordResultError) GetOk() bool { return this.Ok }

type DeleteMercariCrawlSettingExcludeKeywordResultSuccess struct {
	Ok bool `json:"ok"`
}

func (DeleteMercariCrawlSettingExcludeKeywordResultSuccess) IsDeleteMercariCrawlSettingExcludeKeywordResult() {
}

func (DeleteMercariCrawlSettingExcludeKeywordResultSuccess) IsResultBase()    {}
func (this DeleteMercariCrawlSettingExcludeKeywordResultSuccess) GetOk() bool { return this.Ok }

type DeleteMercariCrawlSettingExcludeKeywordResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (DeleteMercariCrawlSettingExcludeKeywordResultValidationFailed) IsDeleteMercariCrawlSettingExcludeKeywordResultErrors() {
}

func (DeleteMercariCrawlSettingExcludeKeywordResultValidationFailed) IsUserError() {}
func (this DeleteMercariCrawlSettingExcludeKeywordResultValidationFailed) GetCode() string {
	return this.Code
}
func (this DeleteMercariCrawlSettingExcludeKeywordResultValidationFailed) GetMessage() string {
	return this.Message
}

type DeleteProductResultError struct {
	Ok    bool                      `json:"ok"`
	Error DeleteProductResultErrors `json:"error"`
}

func (DeleteProductResultError) IsDeleteProductResult() {}

func (DeleteProductResultError) IsResultBase()    {}
func (this DeleteProductResultError) GetOk() bool { return this.Ok }

type DeleteProductResultSuccess struct {
	Ok bool `json:"ok"`
}

func (DeleteProductResultSuccess) IsDeleteProductResult() {}

func (DeleteProductResultSuccess) IsResultBase()    {}
func (this DeleteProductResultSuccess) GetOk() bool { return this.Ok }

type DeleteProductResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (DeleteProductResultValidationFailed) IsDeleteProductResultErrors() {}

func (DeleteProductResultValidationFailed) IsUserError()            {}
func (this DeleteProductResultValidationFailed) GetCode() string    { return this.Code }
func (this DeleteProductResultValidationFailed) GetMessage() string { return this.Message }

type DeleteYahooAuctionCrawlSettingExcludeKeywordResultError struct {
	Ok    bool                                                     `json:"ok"`
	Error DeleteYahooAuctionCrawlSettingExcludeKeywordResultErrors `json:"error"`
}

func (DeleteYahooAuctionCrawlSettingExcludeKeywordResultError) IsDeleteYahooAuctionCrawlSettingExcludeKeywordResult() {
}

func (DeleteYahooAuctionCrawlSettingExcludeKeywordResultError) IsResultBase()    {}
func (this DeleteYahooAuctionCrawlSettingExcludeKeywordResultError) GetOk() bool { return this.Ok }

type DeleteYahooAuctionCrawlSettingExcludeKeywordResultSuccess struct {
	Ok bool `json:"ok"`
}

func (DeleteYahooAuctionCrawlSettingExcludeKeywordResultSuccess) IsDeleteYahooAuctionCrawlSettingExcludeKeywordResult() {
}

func (DeleteYahooAuctionCrawlSettingExcludeKeywordResultSuccess) IsResultBase()    {}
func (this DeleteYahooAuctionCrawlSettingExcludeKeywordResultSuccess) GetOk() bool { return this.Ok }

type DeleteYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (DeleteYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) IsDeleteYahooAuctionCrawlSettingExcludeKeywordResultErrors() {
}

func (DeleteYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) IsUserError() {}
func (this DeleteYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) GetCode() string {
	return this.Code
}
func (this DeleteYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) GetMessage() string {
	return this.Message
}

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type MercariCrawlSetting struct {
	ID                                 string                               `json:"id"`
	ProductID                          int                                  `json:"productId"`
	Keyword                            string                               `json:"keyword"`
	MinPrice                           int                                  `json:"minPrice"`
	MaxPrice                           int                                  `json:"maxPrice"`
	CategoryID                         *int                                 `json:"categoryId,omitempty"`
	Enabled                            bool                                 `json:"enabled"`
	CreatedAt                          string                               `json:"createdAt"`
	UpdatedAt                          string                               `json:"updatedAt"`
	MercariCrawlSettingExcludeKeywords []*MercariCrawlSettingExcludeKeyword `json:"mercariCrawlSettingExcludeKeywords"`
}

func (MercariCrawlSetting) IsNode()            {}
func (this MercariCrawlSetting) GetID() string { return this.ID }

type MercariCrawlSettingExcludeKeyword struct {
	ID                    string  `json:"id"`
	MercariCrawlSettingID int     `json:"mercariCrawlSettingId"`
	Keyword               *string `json:"keyword,omitempty"`
	CreatedAt             string  `json:"createdAt"`
	UpdatedAt             string  `json:"updatedAt"`
}

func (MercariCrawlSettingExcludeKeyword) IsNode()            {}
func (this MercariCrawlSettingExcludeKeyword) GetID() string { return this.ID }

type MercariProduct struct {
	ID           string  `json:"id"`
	ProductID    int     `json:"productId"`
	MercariID    string  `json:"mercariId"`
	Name         string  `json:"name"`
	ThumbnailURL string  `json:"thumbnailUrl"`
	Price        int     `json:"price"`
	Published    bool    `json:"published"`
	BoughtDate   *string `json:"boughtDate,omitempty"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}

func (MercariProduct) IsNode()            {}
func (this MercariProduct) GetID() string { return this.ID }

type Product struct {
	ID                                 string                              `json:"id"`
	Name                               string                              `json:"name"`
	YahooAuctionProducts               []*YahooAuctionProduct              `json:"yahooAuctionProducts"`
	YahooAuctionCrawlSetting           *YahooAuctionCrawlSetting           `json:"yahooAuctionCrawlSetting"`
	YahooAuctionDailyPurchaseSummaries []*YahooAuctionDailyPurchaseSummary `json:"yahooAuctionDailyPurchaseSummaries"`
	MercariProducts                    []*MercariProduct                   `json:"mercariProducts"`
	MercariCrawlSetting                *MercariCrawlSetting                `json:"mercariCrawlSetting"`
}

func (Product) IsNode()            {}
func (this Product) GetID() string { return this.ID }

type UpdateMercariCrawlSettingExcludeKeywordInput struct {
	ID        string  `json:"id"`
	ProductID string  `json:"productId"`
	Keyword   *string `json:"keyword,omitempty"`
}

type UpdateMercariCrawlSettingExcludeKeywordResultError struct {
	Ok    bool                                                `json:"ok"`
	Error UpdateMercariCrawlSettingExcludeKeywordResultErrors `json:"error"`
}

func (UpdateMercariCrawlSettingExcludeKeywordResultError) IsUpdateMercariCrawlSettingExcludeKeywordResult() {
}

func (UpdateMercariCrawlSettingExcludeKeywordResultError) IsResultBase()    {}
func (this UpdateMercariCrawlSettingExcludeKeywordResultError) GetOk() bool { return this.Ok }

type UpdateMercariCrawlSettingExcludeKeywordResultSuccess struct {
	Ok                                bool                               `json:"ok"`
	MercariCrawlSettingExcludeKeyword *MercariCrawlSettingExcludeKeyword `json:"mercariCrawlSettingExcludeKeyword"`
}

func (UpdateMercariCrawlSettingExcludeKeywordResultSuccess) IsUpdateMercariCrawlSettingExcludeKeywordResult() {
}

func (UpdateMercariCrawlSettingExcludeKeywordResultSuccess) IsResultBase()    {}
func (this UpdateMercariCrawlSettingExcludeKeywordResultSuccess) GetOk() bool { return this.Ok }

type UpdateMercariCrawlSettingExcludeKeywordResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (UpdateMercariCrawlSettingExcludeKeywordResultValidationFailed) IsUpdateMercariCrawlSettingExcludeKeywordResultErrors() {
}

func (UpdateMercariCrawlSettingExcludeKeywordResultValidationFailed) IsUserError() {}
func (this UpdateMercariCrawlSettingExcludeKeywordResultValidationFailed) GetCode() string {
	return this.Code
}
func (this UpdateMercariCrawlSettingExcludeKeywordResultValidationFailed) GetMessage() string {
	return this.Message
}

type UpdateMercariCrawlSettingInput struct {
	Keyword    string `json:"keyword"`
	CategoryID *int   `json:"category_id,omitempty"`
	MinPrice   int    `json:"min_price"`
	MaxPrice   int    `json:"max_price"`
	Enabled    bool   `json:"enabled"`
}

type UpdateProductInput struct {
	Name                     string                               `json:"name"`
	YahooAuctionCrawlSetting *UpdateYahooAuctionCrawlSettingInput `json:"yahoo_auction_crawl_setting"`
	MercariCrawlSetting      *UpdateMercariCrawlSettingInput      `json:"mercari_crawl_setting"`
}

type UpdateProductResultError struct {
	Ok    bool                      `json:"ok"`
	Error UpdateProductResultErrors `json:"error"`
}

func (UpdateProductResultError) IsUpdateProductResult() {}

func (UpdateProductResultError) IsResultBase()    {}
func (this UpdateProductResultError) GetOk() bool { return this.Ok }

type UpdateProductResultSuccess struct {
	Ok      bool     `json:"ok"`
	Product *Product `json:"product"`
}

func (UpdateProductResultSuccess) IsUpdateProductResult() {}

func (UpdateProductResultSuccess) IsResultBase()    {}
func (this UpdateProductResultSuccess) GetOk() bool { return this.Ok }

type UpdateProductResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (UpdateProductResultValidationFailed) IsUpdateProductResultErrors() {}

func (UpdateProductResultValidationFailed) IsUserError()            {}
func (this UpdateProductResultValidationFailed) GetCode() string    { return this.Code }
func (this UpdateProductResultValidationFailed) GetMessage() string { return this.Message }

type UpdateYahooAuctionCrawlSettingExcludeKeywordInput struct {
	ID        string  `json:"id"`
	ProductID string  `json:"productId"`
	Keyword   *string `json:"keyword,omitempty"`
}

type UpdateYahooAuctionCrawlSettingExcludeKeywordResultError struct {
	Ok    bool                                                     `json:"ok"`
	Error UpdateYahooAuctionCrawlSettingExcludeKeywordResultErrors `json:"error"`
}

func (UpdateYahooAuctionCrawlSettingExcludeKeywordResultError) IsUpdateYahooAuctionCrawlSettingExcludeKeywordResult() {
}

func (UpdateYahooAuctionCrawlSettingExcludeKeywordResultError) IsResultBase()    {}
func (this UpdateYahooAuctionCrawlSettingExcludeKeywordResultError) GetOk() bool { return this.Ok }

type UpdateYahooAuctionCrawlSettingExcludeKeywordResultSuccess struct {
	Ok                                     bool                                    `json:"ok"`
	YahooAuctionCrawlSettingExcludeKeyword *YahooAuctionCrawlSettingExcludeKeyword `json:"yahooAuctionCrawlSettingExcludeKeyword"`
}

func (UpdateYahooAuctionCrawlSettingExcludeKeywordResultSuccess) IsUpdateYahooAuctionCrawlSettingExcludeKeywordResult() {
}

func (UpdateYahooAuctionCrawlSettingExcludeKeywordResultSuccess) IsResultBase()    {}
func (this UpdateYahooAuctionCrawlSettingExcludeKeywordResultSuccess) GetOk() bool { return this.Ok }

type UpdateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (UpdateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) IsUpdateYahooAuctionCrawlSettingExcludeKeywordResultErrors() {
}

func (UpdateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) IsUserError() {}
func (this UpdateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) GetCode() string {
	return this.Code
}
func (this UpdateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed) GetMessage() string {
	return this.Message
}

type UpdateYahooAuctionCrawlSettingInput struct {
	Keyword    string `json:"keyword"`
	CategoryID *int   `json:"category_id,omitempty"`
	MinPrice   int    `json:"min_price"`
	MaxPrice   int    `json:"max_price"`
	Enabled    bool   `json:"enabled"`
}

type YahooAuctionCrawlSetting struct {
	ID                                      string                                    `json:"id"`
	ProductID                               int                                       `json:"productId"`
	Keyword                                 string                                    `json:"keyword"`
	MinPrice                                int                                       `json:"minPrice"`
	MaxPrice                                int                                       `json:"maxPrice"`
	CategoryID                              *int                                      `json:"categoryId,omitempty"`
	Enabled                                 bool                                      `json:"enabled"`
	CreatedAt                               string                                    `json:"createdAt"`
	UpdatedAt                               string                                    `json:"updatedAt"`
	YahooAuctionCrawlSettingExcludeKeywords []*YahooAuctionCrawlSettingExcludeKeyword `json:"yahooAuctionCrawlSettingExcludeKeywords"`
}

func (YahooAuctionCrawlSetting) IsNode()            {}
func (this YahooAuctionCrawlSetting) GetID() string { return this.ID }

type YahooAuctionCrawlSettingExcludeKeyword struct {
	ID                         string  `json:"id"`
	YahooAuctionCrawlSettingID int     `json:"yahooAuctionCrawlSettingId"`
	Keyword                    *string `json:"keyword,omitempty"`
	CreatedAt                  string  `json:"createdAt"`
	UpdatedAt                  string  `json:"updatedAt"`
}

func (YahooAuctionCrawlSettingExcludeKeyword) IsNode()            {}
func (this YahooAuctionCrawlSettingExcludeKeyword) GetID() string { return this.ID }

type YahooAuctionDailyPurchaseSummary struct {
	ID                   string `json:"id"`
	ProductID            int    `json:"productId"`
	AveragePurchasePrice *int   `json:"averagePurchasePrice,omitempty"`
	PurchaseCount        int    `json:"purchaseCount"`
	Date                 string `json:"date"`
	CreatedAt            string `json:"createdAt"`
	UpdatedAt            string `json:"updatedAt"`
}

func (YahooAuctionDailyPurchaseSummary) IsNode()            {}
func (this YahooAuctionDailyPurchaseSummary) GetID() string { return this.ID }

type YahooAuctionProduct struct {
	ID             string  `json:"id"`
	ProductID      int     `json:"productId"`
	YahooAuctionID string  `json:"yahooAuctionId"`
	Name           string  `json:"name"`
	ThumbnailURL   string  `json:"thumbnailUrl"`
	Price          int     `json:"price"`
	Published      bool    `json:"published"`
	BoughtDate     *string `json:"boughtDate,omitempty"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
}

func (YahooAuctionProduct) IsNode()            {}
func (this YahooAuctionProduct) GetID() string { return this.ID }
