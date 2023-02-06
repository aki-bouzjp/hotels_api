-- CreateTable
CREATE TABLE `hotels` (
    `id` VARCHAR(32) DEFAULT (uuid_short()) NOT NULL,
    `happy_hotel_id` VARCHAR(32) NULL,
    `hotel_type` VARCHAR(64) NOT NULL,
    `lng` DECIMAL(11, 8) NOT NULL,
    `lat` DECIMAL(11, 8) NOT NULL,
    `name` VARCHAR(255) NULL,
    `address` VARCHAR(255) NULL,
    `url` VARCHAR(1024) NULL,
    `homepage` VARCHAR(1024) NULL,
    `tel` VARCHAR(32) NULL,
    `pr` VARCHAR(1024) NULL,
    `score` DECIMAL(2, 1) NOT NULL DEFAULT 0.0,
    `single_price` INTEGER NULL,
    `rest_min_price` INTEGER NULL,
    `lodging_min_price` INTEGER NULL,
    `transfer_time_by_walk` VARCHAR(255) NULL,
    `transfer_time_by_drive` VARCHAR(255) NULL,
    `room_count` TINYINT NULL,
    `parking_count` TINYINT NULL,
    `parking_count_by_highroof` TINYINT NULL,
    `parking_description` VARCHAR(1024) NULL,
    `mapcode` VARCHAR(255) NULL,
    `pricing_plans` JSON NULL,
    `online_booking` BOOLEAN NOT NULL DEFAULT false,
    `online_checking` BOOLEAN NOT NULL DEFAULT false,
    `created_at` DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` DATETIME(0) NOT NULL,

    INDEX `hotels_id_idx`(`id`),
    INDEX `hotels_lng_lat_idx`(`lng`, `lat`),
    UNIQUE INDEX `hotels_happy_hotel_id_key`(`happy_hotel_id`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `vacancies` (
    `id` VARCHAR(32) DEFAULT (uuid_short()) NOT NULL,
    `hotel_id` VARCHAR(32) NOT NULL,
    `empty` BOOLEAN NOT NULL DEFAULT false,
    `created_at` DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` DATETIME(0) NOT NULL,

    INDEX `vacancies_hotel_id_idx`(`hotel_id`),
    UNIQUE INDEX `vacancies_hotel_id_key`(`hotel_id`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `thumbnails` (
    `id` VARCHAR(32) DEFAULT (uuid_short()) NOT NULL,
    `hotel_id` VARCHAR(32) NOT NULL,
    `created_at` DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` DATETIME(0) NOT NULL,

    INDEX `thumbnails_hotel_id_idx`(`hotel_id`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `thumbnail_images` (
    `id` VARCHAR(32) DEFAULT (uuid_short()) NOT NULL,
    `size` VARCHAR(32) NOT NULL DEFAULT 'original',
    `thumbnail_id` VARCHAR(32) NOT NULL,
    `image_id` VARCHAR(32) NOT NULL,
    `created_at` DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` DATETIME(0) NOT NULL,

    INDEX `thumbnail_images_thumbnail_id_idx`(`thumbnail_id`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `rooms` (
    `id` VARCHAR(32) DEFAULT (uuid_short()) NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    `hotel_id` VARCHAR(32) NOT NULL,
    `created_at` DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` DATETIME(0) NOT NULL,

    INDEX `rooms_hotel_id_idx`(`hotel_id`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `room_images` (
    `id` VARCHAR(32) DEFAULT (uuid_short()) NOT NULL,
    `size` VARCHAR(32) NOT NULL DEFAULT 'original',
    `room_id` VARCHAR(32) NOT NULL,
    `image_id` VARCHAR(32) NOT NULL,
    `created_at` DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` DATETIME(0) NOT NULL,

    INDEX `room_images_room_id_idx`(`room_id`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `images` (
    `id` VARCHAR(32) DEFAULT (uuid_short()) NOT NULL,
    `mime_type` VARCHAR(64) NOT NULL,
    `url` VARCHAR(1024) NOT NULL,
    `width` SMALLINT NOT NULL,
    `height` SMALLINT NOT NULL,
    `created_at` DATETIME(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` DATETIME(0) NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- AddForeignKey
ALTER TABLE `vacancies` ADD CONSTRAINT `vacancies_hotel_id_fkey` FOREIGN KEY (`hotel_id`) REFERENCES `hotels`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `thumbnails` ADD CONSTRAINT `thumbnails_hotel_id_fkey` FOREIGN KEY (`hotel_id`) REFERENCES `hotels`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `thumbnail_images` ADD CONSTRAINT `thumbnail_images_thumbnail_id_fkey` FOREIGN KEY (`thumbnail_id`) REFERENCES `thumbnails`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `thumbnail_images` ADD CONSTRAINT `thumbnail_images_image_id_fkey` FOREIGN KEY (`image_id`) REFERENCES `images`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `rooms` ADD CONSTRAINT `rooms_hotel_id_fkey` FOREIGN KEY (`hotel_id`) REFERENCES `hotels`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `room_images` ADD CONSTRAINT `room_images_room_id_fkey` FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `room_images` ADD CONSTRAINT `room_images_image_id_fkey` FOREIGN KEY (`image_id`) REFERENCES `images`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;
