create table articles
(
    id serial not null,
    title varchar not null,
    url varchar not null,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (id)
);
comment on table articles is 'mock.csv';
