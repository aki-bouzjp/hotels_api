-- AlterTable
ALTER TABLE `hotels` ALTER COLUMN `id` DROP DEFAULT;

-- AlterTable
ALTER TABLE `images` ALTER COLUMN `id` DROP DEFAULT;

-- AlterTable
ALTER TABLE `room_images` ALTER COLUMN `id` DROP DEFAULT;

-- AlterTable
ALTER TABLE `rooms` ALTER COLUMN `id` DROP DEFAULT;

-- AlterTable
ALTER TABLE `thumbnail_images` ALTER COLUMN `id` DROP DEFAULT;

-- AlterTable
ALTER TABLE `thumbnails` ALTER COLUMN `id` DROP DEFAULT;

-- AlterTable
ALTER TABLE `vacancies` ALTER COLUMN `id` DROP DEFAULT;
