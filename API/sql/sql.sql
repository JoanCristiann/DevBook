CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE IF NOT EXISTS usuarios(
    id INT PRIMARY KEY AUTO_INCREMENT,
    nome VARCHAR(50) not null,
    username VARCHAR(50) not null unique,
    email VARCHAR(50) not null unique,
    senha VARCHAR(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;