-- +goose Up
alter table bookings
add column cost integer,
add column city text,
add column street text,
add column house integer,
add column rating integer,
add column type text,
add column beds_count integer,
add column short_description text;

-- +goose Down
alter table bookings
drop column cost integer,
drop column city text,
drop column street text,
drop column house integer,
drop column rating integer,
drop column type text,
drop column beds_count integer,
drop column short_description text;
