CREATE DATABASE IF NOT EXISTS p8-sa;
USE p8-sa;

DROP TABLE IF EXISTS Reporte;
DROP TABLE IF EXISTS Estudiante;

CREATE TABLE Estudiante(
    ID_estudiante int NOT NULL AUTO_INCREMENT,
    carnet int NOT NULL,
    nombre varchar(64) NOT NULL,
    curso varchar(128) NOT NULL,

    PRIMARY KEY(ID_estudiante)
);
