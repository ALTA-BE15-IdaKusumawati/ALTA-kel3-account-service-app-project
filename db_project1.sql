create database db_project1;

use db_project1;

create table users(
id varchar(50) primary key not null,
name varchar(50),
telepon int,
email varchar(50),
password varchar(25),
saldo decimal default 0
);

create table transfer(
user_id_pengirim varchar(50),
user_id_penerima varchar(50),
nominal decimal not null,
created_at datetime default current_timestamp,
constraint fk_user_id_pengirim foreign key (user_id_pengirim) 
references users(id) on update cascade on delete cascade,
constraint fk_user_id_penerima foreign key (user_id_penerima) 
references users(id) on update cascade on delete cascade
);

create table topup(
user_id varchar(50),
nominal decimal not null,
created_at datetime default current_timestamp,
constraint fk_user_id foreign key (user_id) 
references users(id) on update cascade on delete cascade
);

UPDATE users INNER JOIN transfer ON users.id="US0001" SET saldo=saldo-5000;

select * from users where id="US0001";