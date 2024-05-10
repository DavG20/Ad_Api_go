/*
  Warnings:

  - A unique constraint covering the columns `[channelId]` on the table `channelDetails` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "channelDetails_channelId_key" ON "channelDetails"("channelId");

-- AddForeignKey
ALTER TABLE "channelCategories" ADD CONSTRAINT "channelCategories_channelId_fkey" FOREIGN KEY ("channelId") REFERENCES "channels"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "hourlyRates" ADD CONSTRAINT "hourlyRates_channelId_fkey" FOREIGN KEY ("channelId") REFERENCES "channels"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "cpmRates" ADD CONSTRAINT "cpmRates_channelId_fkey" FOREIGN KEY ("channelId") REFERENCES "channels"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "channelDetails" ADD CONSTRAINT "channelDetails_channelId_fkey" FOREIGN KEY ("channelId") REFERENCES "channels"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
