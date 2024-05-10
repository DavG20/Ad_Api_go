-- CreateTable
CREATE TABLE "categories" (
    "category" TEXT NOT NULL,

    CONSTRAINT "categories_pkey" PRIMARY KEY ("category")
);

-- CreateTable
CREATE TABLE "channelCategories" (
    "channelId" TEXT NOT NULL,
    "category" TEXT NOT NULL,

    CONSTRAINT "channelCategories_pkey" PRIMARY KEY ("channelId","category")
);
