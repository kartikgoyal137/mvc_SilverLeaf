ALTER TABLE `menu` ADD CONSTRAINT `menu_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`category_id`);

ALTER TABLE `product_ingredients` ADD CONSTRAINT `fk_product_ingredients_menu` FOREIGN KEY (`product_id`) REFERENCES `menu` (`product_id`) ON DELETE CASCADE;

ALTER TABLE `orders` ADD CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `serve` ADD CONSTRAINT `serve_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`);

ALTER TABLE `serve` ADD CONSTRAINT `serve_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `menu` (`product_id`);

ALTER TABLE `payments` ADD CONSTRAINT `payments_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `payments` ADD CONSTRAINT `payments_ibfk_2` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`);