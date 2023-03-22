PRAGMA foreign_keys = ON;

create table lookout(
  id integer primary key,
  name text not null,
  query text not null,
  cron text not null
);

create table review_rule(
  id integer primary key,
  lookout_id integer not null, 
  column_name text not null,
  row_index integer not null,
  exact_value text,

  FOREIGN KEY(lookout_id) REFERENCES lookout(id)
);



-- temp test data
insert into lookout (name, query, cron) values ('test', 'select count(*) as "count" from lookout', '@every 1s');
insert into review_rule (lookout_id, column_name, row_index, exact_value) values (1, 'count', 0, '1');
