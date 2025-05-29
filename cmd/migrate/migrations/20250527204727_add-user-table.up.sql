CREATE TABLE IF NOT EXISTS public.users (
  id           BIGSERIAL       PRIMARY KEY,
  created_at   TIMESTAMPTZ     NOT NULL DEFAULT now(),
  username     VARCHAR         NOT NULL UNIQUE,
  password     VARCHAR         NOT NULL
);
