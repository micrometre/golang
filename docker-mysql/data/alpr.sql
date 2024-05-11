CREATE TABLE items (
            id INTEGER PRIMARY KEY,
            plate TEXT,
            uuid TEXT,
            img TEXT
          );

DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);


INSERT INTO items VALUES(1,'LM9KT0','MicrometreUK-cam1-1700315538650','http://localhost:3000/images/MicrometreUK-cam1-1700317660524.jpg  ');
INSERT INTO items VALUES(2,'LM09KT0','MicrometreUK-cam1-1700315538777','http://localhost:3000/images/MicrometreUK-cam1-1700317660524.jpg  ');
INSERT INTO items VALUES(3,'LM9KT0','MicrometreUK-cam1-1700315538882','http://localhost:3000/images/MicrometreUK-cam1-1700317660524.jpg  ');
INSERT INTO items VALUES(4,'LM9KT0','MicrometreUK-cam1-1700315538989','http://localhost:3000/images/MicrometreUK-cam1-1700317660524.jpg  ');
INSERT INTO items VALUES(1024,'73D44','MicrometreUK-cam1-1700317669928','http://localhost:3000/images/MicrometreUK-cam1-1700317670474.jpg  ');
INSERT INTO items VALUES(1025,'73D44','MicrometreUK-cam1-1700317669928','http://localhost:3000/images/MicrometreUK-cam1-1700317670367.jpg  ');
INSERT INTO items VALUES(1026,'73D44','MicrometreUK-cam1-1700317669928','http://localhost:3000/images/MicrometreUK-cam1-1700317660755.jpg  ');
INSERT INTO items VALUES(1027,'73D44','MicrometreUK-cam1-1700317669928','http://localhost:3000/images/MicrometreUK-cam1-1700317667823.jpg  ');
INSERT INTO items VALUES(1028,'73D44','MicrometreUK-cam1-1700317669928','http://localhost:3000/images/MicrometreUK-cam1-1700317660637.jpg  ');



INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);


COMMIT;
