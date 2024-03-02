CREATE TABLE clients (
    id int,
    account_limit int,
    balance int,
    PRIMARY KEY (id)
);

CREATE TABLE transactions (
    id int,
    client_id int,
    transaction_value int,
    transaction_type varchar(1),
    transaction_description varchar(10),
    created_at date,
    PRIMARY KEY (id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE INDEX transaction_client_id_idx
ON transactions (client_id);

DO $$
BEGIN
  INSERT INTO clients (id, account_limit, balance)
  VALUES
    (1, 1000 * 100, 0),
    (2, 800 * 100, 0),
    (3, 10000 * 100, 0),
    (4, 100000 * 100, 0),
    (5, 5000 * 100, 0);
END; $$