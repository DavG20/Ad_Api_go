/*
  Warnings:

  - The primary key for the `channelDetails` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `channelDetails` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "channelDetails" DROP CONSTRAINT "channelDetails_pkey",
DROP COLUMN "id",
ALTER COLUMN "deletedAt" DROP NOT NULL,
ADD CONSTRAINT "channelDetails_pkey" PRIMARY KEY ("channelId");
