-- migrate:up
create TABLE TodoItem(
    id integer PRIMARY KEY,
    description varchar (255) NOT NULL,
    completed boolean
    created_at TIMESTAMP
);

drop table TodoItem;
-- migrate:down
