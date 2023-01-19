CREATE TABLE Users (
    username    VARCHAR(40) NOT NULL,
    shop        VARCHAR(40) NOT NULL,
    model       VARCHAR(40) NOT NULL,
    url         VARCHAR(2048) NOT NULL,
    PRIMARY KEY(username, shop, model)
);