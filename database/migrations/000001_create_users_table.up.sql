CREATE TABLE `users` (
  `user_id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `first_name` varchar(30) DEFAULT null,
  `last_name` varchar(30) DEFAULT null,
  `contact` varchar(10) DEFAULT null,
  `email` varchar(50) DEFAULT null,
  `password_hash` varchar(255) DEFAULT null,
  `role` ENUM ('administrator', 'customer', 'chef') DEFAULT 'customer'
);

CREATE UNIQUE INDEX `uk_email` ON `users` (`email`);