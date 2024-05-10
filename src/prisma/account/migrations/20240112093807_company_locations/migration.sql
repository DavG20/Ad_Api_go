/*
  Warnings:

  - You are about to drop the `Companies` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "accountBankings" DROP CONSTRAINT "accountBankings_bankId_fkey";

-- DropTable
DROP TABLE "Companies";

-- CreateTable
CREATE TABLE "companies" (
    "id" TEXT NOT NULL,
    "creatorId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "uniqueName" TEXT NOT NULL,
    "tinNumber" TEXT,
    "vatNumber" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "companies_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "companyLocations" (
    "companyId" TEXT NOT NULL,
    "locationId" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "companyLocations_pkey" PRIMARY KEY ("companyId","locationId")
);

-- CreateIndex
CREATE UNIQUE INDEX "companies_uniqueName_key" ON "companies"("uniqueName");

-- AddForeignKey
ALTER TABLE "accountBankings" ADD CONSTRAINT "accountBankings_bankId_fkey" FOREIGN KEY ("bankId") REFERENCES "banks"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "companyLocations" ADD CONSTRAINT "companyLocations_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "companies"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "companyLocations" ADD CONSTRAINT "companyLocations_locationId_fkey" FOREIGN KEY ("locationId") REFERENCES "locations"("id") ON DELETE CASCADE ON UPDATE CASCADE;
