-- Добавляем статусы.
INSERT INTO hike_statuses(status_name)
VALUES ('создан'),
       ('в работе'),
       ('завершён'),
       ('отменён'),
       ('удалён');

-- Добавляем статусы.
INSERT INTO city_statuses(status_name)
VALUES ('существует'),
       ('уничтожен');

-- Добавляем авторов.
INSERT INTO authors(author_name, profession, birthday, image_url)
VALUES
    ('Полина Копилова', 'Историк', '2004-11-25', 'https://avatars.dzeninfra.ru/get-zen_doc/3512693/pub_61ec8ca541196a01edbd82a4_61ec8d73ac280b1dc100b04d/scale_1200'),
    ('Иван Петров', 'Биолог', '1978-08-10', 'https://mykaleidoscope.ru/x/uploads/posts/2022-09/1663637203_11-mykaleidoscope-ru-p-uspeshnie-molodie-lyudi-vkontakte-11.jpg'),
    ('Екатерина Сидорова', 'Юрист', '1985-02-22', 'https://fikiwiki.com/uploads/posts/2022-02/1645006638_13-fikiwiki-com-p-kartinki-schastlivikh-lyudei-13.jpg'),
    ('Андрей Ковалев', 'Физик', '1990-07-15', 'https://kz-russia.ru/wp-content/uploads/2/5/3/253feca7856000ad15632eb6ccd31432.jpeg'),
    ('Надежда Григорьева', 'Психолог', '1982-11-30', 'https://inha.ru/wp-content/uploads/2019/07/interesvobshenii.jpg'),
    ('Сергей Морозов', 'Химик', '1975-06-05', 'https://i09.fotocdn.net/s115/6c537bb917abadaa/public_pin_l/2600280989.jpg'),
    ('Ольга Смирнова', 'Инженер', '1988-04-18', 'https://evehealth.ru/wp-content/uploads/2018/08/1452266223_obschenie-blagopriyatno-vliyaet-na-zdorove-uchenye.jpg'),
    ('Михаил Васильев', 'Астроном', '1983-09-20', 'https://webpulse.imgsmail.ru/imgpreview?mb=webpulse&key=lenta_admin-image-a25168c9-57e5-44f6-a0f9-c9a7ef3ad4c2'),
    ('Анна Дмитриева', 'Искусствовед', '1993-03-08', 'https://hr-portal.ru/files/iimg_uploads/17_10/301.jpg'),
    ('Павел Николаев', 'Экономист', '1970-12-12', 'https://w.forfun.com/fetch/22/22cb92e292f7b97e0ccaa9163aca35d0.jpeg'),
    ('Елена Романова', 'Педагог', '1987-01-25', 'http://natid.co.il/content/top_pic/top_pic8.jpg');

-- Создаем походы.
INSERT INTO hikes (hike_name, date_start, date_end, date_start_preparing, author_id, status_id, description)
VALUES
    ('Поход викингов в Англию', '2023-06-10', '2023-06-20', '2023-05-15', 1, 1, 'Эпический поход викингов в Англию.'),
    ('Завоевание Шотландии', '2023-07-05', '2023-07-15', '2023-06-20', 2, 2, 'Викинги отправляются в Шотландию, чтобы завоевать новые земли.'),
    ('Викинги в Ирландии', '2023-08-15', '2023-08-25', '2023-07-30', 3, 3, 'Поход викингов по берегам Ирландии.'),
    ('Разграбление монастырей', '2023-09-10', '2023-09-20', '2023-08-25', 4, 4, 'Викинги ограбляют монастыри и богатства.'),
    ('По следам Одинсвена', '2023-10-10', '2023-10-20', '2023-09-15', 5, 5, 'Исследовательский поход по следам великого воина Одинсвена.'),
    ('Великое морское путешествие', '2023-11-05', '2023-11-15', '2023-10-20', 6, 1, 'Викинги отправляются в долгое морское путешествие.'),
    ('Завоевание Франции', '2023-12-01', '2023-12-11', '2023-11-15', 7, 2, 'Викинги вторгаются во Францию.'),
    ('Поход на Восток', '2024-01-10', '2024-01-20', '2023-12-15', 8, 3, 'Викинги отправляются на восток, исследуя новые земли.'),
    ('Викинги в Скандинавии', '2024-02-15', '2024-02-25', '2024-01-30', 9, 4, 'Поход викингов в родные земли.'),
    ('Поход в Средиземное море', '2024-03-20', '2024-03-30', '2024-02-25', 10, 5, 'Викинги исследуют Средиземное море и его берега.');

-- Вставляем данные о городах
INSERT INTO cities (city_name, description, status_id, image_url)
VALUES
    ('Категат (Хедебю)', 'Категат был одним из самых важных викингских городов, расположенных на острове Хейланд в Дании. Город был известен своими морскими торговыми маршрутами и фортификациями.', 1, 'https://fiord.org/wp-content/uploads/2016/12/Hedeby-now.jpg'),
    ('Королевство Йорвик', 'Этот викингский город, известный как Йорк, был важным торговым и административным центром во времена викингов.', 1, 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/fc/Kingdom_of_Jórvik.svg/400px-Kingdom_of_Jórvik.svg.png'),
    ('Новгород (Россия)', 'Викинги создали поселение в Новгороде, что сделало его важным торговым центром на востоке.', 1, 'https://avatars.dzeninfra.ru/get-zen_doc/3229639/pub_5ec7ba99f66e3c72366a38ba_5ec9073987eb5a1725e2438d/scale_1200'),
    ('Лунд (Швеция)', 'Лунд был одним из первых викингских городов и центром вероисповедания викингов.', 1, 'https://traveltimes.ru/wp-content/uploads/2021/10/Кафедральный-собор-Лунда-2048x1333.jpg'),
    ('Гриндавик (Исландия)', 'Этот город на Исландии служил базой для викингских мореплавателей и рыбаков.', 1, 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/Iceland_adm_location_map.svg/500px-Iceland_adm_location_map.svg.png'),
    ('Висбю (Швеция)', 'Этот город на острове Готланд был важным торговым и административным центром викингов.', 1, 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/49/Visby_13-.JPG/580px-Visby_13-.JPG'),
    ('Лимфьорд (Дания)', 'Лимфьорд был важным морским перекрестком и базой для викингских экспедиций.', 1, 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9f/Jytland.Limfjord.jpg/240px-Jytland.Limfjord.jpg'),
    ('Рейкьявик (Исландия)', 'Викинги основали Рейкьявик, который со временем стал столицей Исландии.', 1, 'https://traveller-eu.ru/sites/default/files/styles/main_img/public/evelyn-paris-WvPQYDd-3Ow-unsplash.webp?itok=PHKdX3SG'),
    ('Торшавн (Фарерские острова)', 'Этот город служил базой для викингов на Фарерских островах и был важным для контроля торговых путей в Северном Атлантическом регионе.', 1, 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Tórshavn_skansin_8.jpg/500px-Tórshavn_skansin_8.jpg'),
    ('Нормандия', 'Викинг Ролло основал герцогство Нормандия в IX веке после договора с франкским королём. Нормандия стала известной своими влияниями на культуру и историю.', 1, 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/90/Mont-Saint-Michel-2004.jpg/440px-Mont-Saint-Michel-2004.jpg'),
    ('Великая Зима', 'Это викингское поселение было обнаружено на территории современной России, недалеко от Волги.', 1, 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/91/Venice_frozen_lagoon_1708.jpg/400px-Venice_frozen_lagoon_1708.jpg');


-- Создаем викингов.
INSERT INTO vikings (viking_name, post, birthday, day_of_death, city_of_birth_id, image_url)
VALUES
    ('Рагнар Лодброк', 'Вождь', '0795-01-01', '865-12-31', 1, 'https://images.kinorium.com/movie/shot/647748/w1500_48906466.jpg'),
    ('Лагерта', 'Воительница', '0800-03-15', '875-11-28', 2, 'https://proprikol.ru/wp-content/uploads/2020/04/kartinki-vikingi-16.jpg'),
    ('Флоки Варг', 'Изобретатель', '0810-08-20', '890-05-10', 3, 'https://w.forfun.com/fetch/17/1765da239e3c01a40f61133c5e12a345.jpeg'),
    ('Бьорн Железнобок', 'Воин', '0820-05-02', '0892-09-15', 4, 'https://proprikol.ru/wp-content/uploads/2020/04/kartinki-vikingi-3.jpg'),
    ('Ивар Бескостный', 'Стратег', '0830-11-10', '0900-08-02', 5, 'https://fikiwiki.com/uploads/posts/2022-02/1645060045_38-fikiwiki-com-p-kartinki-vikingov-43.jpg'),
    ('Ролло', 'Вождь', '0805-07-25', '0875-02-12', 1, 'https://i.pinimg.com/originals/27/16/5e/27165eb2240e1c6d5ffe2c957100ef0b.jpg'),
    ('Аслауг', 'Королева', '0815-12-05', '0890-06-20', 2, 'https://wp-s.ru/wallpapers/6/18/472622147366450/fotografiya-so-semok-istoricheskogo-seriala-vikingi.jpg'),
    ('Харальд Красноволосый', 'Король', '0850-02-22', '0932-09-23', 3, 'https://i.pinimg.com/originals/d9/7c/15/d97c15eabcd87ca9ab1164f6a4098545.jpg'),
    ('Торвальд Кодранссон', 'Исследователь', '840-09-30', '0910-03-18', 4, 'https://www.teahub.io/photos/full/236-2367625_vikings-ragnar-lothbrok-lagertha-and-rollo-season-vikings.jpg'),
    ('Герда', 'Шаман', '0825-04-12', '900-12-04', 5, 'https://proprikol.ru/wp-content/uploads/2020/04/kartinki-vikingi-43.jpg');
