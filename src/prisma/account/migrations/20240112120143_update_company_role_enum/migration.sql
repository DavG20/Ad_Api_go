/*
  Warnings:

  - Made the column `role` on table `companyMembers` required. This step will fail if there are existing NULL values in that column.

*/
-- AlterTable
ALTER TABLE "companyMembers" ALTER COLUMN "role" SET NOT NULL;
