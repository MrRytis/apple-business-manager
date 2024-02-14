create table goose_db_version
(
    id         serial primary key,
    version_id bigint  not null,
    is_applied boolean not null,
    tstamp     timestamp default now()
);

insert into goose_db_version (version_id, is_applied) values (0, true);
