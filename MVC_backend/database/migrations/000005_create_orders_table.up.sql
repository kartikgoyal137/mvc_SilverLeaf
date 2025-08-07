CREATE TABLE `orders` (
  `order_id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT null,
  `status` ENUM ('Yet to start', 'Cooking', 'Completed') DEFAULT null,
  `created_at` timestamp DEFAULT (CURRENT_TIMESTAMP),
  `instructions` text,
  `table_no` int DEFAULT null
);

CREATE INDEX `user_id` ON `orders` (`user_id`);

CREATE INDEX `idx_status_created_at` ON `orders` (`status`, `created_at`);