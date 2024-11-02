CREATE DATABASE param_db;


CREATE TABLE IF NOT EXISTS testing_api_tab (
  id SERIAL NOT NULL,
  path varchar(256) NOT NULL,
  method SMALLINT NOT NULL,
  host varchar(128) NOT NULL,
  scheme varchar(8) NOT NULL,
  field TEXT NOT NULL,
  is_active SMALLINT NOT NULL DEFAULT 1,
  extra_data TEXT NOT NULL,
  create_time INT NOT NULL DEFAULT 0,
  update_time INT NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS testing_log_tab (
  id SERIAL NOT NULL,
  status SMALLINT NOT NULL DEFAULT 0,
  test_file_name varchar(64) NOT NULL,
  test_result TEXT NOT NULL,
  create_time INT NOT NULL DEFAULT 0,
  finish_time INT NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS testing_log_detail_tab (
  id SERIAL NOT NULL,
  api_id INT NOT NULL,
  log_id INT NOT NULL,
  create_time INT NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);