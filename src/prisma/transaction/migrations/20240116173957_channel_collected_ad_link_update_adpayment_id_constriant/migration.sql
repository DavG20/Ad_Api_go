/*
  Warnings:

  - A unique constraint covering the columns `[adPaymentId]` on the table `channelCollectedAdLinks` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "channelCollectedAdLinks_adPaymentId_key" ON "channelCollectedAdLinks"("adPaymentId");

-- AddForeignKey
ALTER TABLE "channelCollectedAdLinks" ADD CONSTRAINT "channelCollectedAdLinks_adPaymentId_fkey" FOREIGN KEY ("adPaymentId") REFERENCES "adPayments"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
