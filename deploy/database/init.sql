CREATE TABLE users (
      id bigserial primary key,
	name text not null,
	login text not null,
	password text not null,
	phone_number text,
	created_at timestamp not null,
	updated_at timestamp,
	unique(login)
);

CREATE TABLE landlord (
  	id bigserial primary key,   
	rating DECIMAL(3, 2),
	reviews_count INTEGER DEFAULT 0,
	deals_count INTEGER DEFAULT 0,
	user_id bigint NOT NULL,
	constraint fk_users
        FOREIGN KEY (user_id) 
            REFERENCES users(id)
                on delete cascade
                on update cascade
);


CREATE TABLE city (
	id SERIAL PRIMARY KEY,
	name TEXT UNIQUE NOT NULL
);

CREATE TABLE street (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	city_id bigint NOT NULL,
    constraint fk_city
        FOREIGN KEY (city_id) 
            REFERENCES city(id)
                on delete cascade
                on update cascade   
);

CREATE TABLE type_of_housing (
	id SERIAL PRIMARY KEY,
	type TEXT UNIQUE NOT NULL
);


CREATE TABLE offers (
	id bigserial primary key,
	name text not null,
	cost integer not null,
	house integer not null,
	rating integer not null,
	beds_count integer not null,
	short_description text not null,
	city_id bigint NOT NULL,
	street_id bigint NOT NULL,
	type_of_housing_id bigint NOT NULL,
    landlord_id bigint NOT NULL,
    CONSTRAINT fk_city 
    FOREIGN KEY (city_id) 
        REFERENCES city(id)
        on delete cascade
        on update cascade,
    CONSTRAINT fk_street 
    FOREIGN KEY (street_id) 
        REFERENCES street(id)
        on delete cascade
        on update cascade,
    CONSTRAINT fk_type_of_housing 
    FOREIGN KEY (type_of_housing_id) 
        REFERENCES type_of_housing(id)
        on delete cascade
        on update cascade,
    CONSTRAINT fk_landlord 
    FOREIGN KEY (landlord_id) 
        REFERENCES landlord(id)
        on delete cascade
        on update cascade
);


CREATE TABLE bookings (
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

CREATE TABLE reviews (
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
	FOREIGN KEY (user_id) REFERENCES users(id)
	    on delete cascade
        on update cascade,
	FOREIGN KEY (type_id) REFERENCES notification_types(id)
		on delete cascade
		on update cascade
);

create index ix_uuid ON bookings using btree (id);
create index ix_start ON bookings using brin (start_date);

create index ix_end ON bookings using brin (end_date);

create index ix_offer ON bookings using btree (offer_id);
create index ix_owner ON bookings using btree (user_id);

create user otelcol with password 'otelcolpassword';
grant SELECT on pg_stat_database to otelcol;

-- Заполнение таблицы пользователей
INSERT INTO users (name, login, password, phone_number, created_at) VALUES
('Иван Иванов', 'ivan_ivanov', '$2a$14$MfUIiC14PQuFCaCTS09Nhu0GjzIFsncWka4gm6Ic6MdDFdehMB3.K', '88005553535', NOW()),
('Мария Петрова', 'maria_petrova', '$2a$14$MfUIiC14PQuFCaCTS09Nhu0GjzIFsncWka4gm6Ic6MdDFdehMB3.K', '88005553535', NOW()),
('Алексей Сидоров', 'alexey_sidorov', '$2a$14$MfUIiC14PQuFCaCTS09Nhu0GjzIFsncWka4gm6Ic6MdDFdehMB3.K', '88005553535', NOW()),
('Елена Васильева', 'elena_vasilieva', '$2a$14$MfUIiC14PQuFCaCTS09Nhu0GjzIFsncWka4gm6Ic6MdDFdehMB3.K', '88005553535', NOW()),
('Николай Морозов', 'nikolay_morozov', '$2a$14$MfUIiC14PQuFCaCTS09Nhu0GjzIFsncWka4gm6Ic6MdDFdehMB3.K', '88005553535', NOW()),
('Светлана Кузнецова', 'svetlana_kuznetsova', '$2a$14$MfUIiC14PQuFCaCTS09Nhu0GjzIFsncWka4gm6Ic6MdDFdehMB3.K', '88005553535', NOW()),
('Дмитрий Смирнов', 'dmitriy_smirnov', '$2a$14$MfUIiC14PQuFCaCTS09Nhu0GjzIFsncWka4gm6Ic6MdDFdehMB3.K', '88005553535', NOW()),
('Ольга Попова', 'olga_popova', '$2a$14$MfUIiC14PQuFCaCTS09Nhu0GjzIFsncWka4gm6Ic6MdDFdehMB3.K', '88005553535', NOW()),
('Андрей Соколов', 'andrey_sokolov', '$2a$14$MfUIiC14PQuFCaCTS09Nhu0GjzIFsncWka4gm6Ic6MdDFdehMB3.K', '88005553535', NOW()),
('Татьяна Михайлова', 'tatyana_mikhailova', '$2a$14$MfUIiC14PQuFCaCTS09Nhu0GjzIFsncWka4gm6Ic6MdDFdehMB3.K', '88005553535', NOW());

-- Заполнение таблицы городов
INSERT INTO city (name) VALUES
('Москва'),
('Солнечногорск'),
('Ярославль'),
('Торжок'),
('Таганрог');

-- Заполнение таблицы улиц
INSERT INTO street (name, city_id) VALUES
('5-я Авеню', 1),
('Бродвей', 1),
('Уолл-Стрит', 1),
('Мэдисон Авеню', 1),
('Парк Авеню', 1),
('Сансет Бульвар', 2),
('Родео Драйв', 2),
('Голливуд Бульвар', 2),
('Вилшир Бульвар', 2),
('Санта-Моника Бульвар', 2);

-- Заполнение таблицы типов жилья
INSERT INTO type_of_housing (type) VALUES
('Квартира'),
('Дом'),
('Таунхаус');

-- Заполнение таблицы арендодателей
INSERT INTO landlord (user_id, rating, reviews_count, deals_count) VALUES
(1, 4.5, 10, 5);


-- Заполнение таблицы предложений
INSERT INTO offers (name, cost, house, rating, beds_count, short_description, city_id, street_id, type_of_housing_id, landlord_id) VALUES
('Уютная квартира', 100, 1, 5, 2, 'Уютная квартира в центре города', 1, 1, 1, 1),
('Просторный дом', 200, 2, 4, 4, 'Просторный дом с задним двором', 1, 2, 2, 1),
('Таунхаус рядом с парком', 150, 3, 5, 3, 'Таунхаус рядом с парком и легкий доступ к пешеходным тропам', 1, 3, 3, 1),
('Модерная квартира', 120, 4, 4, 2, 'Модерная квартира со всеми удобствами', 1, 4, 1, 1),
('Очаровательный коттедж', 180, 5, 5, 3, 'Очаровательный коттедж в тихом районе', 1, 5, 2, 1),
('Лофт в центре города', 110, 6, 4, 1, 'Лофт в центре города, близко к ночной жизни', 1, 1, 1, 1),
('Дом в пригороде', 210, 7, 3, 4, 'Большой дом в пригороде', 1, 2, 2, 1),
('Квартира в центре города', 130, 8, 5, 2, 'Квартира в самом сердце города', 1, 3, 1, 1),
('Кондо у реки', 140, 9, 4, 2, 'Кондо с видом на реку', 1, 4, 1, 1),
('Кабинка в горах', 160, 10, 5, 3, 'Кабинка с видом на горы', 1, 5, 2, 1);
