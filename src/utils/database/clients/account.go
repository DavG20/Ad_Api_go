package clients

import (
	"adtec/backend/src/prisma/account/db"
	"context"
	"log"

	"github.com/getsentry/sentry-go"
)

func AccountPrismaClient() (*db.PrismaClient, context.Context, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Println("prisma setup error:", err)
		sentry.CaptureException(err)
		return nil, nil, err
	}

	ctx := context.Background()

	return client, ctx, nil
}
