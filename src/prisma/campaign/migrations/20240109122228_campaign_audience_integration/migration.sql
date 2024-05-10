-- CreateTable
CREATE TABLE "audiences" (
    "campaignId" TEXT NOT NULL,
    "category" TEXT,
    "language" TEXT,

    CONSTRAINT "audiences_pkey" PRIMARY KEY ("campaignId")
);

-- CreateIndex
CREATE UNIQUE INDEX "audiences_campaignId_key" ON "audiences"("campaignId");

-- AddForeignKey
ALTER TABLE "audiences" ADD CONSTRAINT "audiences_campaignId_fkey" FOREIGN KEY ("campaignId") REFERENCES "campaigns"("id") ON DELETE CASCADE ON UPDATE CASCADE;
