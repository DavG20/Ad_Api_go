/*
  Warnings:

  - You are about to drop the column `createdAt` on the `companyLocations` table. All the data in the column will be lost.
  - You are about to drop the column `deletedAt` on the `companyLocations` table. All the data in the column will be lost.
  - You are about to drop the column `updatedAt` on the `companyLocations` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "companies" ALTER COLUMN "updatedAt" SET DEFAULT CURRENT_TIMESTAMP;

-- AlterTable
ALTER TABLE "companyLocations" DROP COLUMN "createdAt",
DROP COLUMN "deletedAt",
DROP COLUMN "updatedAt";
