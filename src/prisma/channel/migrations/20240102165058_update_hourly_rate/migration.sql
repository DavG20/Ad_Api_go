-- CreateTable
CREATE TABLE "hourlyRates" (
    "id" TEXT NOT NULL,
    "channelId" TEXT NOT NULL,
    "active" BOOLEAN NOT NULL DEFAULT true,
    "hourlyRate" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "minHourlyVolume" DOUBLE PRECISION NOT NULL DEFAULT 0,
    "maxHourlyVolume" DOUBLE PRECISION NOT NULL DEFAULT 10,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "removeAt" TIMESTAMP(3),

    CONSTRAINT "hourlyRates_pkey" PRIMARY KEY ("id")
);
