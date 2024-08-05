CREATE TABLE `customers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 NOT NULL,
  `email` varchar(255) CHARACTER SET utf8 NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `customer_loans` (
  `id` int NOT NULL AUTO_INCREMENT,
  `customer_id` int NOT NULL,
  `last_payment_term` int NOT NULL,
  `interest_rate` decimal(3,2) NOT NULL,
  `term_in_week` SMALLINT UNSIGNED NOT NULL,
  `principal` BIGINT UNSIGNED NOT NULL,
  `installment` BIGINT UNSIGNED NOT NULL,
  `outstanding_loan` BIGINT UNSIGNED NOT NULL,
  `delinquent` TINYINT(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES customers(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `customer_billing_schedules` (
  `id` int NOT NULL AUTO_INCREMENT,
  `customer_id` int NOT NULL,
  `customer_loan_id` int NOT NULL,
  `payment_term` int NOT NULL,
  `start_date` timestamp NOT NULL,
  `due_date` timestamp NOT NULL,
  `paid_at` timestamp NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT fk_customer_loan_id FOREIGN KEY (customer_id) REFERENCES customers(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;