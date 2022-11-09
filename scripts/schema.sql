USE usersbalance;

CREATE TABLE IF NOT EXISTS accounts (
    user_id BIGINT UNSIGNED UNIQUE NOT NULL PRIMARY KEY,
    balance MEDIUMINT UNSIGNED NOT NULL
);

CREATE TABLE IF NOT EXISTS report_accounting (
    service_id BIGINT UNSIGNED UNIQUE NOT NULL,
    year YEAR NOT NULL,
    month SMALLINT UNSIGNED NOT NULL,
    income INT UNSIGNED NOT NULL,
    PRIMARY KEY (service_id, year, month)
);

CREATE TABLE IF NOT EXISTS reserved_accounts (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    order_id BIGINT UNSIGNED UNIQUE NOT NULL,
    service_id BIGINT UNSIGNED NOT NULL,
    cost MEDIUMINT UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES accounts(user_id)
);



CREATE TABLE IF NOT EXISTS transactions_history(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    operation TEXT NOT NULL,
    comments TEXT,
    time DATETIME,
    sum MEDIUMINT UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES accounts(user_id),
    CHECK ( operation IN ('начисление', 'списание', 'возврат'))
);