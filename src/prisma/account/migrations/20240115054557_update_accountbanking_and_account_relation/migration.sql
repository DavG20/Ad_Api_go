-- AddForeignKey
ALTER TABLE "accountBankings" ADD CONSTRAINT "accountBankings_accountId_fkey" FOREIGN KEY ("accountId") REFERENCES "accounts"("id") ON DELETE CASCADE ON UPDATE CASCADE;
