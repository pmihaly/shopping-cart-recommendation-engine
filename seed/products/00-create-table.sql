create extension if not exists "uuid-ossp";

create table product (
    id uuid primary key default uuid_generate_v4 (),
    name varchar(120) not null,
    description text not null,
    category varchar(80)[] not null,
    image_url text not null,
    price integer not null
  );
