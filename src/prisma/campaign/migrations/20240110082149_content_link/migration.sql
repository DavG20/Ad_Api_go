-- CreateTable
CREATE TABLE "contentLinks" (
    "contentId" TEXT NOT NULL,
    "link" TEXT,
    "title" TEXT,

    CONSTRAINT "contentLinks_pkey" PRIMARY KEY ("contentId")
);

-- CreateIndex
CREATE UNIQUE INDEX "contentLinks_contentId_key" ON "contentLinks"("contentId");

-- AddForeignKey
ALTER TABLE "contentLinks" ADD CONSTRAINT "contentLinks_contentId_fkey" FOREIGN KEY ("contentId") REFERENCES "contents"("id") ON DELETE CASCADE ON UPDATE CASCADE;
