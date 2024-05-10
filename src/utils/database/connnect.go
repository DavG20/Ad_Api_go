package database

import (
	profileDB "adtec/backend/src/prisma/account/db"
	campaignDB "adtec/backend/src/prisma/campaign/db"
	channelDB "adtec/backend/src/prisma/channel/db"
	transactionDB "adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database/clients"
	"context"
)

var accountClient *profileDB.PrismaClient
var accountContext context.Context
var accountErr error

var channelClient *channelDB.PrismaClient
var channelContext context.Context
var channelErr error

var campaignClient *campaignDB.PrismaClient
var campaignContext context.Context
var campaignErr error

var transactionClient *transactionDB.PrismaClient
var transactionContext context.Context
var transactionErr error

func ConnectToDatabase() {
	accountClient, accountContext, accountErr = clients.AccountPrismaClient()
	channelClient, channelContext, channelErr = clients.ChannelPrismaClient()
	campaignClient, campaignContext, campaignErr = clients.CampaignPrismaClient()
	transactionClient, transactionContext, transactionErr = clients.TransactionPrismaClient()

}

func GetAccountClient() (*profileDB.PrismaClient, context.Context, error) {
	if accountErr != nil || accountClient == nil || accountContext == nil {
		if accountClient != nil {
			accountClient.Prisma.Disconnect()
		}
		accountClient, accountContext, accountErr = clients.AccountPrismaClient()
	}

	return accountClient, accountContext, accountErr
}

func GetChannelClient() (*channelDB.PrismaClient, context.Context, error) {
	if channelErr != nil || channelClient == nil || channelContext == nil {
		if channelClient != nil {
			channelClient.Prisma.Disconnect()
		}
		channelClient, channelContext, channelErr = clients.ChannelPrismaClient()
	}

	return channelClient, channelContext, channelErr
}

func GetCampaignClient() (*campaignDB.PrismaClient, context.Context, error) {
	if campaignErr != nil || campaignClient == nil || campaignContext == nil {
		if campaignClient != nil {
			campaignClient.Prisma.Disconnect()
		}
		campaignClient, campaignContext, campaignErr = clients.CampaignPrismaClient()
	}

	return campaignClient, campaignContext, campaignErr
}

func GetTransactionClient() (*transactionDB.PrismaClient, context.Context, error) {
	if transactionErr != nil || transactionClient == nil || transactionContext == nil {
		if transactionClient != nil {
			transactionClient.Prisma.Disconnect()
		}
		transactionClient, transactionContext, transactionErr = clients.TransactionPrismaClient()
	}

	return transactionClient, transactionContext, transactionErr
}
