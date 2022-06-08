drop DATABASE if exists task_manager ;

CREATE DATABASE task_manager;

USE task_manager; ;

CREATE TABLE IF NOT EXISTS projects (
 		id INT(11) NOT NULL AUTO_INCREMENT,
 		name VARCHAR(255) NOT NULL,
 		PRIMARY KEY (id)
 	);

CREATE TABLE IF NOT EXISTS tasks (
 		id INT(11) NOT NULL AUTO_INCREMENT,
 		description VARCHAR(255) NOT NULL,
 		assignee VARCHAR(255) NOT NULL,
 		status VARCHAR(255) NOT NULL,
 		project_id INT(11) NOT NULL,
 		PRIMARY KEY (id),
 		FOREIGN KEY (project_id) REFERENCES projects(id)
 	)


CREATE TABLE IF NOT EXISTS users (
    id INT(11) NOT NULL AUTO_INCREMENT,
 		name VARCHAR(255) NOT NULL,
 		role VARCHAR(255) NOT NULL,
 		PRIMARY KEY (id)
)