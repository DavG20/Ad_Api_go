-- CreateTable
CREATE TABLE "withdrawals" (
    "id" TEXT NOT NULL,
    "accountId" TEXT NOT NULL,
    "accountBankId" TEXT NOT NULL,
    "totalAmount" DOUBLE PRECISION NOT NULL,
    "currency" TEXT NOT NULL DEFAULT 'ETB',
    "fundingKey" TEXT,
    "method" TEXT,
    "reference" TEXT,
    "paymentDate" TIMESTAMP(3),
    "status" "FundingStatus" NOT NULL DEFAULT 'Processing',
    "extra" BYTEA,
    "log" BYTEA,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "withdrawals_pkey" PRIMARY KEY ("id")
);
