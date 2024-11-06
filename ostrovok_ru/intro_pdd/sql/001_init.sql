CREATE SCHEMA intro;
CREATE TYPE intro.KindOfOperation AS ENUM ('INC', 'DEC');
CREATE TYPE intro.KindOfTransfer AS ENUM ('INTERNAL', 'EXTERNAL');

CREATE TABLE intro.accounts (
                                id UUID PRIMARY KEY,
                                amount BIGINT default 0,
                                user_id uuid NOT NULL ,
                                tx_id uuid NOT NULL ,
                                operation intro.KindOfOperation NOT NULL,
                                kind intro.KindOfTransfer NOT NULL
);