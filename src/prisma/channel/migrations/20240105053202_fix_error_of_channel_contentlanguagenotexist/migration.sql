-- DropForeignKey
ALTER TABLE "channelCategories" DROP CONSTRAINT "channelCategories_channelId_fkey";

-- DropForeignKey
ALTER TABLE "channelDetails" DROP CONSTRAINT "channelDetails_channelId_fkey";

-- DropForeignKey
ALTER TABLE "cpmRates" DROP CONSTRAINT "cpmRates_channelId_fkey";

-- DropForeignKey
ALTER TABLE "hourlyRates" DROP CONSTRAINT "hourlyRates_channelId_fkey";

-- AddForeignKey
ALTER TABLE "channelCategories" ADD CONSTRAINT "channelCategories_channelId_fkey" FOREIGN KEY ("channelId") REFERENCES "channels"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "channelContentLanguages" ADD CONSTRAINT "channelContentLanguages_channelId_fkey" FOREIGN KEY ("channelId") REFERENCES "channels"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "hourlyRates" ADD CONSTRAINT "hourlyRates_channelId_fkey" FOREIGN KEY ("channelId") REFERENCES "channels"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "cpmRates" ADD CONSTRAINT "cpmRates_channelId_fkey" FOREIGN KEY ("channelId") REFERENCES "channels"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "channelDetails" ADD CONSTRAINT "channelDetails_channelId_fkey" FOREIGN KEY ("channelId") REFERENCES "channels"("id") ON DELETE CASCADE ON UPDATE CASCADE;
