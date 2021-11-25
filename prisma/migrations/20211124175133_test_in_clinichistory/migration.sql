/*
  Warnings:

  - Added the required column `patientId` to the `ClinicHistory` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "Patient" DROP CONSTRAINT "Patient_clinicHistoryId_fkey";

-- AlterTable
ALTER TABLE "ClinicHistory" ADD COLUMN     "patientId" INTEGER NOT NULL;

-- AddForeignKey
ALTER TABLE "ClinicHistory" ADD FOREIGN KEY ("patientId") REFERENCES "Patient"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
