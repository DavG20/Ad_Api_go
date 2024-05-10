-- CreateTable
CREATE TABLE "languages" (
    "language" TEXT NOT NULL,

    CONSTRAINT "languages_pkey" PRIMARY KEY ("language")
);

-- CreateTable
CREATE TABLE "channelContentLanguages" (
    "channelId" TEXT NOT NULL,
    "language" TEXT NOT NULL,

    CONSTRAINT "channelContentLanguages_pkey" PRIMARY KEY ("channelId","language")
);
