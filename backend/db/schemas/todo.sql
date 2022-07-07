CREATE TABLE `todo` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(40) COLLATE utf8mb4_bin NOT NULL,
  `content` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
