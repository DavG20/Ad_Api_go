-- CreateTable
CREATE TABLE "channelBalances" (
    "id" TEXT NOT NULL,
    "accountId" TEXT NOT NULL,
    "channelId" TEXT NOT NULL,
    "amount" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "currency" TEXT NOT NULL DEFAULT 'ETB',
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "channelBalances_pkey" PRIMARY KEY ("id")
);
