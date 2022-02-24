USE test;
DROP TABLE IF EXISTS users;
CREATE TABLE users(id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255), phone VARCHAR(30), active BOOLEAN);
INSERT INTO users(name, phone, active) VALUES('Loren Uptown Fisher', "+1363333671000", false),('John Lipxten So Jr', "+475525012025520", true),('Helen Matthew', "+98325622522523", true);