/*
  Warnings:

  - You are about to drop the column `postFrequecy` on the `channelDetails` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "channelDetails" DROP COLUMN "postFrequecy",
ADD COLUMN     "PostFrequency" DOUBLE PRECISION NOT NULL DEFAULT 0;
