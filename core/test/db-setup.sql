create table test_table
(
    id             int           not null primary key,
    text_value     varchar(255)  not null,
    number_value   int
);

insert into test_table (id, text_value, number_value) values(0, 'works', 1);
