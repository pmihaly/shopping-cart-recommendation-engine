create extension if not exists unaccent;

create extension if not exists "uuid-ossp";

create table product (
  id varchar(32) primary key,
  name varchar(120) not null,
  description text not null,
  category varchar(80) [] not null,
  image_url text not null,
  price integer not null,
  name_search tsvector,
  description_search tsvector
);

create
or replace function populate_search_vectors () returns trigger as $$
BEGIN
  NEW.name_search := to_tsvector('simple', unaccent(NEW.name));
  NEW.description_search := to_tsvector('simple', unaccent(NEW.description));
  RETURN NEW;
END;
$$ language plpgsql;

create trigger tsvectorupdate before insert
or
update on product for each row
execute function populate_search_vectors ();

create index idx_product_search on product using gin (name_search);

create index idx_product_description_search on product using gin (description_search);

create table cart (id uuid primary key default gen_random_uuid ());

create table cart_items (
  cart_id uuid references cart (id),
  product_id varchar(32) references product (id),
  primary key (cart_id, product_id)
);
