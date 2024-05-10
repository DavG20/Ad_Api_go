-- CreateTable
CREATE TABLE "channelLifeTimeBalances" (
    "id" TEXT NOT NULL,
    "accountId" TEXT NOT NULL,
    "channelId" TEXT NOT NULL,
    "totalAmount" DOUBLE PRECISION NOT NULL,
    "currency" TEXT NOT NULL DEFAULT 'ETB',
    "totalHour" DOUBLE PRECISION NOT NULL,
    "totalAd" DOUBLE PRECISION NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "channelLifeTimeBalances_pkey" PRIMARY KEY ("id")
);
