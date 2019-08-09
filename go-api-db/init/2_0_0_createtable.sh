#!/bin/bash
psql -U root -d goapi << "EOSQL"

CREATE SEQUENCE ID_SEQ
    INCREMENT BY 1
    MAXVALUE 99999999
    START WITH 1
    NO CYCLE
;

CREATE TABLE ACCOUNT (
    id int PRIMARY KEY,
    name text
);

EOSQL