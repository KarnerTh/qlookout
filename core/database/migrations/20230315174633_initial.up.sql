PRAGMA foreign_keys = ON;

create table lookout(
  id integer primary key,
  name text not null,
  query text not null,
  cron text not null,
  notify_local boolean not null,
  notify_mail boolean not null
);

create table review_rule(
  id integer primary key,
  lookout_id integer not null, 
  column_name text not null,
  column_type text check(column_type in ('text', 'int', 'float')) not null,
  row_index integer not null,
  exact_value text,
  greater_than text,
  less_than text,
  should_be_null boolean,

  FOREIGN KEY(lookout_id) REFERENCES lookout(id)
);


-- TODO: remove
-- temp test data
insert into lookout (name, query, cron, notify_local, notify_mail) values ('test', 'select count(*) as "count" from lookout', '@every 1s', true, false);
insert into review_rule (lookout_id, column_name, column_type, row_index, exact_value) values (1, 'count', 'int', 0, '1');
