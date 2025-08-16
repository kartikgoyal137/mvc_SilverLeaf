CREATE TABLE IF NOT EXISTS `categories`  (
  `category_id` int PRIMARY KEY NOT NULL,
  `category_name` varchar(20) DEFAULT null,
  `description` text,
  `image_url` varchar(255) DEFAULT null
);

CREATE INDEX `idx_category_name` ON `categories` (`category_name`);