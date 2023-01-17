CREATE TABLE Products (
    shop        VARCHAR(40) NOT NULL,
    model       VARCHAR(40) NOT NULL,
    url         TEXT NOT NULL,
    update_ts   TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    price       INTEGER NOT NULL, -- price in polish cents
    PRIMARY KEY(shop, model, update_ts)
);