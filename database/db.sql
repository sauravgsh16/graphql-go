DROP TABLE books;
DROP TABLE authors;

CREATE TABLE IF NOT EXISTS authors (
    ID SERIAL PRIMARY KEY,
    NAME TEXT NOT NULL,
    AGE TEXT
);

CREATE TABLE IF NOT EXISTS books (
    ID SERIAL PRIMARY KEY,
    NAME TEXT NOT NULL,
    GENRE TEXT NOT NULL,
    AUTHOR_ID INTEGER,
    FOREIGN KEY (AUTHOR_ID) REFERENCES authors (ID)
);

insert into authors (NAME, Age) values ('John', 42);
insert into authors (NAME, Age) values ('Dave', 52);
insert into authors (NAME, Age) values ('Frank', 62);

insert into books (NAME, GENRE, AUTHOR_ID) values ('First Book', 'Fantasy', 1);
insert into books (NAME, GENRE, AUTHOR_ID) values ('Second Book', 'Tech', 2);
insert into books (NAME, GENRE, AUTHOR_ID) values ('Third Book', 'Music', 1);
