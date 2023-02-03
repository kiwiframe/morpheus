-- +goose Up
DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA users;

CREATE TYPE users.privilege AS ENUM ('user', 'moderator', 'admin');
CREATE TYPE users.restriction AS ENUM ('none', 'communication', 'full');

CREATE TABLE users.accounts
(
    id                 uuid PRIMARY KEY           DEFAULT gen_random_uuid(),
    username           varchar(16)       NOT NULL UNIQUE,
    email              varchar(254)      NOT NULL UNIQUE,
    hashed_password    text              NOT NULL,
    created_at         timestamptz       NOT NULL DEFAULT now(),
    updated_at         timestamptz       NOT NULL DEFAULT now(),
    last_logged_in     timestamptz,
    privilege_level    users.privilege   NOT NULL DEFAULT 'user',
    restriction_level  users.restriction NOT NULL DEFAULT 'none',
    restriction_expiry timestamptz,
    verified           boolean           NOT NULL DEFAULT false
);

CREATE TABLE users.verifications
(
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    uuid        NOT NULL REFERENCES users.accounts ON DELETE CASCADE,
    token      varchar(32) NOT NULL UNIQUE,
    expires_at timestamptz NOT NULL
);

-- +goose Down
DROP SCHEMA IF EXISTS users CASCADE;
