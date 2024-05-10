-- CreateTable
CREATE TABLE "channels" (
    "id" TEXT NOT NULL,
    "ownerId" TEXT NOT NULL,
    "userName" TEXT NOT NULL,
    "name" TEXT,
    "description" TEXT,
    "botAddAsAdmin" BOOLEAN NOT NULL DEFAULT false,
    "channelCreatedAt" TIMESTAMP(3),
    "country" TEXT,
    "currency" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deletedAt" TIMESTAMP(3),

    CONSTRAINT "channels_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "channels_userName_key" ON "channels"("userName");
