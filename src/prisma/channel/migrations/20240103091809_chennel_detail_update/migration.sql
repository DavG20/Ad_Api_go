-- CreateTable
CREATE TABLE "channelDetails" (
    "id" TEXT NOT NULL,
    "channelId" TEXT NOT NULL,
    "subCount" INTEGER NOT NULL DEFAULT 0,
    "averagePostView" INTEGER NOT NULL DEFAULT 0,
    "postToSubRatio" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "postFrequecy" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "lastPostId" TEXT,
    "lastPost" TIMESTAMP(3),
    "collectedDate" TIMESTAMP(3),
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "channelDetails_pkey" PRIMARY KEY ("id")
);
