CREATE TABLE schools (
  ID INT AUTO_INCREMENT NOT NULL,
  name VARCHAR(40),
  total_students Int,
  PRIMARY KEY (ID)
);
CREATE TABLE students (
  ID INT(13) NOT NULL,
  name VARCHAR(40),
  age INT,
  grade INT,
  PRIMARY KEY (ID)
);