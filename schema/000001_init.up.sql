CREATE TYPE cottage_type as ENUM ('house', 'dacha', 'area');
CREATE TYPE rent as ENUM ('flatDayRent', 'flatLongRent', 'cottageDayRent', 'cottageLongRent');

CREATE TABLE flats
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR      NOT NULL,
    map_marker  VARCHAR(128) NOT NULL,
    price       VARCHAR(32)  NOT NULL,
    inf         VARCHAR      NOT NULL,
    description VARCHAR      NOT NULL,
    link        VARCHAR      NOT NULL UNIQUE,
    rooms       INT          NOT NULL
);

CREATE TABLE cottages
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR      NOT NULL,
    map_marker  VARCHAR(128) NOT NULL,
    price       VARCHAR(32)  NOT NULL,
    inf         VARCHAR      NOT NULL,
    description VARCHAR      NOT NULL,
    link        VARCHAR      NOT NULL UNIQUE,
    typ         cottage_type NOT NULL
);

CREATE TABLE flats_for_rent
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR      NOT NULL,
    map_marker  VARCHAR(128) NOT NULL,
    price       VARCHAR      NOT NULL,
    inf         VARCHAR      NOT NULL,
    description VARCHAR      NOT NULL,
    link        VARCHAR      NOT NULL UNIQUE,
    rooms       INT,
    typ         rent         NOT NULL
);