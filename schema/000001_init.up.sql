CREATE TYPE cottageType as ENUM ('House', 'Dacha','Area');
CREATE TYPE rent as ENUM ('flatsForDay', 'flatsLongRent','cottagesForDay','cottagesLongRent');

CREATE TABLE flats
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR      NOT NULL,
    mapMarker   VARCHAR(128) NOT NULL,
    price       VARCHAR(32)  NOT NULL,
    info        VARCHAR      NOT NULL UNIQUE,
    description VARCHAR      NOT NULL UNIQUE,
    link        VARCHAR      NOT NULL UNIQUE,
    rooms       INT          NOT NULL
);



CREATE TABLE cottages
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR      NOT NULL,
    mapMarker   VARCHAR(128) NOT NULL,
    price       VARCHAR(32)  NOT NULL,
    info        VARCHAR      NOT NULL UNIQUE,
    description VARCHAR      NOT NULL UNIQUE,
    link        VARCHAR      NOT NULL UNIQUE,
    type        cottageType  NOT NULL
);


CREATE TABLE flats_for_rent
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR      NOT NULL,
    mapMarker   VARCHAR(128) NOT NULL,
    price       VARCHAR(32)  NOT NULL,
    info        VARCHAR      NOT NULL UNIQUE,
    description VARCHAR      NOT NULL UNIQUE,
    link        VARCHAR      NOT NULL UNIQUE,
    rooms       INT,
    type        rent         NOT NULL
);