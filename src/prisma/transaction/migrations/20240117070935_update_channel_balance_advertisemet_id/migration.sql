/*
  Warnings:

  - A unique constraint covering the columns `[advertisementId]` on the table `adPayments` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[channelId]` on the table `channelBalances` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "adPayments_advertisementId_key" ON "adPayments"("advertisementId");

-- CreateIndex
CREATE UNIQUE INDEX "channelBalances_channelId_key" ON "channelBalances"("channelId");
