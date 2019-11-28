CREATE TABLE authors (
    ID SERIAL PRIMARY KEY,
    NAME TEXT NOT NULL,
    AGE TEXT
);

CREATE TABLE books (
    ID SERIAL PRIMARY KEY,
    NAME TEXT NOT NULL,
    GENRE TEXT NOT NULL,
    AUTHOR_ID INTEGER REFERENCES authors (ID)
);


insert into authors (NAME, Age) values ('AAA', 42);
insert into authors (NAME, Age) values ('BBB', 52);
insert into authors (NAME, Age) values ('CCC', 62);



insert into books (NAME, GENRE, AUTHOR_ID) values ('First Book', 'Fantasy', 1);
insert into books (NAME, GENRE, AUTHOR_ID) values ('Second Book', 'Tech', 2);
insert into books (NAME, GENRE, AUTHOR_ID) values ('Third Book', 'Music', 1);
