-- DropForeignKey
ALTER TABLE "advertisementCPMs" DROP CONSTRAINT "advertisementCPMs_advertisementId_fkey";

-- CreateTable
CREATE TABLE "advertisementResults" (
    "advertisementId" TEXT NOT NULL,
    "startedAt" TIMESTAMP(3) NOT NULL,
    "budget" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "providerBudgetShare" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "channelBudgetShare" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "totalHour" INTEGER NOT NULL DEFAULT 0,
    "totalClick" INTEGER NOT NULL DEFAULT 0,
    "totalViews" INTEGER NOT NULL DEFAULT 0,
    "totalForward" INTEGER NOT NULL DEFAULT 0,
    "totalReaction" INTEGER NOT NULL DEFAULT 0,

    CONSTRAINT "advertisementResults_pkey" PRIMARY KEY ("advertisementId")
);

-- CreateIndex
CREATE UNIQUE INDEX "advertisementResults_advertisementId_key" ON "advertisementResults"("advertisementId");

-- AddForeignKey
ALTER TABLE "advertisementCPMs" ADD CONSTRAINT "advertisementCPMs_advertisementId_fkey" FOREIGN KEY ("advertisementId") REFERENCES "advertisements"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "advertisementResults" ADD CONSTRAINT "advertisementResults_advertisementId_fkey" FOREIGN KEY ("advertisementId") REFERENCES "advertisements"("id") ON DELETE CASCADE ON UPDATE CASCADE;
