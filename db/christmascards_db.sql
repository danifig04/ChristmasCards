DROP DATABASE IF EXISTS christmascards;

CREATE DATABASE christmascards;

USE arcades;

CREATE TABLE arcades (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    arcade_name varchar(255),
    entry_price int
);

CREATE TABLE games (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    game_title varchar(255),
    rating varchar(255)
);

CREATE TABLE game_cabinets (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    game_id int NOT NULL,
    arcade_id int NOT NULL,
    price_per_play int NOT NULL,

    FOREIGN KEY (game_id) REFERENCES games(id),
    FOREIGN KEY (arcade_id) REFERENCES arcades(id)
);

INSERT INTO arcades (id, arcade_name, entry_price) VALUES (1, 'Andrea''s Palace', 20);
INSERT INTO arcades (id, arcade_name, entry_price) VALUES (2, 'Holly''s Hovel', 2);
INSERT INTO arcades (id, arcade_name, entry_price) VALUES (3, 'Lindsay''s Lair', 3);
INSERT INTO arcades (id, arcade_name, entry_price) VALUES (4, 'Buffy''s Business', 12);

INSERT INTO games (id, game_title, rating) VALUES (1, 'Pac Man', 'E');
INSERT INTO games (id, game_title, rating) VALUES (2, 'Galaga', 'E');
INSERT INTO games (id, game_title, rating) VALUES (3, 'Asteroids', 'M');
INSERT INTO games (id, game_title, rating) VALUES (4, 'Simpsons', 'T');
INSERT INTO games (id, game_title, rating) VALUES (5, 'Deal or No Deal', 'AO');
INSERT INTO games (id, game_title, rating) VALUES (6, 'Mario', 'E');
INSERT INTO games (id, game_title, rating) VALUES (7, 'Jurassic Park', 'T');

INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (1, 1, 2);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (2, 1, 1);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (3, 1, 3);

INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (2, 2, 1);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (4, 2, 1);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (4, 2, 1);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (5, 2, 1);

INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (6, 3, 3);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (2, 3, 5);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (3, 3, 1);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (5, 3, 10);

INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (6, 4, 1);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (7, 4, 4);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (3, 4, 1);
INSERT INTO game_cabinets (game_id, arcade_id, price_per_play) VALUES (5, 4, 1);


