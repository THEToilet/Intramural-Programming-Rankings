DROP DATABASE IF EXISTS user_db;
CREATE DATABASE user_db;
USE user_db;

CREATE TABLE users (
  id integer not null auto_increment,
  name text,
  accepted_count integer,
  accepted_count_rank integer,
  rated_point_sum integer,
  rated_point_sum_rank integer,
  created_time datetime default current_timestamp,
  updated_time datetime not null default current_timestamp on update current_timestamp,
  primary key (id)
)DEFAULT CHARACTER SET=utf8;

-- INSERT INTO users (name, accepted_count) VALUES ("太郎",3),("花子",34),("令和",324);
