-- Добавляем статусы походов.
INSERT INTO hike_statuses(status_name)
VALUES ('создан'),
       ('в работе'),
       ('завершён'),
       ('отменён'),
       ('удалён');

-- Добавляем статусы городов.
INSERT INTO city_statuses(status_name)
VALUES ('существует'), ('уничтожен');

-- Добавляем города.
INSERT INTO cities(city_name, status_id, description, image_url)
VALUES ('Категат (Хедебю)', '1', 'Категат был одним из самых важных викингских городов, расположенных на острове Хейланд в Дании. Город был известен своими морскими торговыми маршрутами и фортификациями.', 'https://fiord.org/wp-content/uploads/2016/12/Hedeby-now.jpg');

INSERT INTO cities(city_name, status_id, description, image_url)
VALUES('Королевство Йорвик', '1', 'Этот викингский город, известный как Йорк, был важным торговым и административным центром во времена викингов.\nДублин (Ирландия): Викинги основали Дублин в 9 веке. Этот город был известен своей торговлей и культурным влиянием в регионе.', 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/fc/Kingdom_of_Jórvik.svg/400px-Kingdom_of_Jórvik.svg.png');

INSERT INTO cities(city_name, status_id, description, image_url)
VALUES('Новгород (Россия)', '1', 'Викинги создали поселение в Новгороде, что сделало его важным торговым центром на востоке.", Area: "1000 кв. м.', 'https://avatars.dzeninfra.ru/get-zen_doc/3229639/pub_5ec7ba99f66e3c72366a38ba_5ec9073987eb5a1725e2438d/scale_1200');

INSERT INTO cities(city_name, status_id, description, image_url)
VALUES('Лунд (Швеция)', '1', 'Лунд был одним из первых викингских городов и центром вероисповедания викингов.', 'https://traveltimes.ru/wp-content/uploads/2021/10/Кафедральный-собор-Лунда-2048x1333.jpg');

INSERT INTO cities(city_name, status_id, description, image_url)
VALUES('Гриндавик (Исландия)', '1', 'Этот город на Исландии служил базой для викингских мореплавателей и рыбаков.', 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/Iceland_adm_location_map.svg/500px-Iceland_adm_location_map.svg.png');

INSERT INTO cities(city_name, status_id, description, image_url)
VALUES('Висбю (Швеция)', '1', 'Этот город на острове Готланд был важным торговым и административным центром викингов.', 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/49/Visby_13-.JPG/580px-Visby_13-.JPG');

INSERT INTO cities(city_name, status_id, description, image_url)
VALUES('Лимфьорд (Дания)', '1', 'Лимфьорд был важным морским перекрестком и базой для викингских экспедиций.', 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9f/Jytland.Limfjord.jpg/240px-Jytland.Limfjord.jpg');

INSERT INTO cities(city_name, status_id, description, image_url)
VALUES('Рейкьявик (Исландия)', '1', 'Викинги основали Рейкьявик, который со временем стал столицей Исландии.', 'https://traveller-eu.ru/sites/default/files/styles/main_img/public/evelyn-paris-WvPQYDd-3Ow-unsplash.webp?itok=PHKdX3SG');

INSERT INTO cities(city_name, status_id, description, image_url)
VALUES('Торшавн (Фарерские острова)', '1', 'Этот город служил базой для викингов на Фарерских островах и был важным для контроля торговых путей в Северном Атлантическом регионе.', 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Tórshavn_skansin_8.jpg/500px-Tórshavn_skansin_8.jpg');

INSERT INTO cities(city_name, status_id, description, image_url)
VALUES('Нормандия', '1', 'Викинг Ролло основал герцогство Нормандия в IX веке после договора с франкским королём. Нормандия стала известной своими влияниями на культуру и историю.', 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/90/Mont-Saint-Michel-2004.jpg/440px-Mont-Saint-Michel-2004.jpg');

INSERT INTO cities(city_name, status_id, description, image_url)
VALUES('Великая Зима', '1', 'Это викингское поселение было обнаружено на территории современной России, недалеко от Волги.', 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/91/Venice_frozen_lagoon_1708.jpg/400px-Venice_frozen_lagoon_1708.jpg');