create table lookout(
  id integer primary key,
  name text not null,
  query text not null,
  cron text not null
);
