-- CreateTable
CREATE TABLE "channelCollectedAdLinks" (
    "channelBalanceId" TEXT NOT NULL,
    "adPaymentId" TEXT NOT NULL,

    CONSTRAINT "channelCollectedAdLinks_pkey" PRIMARY KEY ("channelBalanceId","adPaymentId")
);

-- AddForeignKey
ALTER TABLE "channelCollectedAdLinks" ADD CONSTRAINT "channelCollectedAdLinks_adPaymentId_fkey" FOREIGN KEY ("adPaymentId") REFERENCES "adPayments"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "channelCollectedAdLinks" ADD CONSTRAINT "channelCollectedAdLinks_channelBalanceId_fkey" FOREIGN KEY ("channelBalanceId") REFERENCES "channelBalances"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
