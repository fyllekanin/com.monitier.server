CREATE TABLE IF NOT EXISTS pings (
     ID INTEGER PRIMARY KEY AUTOINCREMENT,
     serviceName varchar(255) NOT NULL,
     responseTime INTEGER,
     isUp INTEGER
);