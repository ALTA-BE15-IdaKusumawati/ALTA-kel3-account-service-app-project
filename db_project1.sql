CREATE DATABASE db_project1;

USE db_project1;

-- create tabel users
CREATE TABLE users(
id varchar(50) primary key,
name varchar(50),
telepon int,
email varchar(50),
password varchar(25),
saldo decimal default 0
);

CREATE TABLE topup(
id int primary key auto_increment,
user_id varchar(50),
nominal decimal not null,
created_at datetime default current_timestamp,
FOREIGN KEY (user_id) REFERENCES users(id) on update cascade on delete cascade
);

CREATE TABLE transfer(
id int primary key auto_increment,
user_id_pengirim varchar(50),
user_id_penerima varchar(50),
nominal decimal not null,
created_at datetime default current_timestamp,
FOREIGN KEY (user_id_pengirim) references users(id) on update cascade on delete cascade,
FOREIGN KEY (user_id_penerima) references users(id)	on update cascade on delete cascade
);

DROP TABLE users;
DROP TABLE saldo;
DROP TABLE topup;
DROP TABLE transfer;