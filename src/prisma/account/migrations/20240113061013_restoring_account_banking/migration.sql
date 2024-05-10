-- CreateTable
CREATE TABLE "accountBankings" (
    "id" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "bankId" TEXT NOT NULL,
    "fullNameOnBank" TEXT NOT NULL,
    "bankAccount" TEXT NOT NULL,
    "currency" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "accountBankings_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "accountBankings_bankAccount_bankId_key" ON "accountBankings"("bankAccount", "bankId");

-- AddForeignKey
ALTER TABLE "accountBankings" ADD CONSTRAINT "accountBankings_bankId_fkey" FOREIGN KEY ("bankId") REFERENCES "banks"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
