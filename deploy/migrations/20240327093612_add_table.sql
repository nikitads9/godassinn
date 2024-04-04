-- +goose Up
create table users (
	id bigserial primary key,
    name text not null,
	telegram_id bigint not null,
    telegram_nickname text not null,
    password text not null,
    created_at timestamp not null,
    updated_at timestamp,
    unique(telegram_id),
    unique(telegram_nickname)
);

create table offers (
    id bigserial primary key,
    name text not null,
    cost integer not null,
    city text not null,
    street text not null,
    house integer not null,
    rating integer not null,
    type text not null,
    beds_count integer not null,
    short_description text not null
);

create table bookings (
    id uuid primary key,
    start_date timestamp not null,
    end_date timestamp not null,
    notify_at interval default '0s',
    created_at timestamp not null,
    updated_at timestamp,
    offer_id bigint not null,
    user_id bigint not null,
    constraint fk_offers
        foreign key(offer_id) 
            references offers(id) 
            on delete cascade
            on update cascade,
    constraint fk_users
        foreign key(user_id) 
            references users(id)
            on delete cascade
            on update cascade
);

create index ix_uuid ON bookings using btree (id);
create index ix_start ON bookings using brin (start_date);
create index ix_end ON bookings using brin (end_date);
create index ix_offer ON bookings using btree (offer_id);
create index ix_owner ON bookings using btree (user_id);
create user otelcol with password 'otelcolpassword';
grant select on pg_stat_database to otelcol;

-- +goose Down
drop table bookings;
drop table users;
drop table offers;
revoke select on pg_stat_database from otelcol;
drop user otelcol;

