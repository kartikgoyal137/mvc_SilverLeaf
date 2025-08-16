ALTER TABLE `menu` DROP FOREIGN KEY `menu_ibfk_1`;

ALTER TABLE `product_ingredients` DROP FOREIGN KEY `fk_product_ingredients_menu`;

ALTER TABLE `orders` DROP FOREIGN KEY `orders_ibfk_1`;

ALTER TABLE `serve` DROP FOREIGN KEY `serve_ibfk_1`;

ALTER TABLE `serve` DROP FOREIGN KEY `serve_ibfk_2`;

ALTER TABLE `payments` DROP FOREIGN KEY `payments_ibfk_1`;

ALTER TABLE `payments` DROP FOREIGN KEY `payments_ibfk_2`;