-- +goose Up
ALTER TABLE users
DROP COLUMN telegram_id,
DROP COLUMN telegram_nickname,
ADD COLUMN login TEXT NOT NULL,
ADD COLUMN phone_number TEXT,
ADD UNIQUE (login);

CREATE TABLE city (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE street (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    city_id bigint NOT NULL,
    FOREIGN KEY (city_id) REFERENCES city(id)
);

CREATE TABLE type_of_housing (
    id SERIAL PRIMARY KEY,
    type TEXT UNIQUE NOT NULL
);

ALTER TABLE offers
ADD COLUMN city_id BIGINT NOT NULL,
ADD COLUMN street_id BIGINT NOT NULL,
ADD COLUMN type_of_housing_id BIGINT NOT NULL,
DROP COLUMN city,
DROP COLUMN street,
DROP COLUMN type,
ADD CONSTRAINT fk_city FOREIGN KEY (city_id) REFERENCES city(id),
ADD CONSTRAINT fk_street FOREIGN KEY (street_id) REFERENCES street(id),
ADD CONSTRAINT fk_type_of_housing FOREIGN KEY (type_of_housing_id) REFERENCES type_of_housing(id);

CREATE TABLE landlord (
   	id bigserial primary key,   
    rating DECIMAL(3, 2),
    reviews_count INTEGER DEFAULT 0,
    deals_count INTEGER DEFAULT 0,
    user_id bigint NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

create table reviews (
    id bigserial primary key,
  	review text not null,
    rate SMALLINT not null,
  	created_at timestamp not null,
    landlord_id bigint not null,
    author_id bigint not null,
    constraint fk_landlord
        foreign key(landlord_id) 
            references landlord(id) 
            on delete cascade
            on update cascade,
    constraint fk_author
        foreign key(author_id) 
            references users(id)
            on delete cascade
            on update cascade
);

CREATE TABLE notification_types (
    id SERIAL PRIMARY KEY,
    type text UNIQUE NOT NULL
);

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    status text NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    sent_at TIMESTAMP,
    user_id BIGINT NOT NULL,
    type_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (type_id) REFERENCES notification_types(id)
);

-- +goose Down
ALTER TABLE offers
DROP CONSTRAINT fk_city,
DROP CONSTRAINT fk_street,
DROP CONSTRAINT fk_type_of_housing,
DROP COLUMN city_id,
DROP COLUMN street_id,
DROP COLUMN type_of_housing_id,
ADD COLUMN city text not null,
ADD COLUMN street text not null,
ADD COLUMN type text not null;

drop table reviews;
drop table landlord;
drop table street;
drop table city;
drop table type_of_housing;
drop table notifications;
drop table notification_types;

ALTER TABLE users
DROP COLUMN login,
DROP COLUMN phone_number,
ADD COLUMN telegram_id bigint not null,
ADD COLUMN telegram_nickname text not null,
ADD UNIQUE (telegram_id),
ADD UNIQUE (telegram_nickname);