/*
  Warnings:

  - A unique constraint covering the columns `[channelId]` on the table `channelLifeTimeBalances` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "channelLifeTimeBalances_channelId_key" ON "channelLifeTimeBalances"("channelId");
