-- AddForeignKey
ALTER TABLE "channelCategories" ADD CONSTRAINT "channelCategories_category_fkey" FOREIGN KEY ("category") REFERENCES "categories"("category") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "channelContentLanguages" ADD CONSTRAINT "channelContentLanguages_language_fkey" FOREIGN KEY ("language") REFERENCES "languages"("language") ON DELETE RESTRICT ON UPDATE CASCADE;
