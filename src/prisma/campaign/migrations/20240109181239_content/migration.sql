-- CreateEnum
CREATE TYPE "ContentType" AS ENUM ('ImageMediaGroup', 'JustText');

-- CreateTable
CREATE TABLE "contents" (
    "id" TEXT NOT NULL,
    "campaignId" TEXT NOT NULL,
    "contentType" "ContentType",
    "description" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "contents_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "contents_campaignId_key" ON "contents"("campaignId");

-- AddForeignKey
ALTER TABLE "contents" ADD CONSTRAINT "contents_campaignId_fkey" FOREIGN KEY ("campaignId") REFERENCES "campaigns"("id") ON DELETE CASCADE ON UPDATE CASCADE;
