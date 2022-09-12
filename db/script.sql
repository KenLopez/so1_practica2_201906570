USE practica2;
DROP TABLE IF EXISTS PROCESO;
DROP TABLE IF EXISTS LOG;


CREATE TABLE IF NOT EXISTS LOG (
	id INT PRIMARY KEY AUTO_INCREMENT,
    fecha datetime NOT NULL,
    cpu FLOAT NOT NULL,
    ram FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS PROCESO (
	id INT PRIMARY KEY AUTO_INCREMENT,
	pid INT NOT NULL,
    nombre VARCHAR(50) NOT NULL,
    usuario INT NOT NULL,
    estado VARCHAR(50) NOT NULL,
    ram FLOAT NOT NULL,
    padre INT NULL,
    log INT NOT NULL,
    FOREIGN KEY (log) REFERENCES LOG(id) ON DELETE CASCADE,
    FOREIGN KEY (padre) REFERENCES PROCESO(id) ON DELETE CASCADE
);
