package service

type FilterData struct {
	After  *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}
type MinMaxFilterInput struct {
	Min float64  `json:"min"`
	Max *float64 `json:"max,omitempty"`
}
type ObjectiveType string

const (
	Reach      ObjectiveType = "Reach"
	Impression ObjectiveType = "Impression"
	Clicks     ObjectiveType = "Clicks"
)

type ContentType string

const (
	ImageMediaGroup ContentType = "ImageMediaGroup"
	JustText        ContentType = "JustText"
)

type AdStatus string

const (
	AdStatusPending  AdStatus = "Pending"
	AdStatusRunning  AdStatus = "Running"
	AdStatusClosed   AdStatus = "Closed"
	AdStatusRejected AdStatus = "Rejected"
)

type AccountType string

const (
	AccountTypePublisher  AccountType = "Publisher"
	AccountTypeAdvertiser AccountType = "Advertiser"
)

type CompanyRole string

const (
	CompanyRoleAdmin  CompanyRole = "Admin"
	CompanyRoleMember CompanyRole = "Member"
)

type FundingStatus string

const (
	StatusProcessing FundingStatus = "Processing"
	StatusSuccessful FundingStatus = "Successful"
	StatusFailed     FundingStatus = "Failed"
)
