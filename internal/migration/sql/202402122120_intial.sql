-- +goose Up
CREATE TABLE transactions
(
    id         SERIAL PRIMARY KEY,
    number     VARCHAR(255) NOT NULL,
    apple_id   VARCHAR(255),
    status     VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders
(
    id             SERIAL PRIMARY KEY,
    number         VARCHAR(255) NOT NULL,
    customer_id    VARCHAR(255) NOT NULL,
    order_type     VARCHAR(255) NOT NULL,
    contract_date  TIMESTAMP    NOT NULL,
    transaction_id INTEGER,
    created_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE orders
    ADD CONSTRAINT fk_transaction_id
        FOREIGN KEY (transaction_id)
            REFERENCES transactions (id);

CREATE TABLE deliveries
(
    id         SERIAL PRIMARY KEY,
    number     VARCHAR(255) NOT NULL,
    order_id   INTEGER,
    status     VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE deliveries
    ADD CONSTRAINT fk_order_id
        FOREIGN KEY (order_id)
            REFERENCES orders (id);

CREATE TABLE devices
(
    id          SERIAL PRIMARY KEY,
    imei        VARCHAR(255) NOT NULL,
    delivery_id INTEGER,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE devices
    ADD CONSTRAINT fk_delivery_id
        FOREIGN KEY (delivery_id)
            REFERENCES deliveries (id);


-- +goose Down
DROP TABLE devices;
DROP TABLE deliveries;
DROP TABLE orders;
DROP TABLE transactions;