/*
  Warnings:

  - You are about to drop the column `active` on the `accounts` table. All the data in the column will be lost.
  - You are about to drop the column `deletedAt` on the `accounts` table. All the data in the column will be lost.
  - You are about to drop the `profiles` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `accountType` to the `accounts` table without a default value. This is not possible if the table is not empty.

*/
-- CreateEnum
CREATE TYPE "AccountType" AS ENUM ('Publisher', 'Advertiser');

-- AlterTable
ALTER TABLE "accounts" DROP COLUMN "active",
DROP COLUMN "deletedAt",
ADD COLUMN     "accountType" "AccountType" NOT NULL,
ADD COLUMN     "birthDate" TEXT,
ADD COLUMN     "fullName" TEXT;

-- DropTable
DROP TABLE "profiles";
