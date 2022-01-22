
--
-- Creating Database`
--

create database tododb;
use tododb;


--
-- Table structure for table `User`
--
create table user(
  id int auto_increment not null,
  first_name varchar(50) not null,
  last_name varchar(50) not null,
   primary key(id)
  

);


--
-- Table structure for table `Task`
--

create table task(
  id int auto_increment not null,
  user_id  int not null,
  task_name  varchar(100) not null,
  description text,
  status varchar(50) not null,
  due_date datetime NOT NULL DEFAULT current_timestamp(),
  primary key(id)
);
 