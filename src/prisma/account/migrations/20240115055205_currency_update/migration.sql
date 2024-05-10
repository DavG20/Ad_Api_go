/*
  Warnings:

  - Made the column `currency` on table `accountBankings` required. This step will fail if there are existing NULL values in that column.

*/
-- AlterTable
ALTER TABLE "accountBankings" ALTER COLUMN "currency" SET NOT NULL;
