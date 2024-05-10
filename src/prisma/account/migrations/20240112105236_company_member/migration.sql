-- CreateEnum
CREATE TYPE "CompanyRole" AS ENUM ('ADMIN', 'MEMBER');

-- CreateTable
CREATE TABLE "companyMembers" (
    "id" TEXT NOT NULL,
    "accountId" TEXT NOT NULL,
    "companyId" TEXT NOT NULL,
    "role" "CompanyRole" DEFAULT 'ADMIN',
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "companyMembers_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "companyMembers_accountId_companyId_key" ON "companyMembers"("accountId", "companyId");

-- AddForeignKey
ALTER TABLE "companyMembers" ADD CONSTRAINT "companyMembers_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "companies"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "companyMembers" ADD CONSTRAINT "companyMembers_accountId_fkey" FOREIGN KEY ("accountId") REFERENCES "accounts"("id") ON DELETE CASCADE ON UPDATE CASCADE;
