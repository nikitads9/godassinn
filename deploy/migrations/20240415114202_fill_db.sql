-- +goose Up
-- +goose StatementBegin
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

-- Заполнение таблицы пользователей
INSERT INTO users (name, login, password, phone_number, created_at) VALUES
('Pavel Durov', 'pavel_durov', '$2a$14$6VcKoOnnXRmjZ4eZKL4tlOWHyVzNg.ph/1dZ8vbapv/PNlai3czaq', '89771384545', NOW());

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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users;
DELETE FROM landlord;
DELETE FROM offers;

-- +goose StatementEnd