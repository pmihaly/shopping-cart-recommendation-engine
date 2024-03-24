create extension if not exists "uuid-ossp";

create extension if not exists unaccent;

create table product (
  id uuid primary key default uuid_generate_v4 (),
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
