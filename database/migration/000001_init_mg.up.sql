CREATE TABLE `users` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255),
  `role` varchar(255),
  `created_at` timestamp
  );

CREATE TABLE `product` (
  `product_id` integer PRIMARY KEY AUTO_INCREMENT,
  `product_name` varchar(255),
  `product_description` varchar(255),
  `stock` integer,
  `product_price` double,
  `status` varchar(255),
  `category_id` integer,
  `created_at` timestamp
);

CREATE TABLE `category` (
  `category_id` integer PRIMARY KEY AUTO_INCREMENT,
  `category_name` varchar(255),
  `category_description` varchar(255),
  `created_at` timestamp
);

ALTER TABLE `product` ADD FOREIGN KEY (`category_id`) REFERENCES `category` (`category_id`);
