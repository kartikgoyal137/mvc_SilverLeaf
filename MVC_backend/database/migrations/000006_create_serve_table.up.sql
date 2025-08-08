CREATE TABLE `serve` (
  `order_id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `quantity` int DEFAULT '0',
  PRIMARY KEY (`order_id`, `product_id`)
);
CREATE INDEX `product_id` ON `serve` (`product_id`);