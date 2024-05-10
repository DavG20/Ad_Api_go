-- CreateTable
CREATE TABLE "budgets" (
    "campaignId" TEXT NOT NULL,
    "initialAmount" DOUBLE PRECISION NOT NULL,
    "usedAmount" DOUBLE PRECISION NOT NULL,
    "refundedAmount" DOUBLE PRECISION,
    "currency" TEXT NOT NULL DEFAULT 'ETB',

    CONSTRAINT "budgets_pkey" PRIMARY KEY ("campaignId")
);

-- CreateIndex
CREATE UNIQUE INDEX "budgets_campaignId_key" ON "budgets"("campaignId");

-- AddForeignKey
ALTER TABLE "budgets" ADD CONSTRAINT "budgets_campaignId_fkey" FOREIGN KEY ("campaignId") REFERENCES "campaigns"("id") ON DELETE CASCADE ON UPDATE CASCADE;
