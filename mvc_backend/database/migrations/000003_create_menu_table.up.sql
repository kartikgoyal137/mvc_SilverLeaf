CREATE TABLE IF NOT EXISTS `menu`(
  `product_id` int PRIMARY KEY NOT NULL,
  `product_name` varchar(100) DEFAULT null,
  `category_id` int DEFAULT null,
  `price` decimal(10,2) DEFAULT null,
  `image_url` varchar(255) DEFAULT null
);

CREATE INDEX `category_id` ON `menu` (`category_id`);

CREATE INDEX `idx_product_name` ON `menu` (`product_name`);