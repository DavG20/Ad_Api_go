-- CreateEnum
CREATE TYPE "FundingStatus" AS ENUM ('Processing', 'Successful', 'Failed');

-- CreateTable
CREATE TABLE "funding" (
    "id" TEXT NOT NULL,
    "companyId" TEXT NOT NULL,
    "amount" DOUBLE PRECISION NOT NULL,
    "currency" TEXT NOT NULL DEFAULT 'ETB',
    "key" TEXT NOT NULL,
    "method" TEXT,
    "fundingTxRef" TEXT,
    "reference" TEXT,
    "paymentDate" TIMESTAMP(3),
    "redirectUrl" TEXT,
    "status" "FundingStatus" NOT NULL DEFAULT 'Processing',
    "extra" BYTEA,
    "log" BYTEA,
    "tax" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "funding_pkey" PRIMARY KEY ("id")
);
