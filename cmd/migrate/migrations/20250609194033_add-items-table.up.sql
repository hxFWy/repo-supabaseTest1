CREATE TABLE IF NOT EXISTS public.items (
   id           BIGSERIAL       PRIMARY KEY,
   NAME         VARCHAR(64)     NOT NULL UNIQUE,
   slot         VARCHAR(16)     NOT NULL,
   cost         BIGINT          NOT NULL,
   skill_bonus  INT             NOT NULL,
   created_at   TIMESTAMPTZ     NOT NULL    DEFAULT now()
);
