CREATE TABLE IF NOT EXISTS tasks (
    id integer GENERATED ALWAYS AS IDENTITY,
    title text,
    completed integer,
    created_at text
);