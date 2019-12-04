USE example;

CREATE TABLE example1 (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(50) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE example2 (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(50) NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO example1 (name) VALUES ("mike");
INSERT INTO example1 (name) VALUES ("tony");
INSERT INTO example1 (name) VALUES ("billy");
INSERT INTO example1 (name) VALUES ("grace");
INSERT INTO example1 (name) VALUES ("ashley");
INSERT INTO example1 (name) VALUES ("jesus");
INSERT INTO example1 (name) VALUES ("velma");
INSERT INTO example1 (name) VALUES ("dawn");
INSERT INTO example1 (name) VALUES ("marvin");
INSERT INTO example1 (name) VALUES ("trey");
INSERT INTO example1 (name) VALUES ("tandy");

INSERT INTO example2 (name) VALUES ("jeff");
INSERT INTO example2 (name) VALUES ("hector");
INSERT INTO example2 (name) VALUES ("allison");
INSERT INTO example2 (name) VALUES ("marshall");
INSERT INTO example2 (name) VALUES ("leticia");
INSERT INTO example2 (name) VALUES ("carol");
INSERT INTO example2 (name) VALUES ("dave");
INSERT INTO example2 (name) VALUES ("phil");
INSERT INTO example2 (name) VALUES ("frank");
INSERT INTO example2 (name) VALUES ("emmy");
INSERT INTO example2 (name) VALUES ("sara");