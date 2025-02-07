CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE IF NOT EXISTS usuarios(
    id INT PRIMARY KEY AUTO_INCREMENT,
    nome VARCHAR(50) not null,
    username VARCHAR(50) not null unique,
    email VARCHAR(50) not null unique,
    senha VARCHAR(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id) 
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    PRIMARY KEY(usuario_id, seguidor_id)
) ENGINE=INNODB;

CREATE TABLE publicacoes(
    id INT PRIMARY KEY AUTO_INCREMENT,
    titulo VARCHAR(50) NOT NULL,
    conteudo VARCHAR(400) NOT NULL,
    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id) 
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    likes int default 0,
    criadaEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;