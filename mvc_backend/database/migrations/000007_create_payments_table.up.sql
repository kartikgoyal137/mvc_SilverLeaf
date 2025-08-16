CREATE TABLE IF NOT EXISTS `payments` (
  `transaction_id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `order_id` int DEFAULT null,
  `user_id` int DEFAULT null,
  `food_total` float DEFAULT null,
  `created_at` timestamp DEFAULT null,
  `tip` int DEFAULT null,
  `status` ENUM ('Pending', 'Completed') DEFAULT 'Pending'
);

CREATE INDEX `Payments_index_0` ON `payments` (`transaction_id`);

CREATE INDEX `user_id` ON `payments` (`user_id`);

CREATE INDEX `payments_ibfk_2` ON `payments` (`order_id`);