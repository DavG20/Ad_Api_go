package enumConverter

import (
	accountDB "adtec/backend/src/prisma/account/db"
	"adtec/backend/src/prisma/campaign/db"
	transactionDB "adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/service"
)

func ObjectiveTypeConverter(userData []service.ObjectiveType) []db.ObjectiveType {
	var objectiveTypes []db.ObjectiveType
	for _, object := range userData {

		objectiveTypes = append(objectiveTypes, db.ObjectiveType(object))
	}
	return objectiveTypes
}
func ContentTypeConverter(userData []service.ContentType) []db.ContentType {
	var contentTypes []db.ContentType
	for _, object := range userData {

		contentTypes = append(contentTypes, db.ContentType(object))
	}
	return contentTypes
}
func AdvertisementStatusConverter(userData []service.AdStatus) []db.AdStatus {
	var advertisementStatus []db.AdStatus
	for _, object := range userData {

		advertisementStatus = append(advertisementStatus, db.AdStatus(object))
	}
	return advertisementStatus
}

func AccountTypeConverter(userData []service.AccountType) []accountDB.AccountType {
	var accountType []accountDB.AccountType
	for _, object := range userData {

		accountType = append(accountType, accountDB.AccountType(object))
	}
	return accountType
}
func CompanyRoleConverter(userData []service.CompanyRole) []accountDB.CompanyRole {
	var companyRole []accountDB.CompanyRole
	for _, object := range userData {

		companyRole = append(companyRole, accountDB.CompanyRole(object))
	}
	return companyRole
}
func FundingStatusConverter(userData []service.FundingStatus) []transactionDB.FundingStatus {
	var fundingStatus []transactionDB.FundingStatus
	for _, object := range userData {

		fundingStatus = append(fundingStatus, transactionDB.FundingStatus(object))
	}
	return fundingStatus
}
