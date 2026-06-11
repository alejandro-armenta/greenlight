
create table if not exists users(

  id bigint primary key generated always as identity,
  
  created_at timestamp(0) with time zone not null default now(),

  name text not null,

  email citext unique not null,

  password_hash bytea not null,

  activated bool not null,

  version integer not null default 1

);

