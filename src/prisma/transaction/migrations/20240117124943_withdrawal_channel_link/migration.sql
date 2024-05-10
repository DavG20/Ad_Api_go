-- CreateTable
CREATE TABLE "withdrawalChannelLinks" (
    "withdrawalId" TEXT NOT NULL,
    "channelBalanceId" TEXT NOT NULL,

    CONSTRAINT "withdrawalChannelLinks_pkey" PRIMARY KEY ("withdrawalId","channelBalanceId")
);

-- CreateIndex
CREATE UNIQUE INDEX "withdrawalChannelLinks_withdrawalId_key" ON "withdrawalChannelLinks"("withdrawalId");

-- AddForeignKey
ALTER TABLE "withdrawalChannelLinks" ADD CONSTRAINT "withdrawalChannelLinks_withdrawalId_fkey" FOREIGN KEY ("withdrawalId") REFERENCES "withdrawals"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "withdrawalChannelLinks" ADD CONSTRAINT "withdrawalChannelLinks_channelBalanceId_fkey" FOREIGN KEY ("channelBalanceId") REFERENCES "channelBalances"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
