CREATE DATABASE IF NOT EXISTS adventofcode;
USE adventofcode;

DROP TABLE IF EXISTS puzzleInputs;
CREATE TABLE puzzleInputs (round char(3) NOT NULL);
LOAD DATA INFILE '/tmp/puzzleinput.txt' INTO TABLE puzzleInputs;

DROP TABLE IF EXISTS rules1;
CREATE TABLE rules1 (game char(3) NOT NULL, score tinyint NOT NULL);
INSERT INTO rules1 (game, score) VALUES ("C X", 7), ("A Y", 8), ("B Z", 9), ("B X", 1), ("C Y", 2), ("A Z", 3), ("A X", 4), ("B Y", 5), ("C Z", 6);

SELECT "part 1:", SUM(score) FROM puzzleInputs p JOIN rules1 r ON p.round = r.game;

DROP TABLE IF EXISTS rules2;
CREATE TABLE rules2 (game char(3) NOT NULL, score tinyint NOT NULL);
INSERT INTO rules2 (game, score) VALUES ("A Z", 8), ("B Z", 9), ("C Z", 7), ("A X", 3), ("B X", 1), ("C X", 2), ("A Y", 4), ("B Y", 5), ("C Y", 6);

SELECT "part 2:", SUM(score) FROM puzzleInputs p JOIN rules2 r ON p.round = r.game;