-- CreateTable
CREATE TABLE "cpmRates" (
    "id" TEXT NOT NULL,
    "channelId" TEXT NOT NULL,
    "active" BOOLEAN NOT NULL DEFAULT true,
    "CPM" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "minCPMVolume" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "cpmRates_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "cpmRates_channelId_key" ON "cpmRates"("channelId");
