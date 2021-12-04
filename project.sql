CREATE DATABASE IF NOT EXISTS software DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
/* CREATE DATABASE project; */

use software;
DROP TABLE IF EXISTS project;
DROP TABLE IF EXISTS user;

CREATE TABLE IF NOT EXISTS user
(
    id        BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    user_name VARCHAR(100) NOT NULL,
    password  CHAR(32)     NOT NULL
);



CREATE TABLE IF NOT EXISTS project
(
    id                   BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    project_name         VARCHAR(100)     NOT NULL,
    user_id              BIGINT UNSIGNED  NOT NULL,
    operating_system     TINYINT UNSIGNED NOT NULL,
    programming_language TINYINT UNSIGNED NOT NULL,
    natural_language     TINYINT UNSIGNED NOT NULL,
    topic                TINYINT UNSIGNED NOT NULL,
    tags                 VARCHAR(20),
    code_addr            CHAR(24),
    binary_addr          CHAR(24),
    license              VARCHAR(50),
    update_time          TIMESTAMP        NOT NULL,
    project_description  TEXT,
    FULLTEXT (project_name),
    foreign key (user_id) references user (id)
        on update cascade
        on delete cascade
);

