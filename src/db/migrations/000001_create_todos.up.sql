CREATE TABLE `todos` (
  `id` varchar(64) NOT NULL,
  `title` varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL NULL,
  `desc` varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;