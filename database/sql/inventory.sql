CREATE TABLE `users` (
  `id` integer PRIMARY KEY,
  `username` varchar(255),
  `role` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `product` (
  `product_id` integer PRIMARY KEY,
  `product_name` varchar(255),
  `product_description` varchar(255),
  `stock` integer,
  `product_price` double,
  `status` varchar(255),
  `category_id` integer,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `category` (
  `category_id` integer PRIMARY KEY,
  `category_name` varchar(255),
  `category_description` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `category` ADD FOREIGN KEY (`category_id`) REFERENCES `product` (`category_id`);
