-- CreateTable
CREATE TABLE "advertisementCPMs" (
    "advertisementId" TEXT NOT NULL,
    "Rate" DOUBLE PRECISION NOT NULL,
    "channelShare" DOUBLE PRECISION NOT NULL,
    "providerShare" DOUBLE PRECISION NOT NULL,
    "totalBudget" DOUBLE PRECISION NOT NULL,
    "maxLifeCycle" INTEGER NOT NULL,
    "requiredViews" INTEGER NOT NULL,

    CONSTRAINT "advertisementCPMs_pkey" PRIMARY KEY ("advertisementId")
);

-- CreateIndex
CREATE UNIQUE INDEX "advertisementCPMs_advertisementId_key" ON "advertisementCPMs"("advertisementId");

-- AddForeignKey
ALTER TABLE "advertisementCPMs" ADD CONSTRAINT "advertisementCPMs_advertisementId_fkey" FOREIGN KEY ("advertisementId") REFERENCES "advertisements"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
