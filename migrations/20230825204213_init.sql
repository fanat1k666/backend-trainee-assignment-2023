-- +goose Up

-- +goose StatementBegin
create table segment
(
    id   serial primary key,
    name varchar(255) NOT NULL,
    CONSTRAINT segment_unique UNIQUE (name)
);
-- +goose StatementEnd

-- +goose StatementBegin
create table "user"
(
    id      serial primary key,
    user_id int NOT NULL,
    CONSTRAINT user_unique UNIQUE (user_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
create table user_segment
(
    id             serial primary key,
    entry_time     timestamp NOT NULL DEFAULT NOW(),
    exit_time      timestamp,
    plan_exit_time timestamp,
    user_id        int       NOT NULL,
    segment_id     int       NOT NULL,
    constraint user_fk foreign key (user_id) references "user" on update cascade on delete cascade,
    constraint segment_fk foreign key (segment_id) references segment on update cascade on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user_segment;
-- +goose StatementEnd

-- +goose StatementBegin
drop table segment;
-- +goose StatementEnd

-- +goose StatementBegin
drop table "user";
-- +goose StatementEnd

