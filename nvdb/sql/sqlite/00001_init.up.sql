create table files
(
    id              INTEGER                           not null
        constraint files_pk
            primary key autoincrement,
    name            TEXT                              not null,
    extension       TEXT,
    base_path       TEXT                              not null,
    full_path       TEXT                              not null,
    size            INTEGER                           not null,
    library_path_id INTEGER                           not null,
    modified_at     INTEGER                           not null,
    created_at      INTEGER default CURRENT_TIMESTAMP not null,
    updated_at      INTEGER default CURRENT_TIMESTAMP not null
);

create unique index files_full_path_unique_index
    on files (full_path);

create table library
(
    id              INTEGER                           not null
        constraint files_pk
            primary key autoincrement,
    name            TEXT                              not null,
    created_at      INTEGER default CURRENT_TIMESTAMP not null,
    updated_at      INTEGER default CURRENT_TIMESTAMP not null
);

create table library_path
(
    id         INTEGER                           not null
        constraint files_pk
            primary key autoincrement,
    library_id INTEGER                           not null,
    path       TEXT                              not null,
    created_at INTEGER default CURRENT_TIMESTAMP not null,
    updated_at INTEGER default CURRENT_TIMESTAMP not null
);
INSERT INTO library (name) VALUES ('Music');
INSERT INTO library_path (library_id, path)
VALUES ('1', '/home/tbruno/Music');

create table media_item
(
    id               INTEGER                           not null
        constraint files_pk
            primary key autoincrement,
    library_path_id  INTEGER                           not null,
    path             TEXT                              not null,
    mime_type        TEXT,
    file_modified_at INTEGER,
    created_at       INTEGER default CURRENT_TIMESTAMP not null,
    updated_at       INTEGER default CURRENT_TIMESTAMP not null
);

create table media_item_stream
(
    id            INTEGER not null
        constraint files_pk
            primary key autoincrement,
    media_item_id INTEGER not null,
    stream_index  INTEGER not null,
    type          TEXT    not null,
    codec         TEXT    not null,
    language      TEXT,
    height        INTEGER,
    width         INTEGER
);