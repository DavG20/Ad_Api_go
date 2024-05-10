-- CreateEnum
CREATE TYPE "AdStatus" AS ENUM ('Pending', 'Running', 'Closed', 'Rejected');

-- CreateTable
CREATE TABLE "advertisements" (
    "id" TEXT NOT NULL,
    "campaignId" TEXT NOT NULL,
    "contentId" TEXT NOT NULL,
    "channelId" TEXT NOT NULL,
    "messageId" INTEGER,
    "status" "AdStatus" NOT NULL DEFAULT 'Pending',
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "advertisements_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "advertisements" ADD CONSTRAINT "advertisements_contentId_fkey" FOREIGN KEY ("contentId") REFERENCES "contents"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "advertisements" ADD CONSTRAINT "advertisements_campaignId_fkey" FOREIGN KEY ("campaignId") REFERENCES "campaigns"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
