/*
  Warnings:

  - You are about to drop the `accountBankings` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "accountBankings" DROP CONSTRAINT "accountBankings_bankId_fkey";

-- DropTable
DROP TABLE "accountBankings";

-- CreateTable
CREATE TABLE "companyBankings" (
    "id" TEXT NOT NULL,
    "companyId" TEXT NOT NULL,
    "bankId" TEXT NOT NULL,
    "fullNameOnBank" TEXT NOT NULL,
    "bankAccount" TEXT NOT NULL,
    "currency" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "companyBankings_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "companyBankings_bankAccount_bankId_key" ON "companyBankings"("bankAccount", "bankId");

-- AddForeignKey
ALTER TABLE "companyBankings" ADD CONSTRAINT "companyBankings_bankId_fkey" FOREIGN KEY ("bankId") REFERENCES "banks"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "companyBankings" ADD CONSTRAINT "companyBankings_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "companies"("id") ON DELETE CASCADE ON UPDATE CASCADE;
