CREATE TABLE IF NOT EXISTS public.items (
   id           BIGSERIAL       PRIMARY KEY,
   name         VARCHAR(64)     NOT NULL,
   slot         VARCHAR(16)     NOT NULL,
   cost         BIGINT          NOT NULL,
   skill_bonus  INT             NOT NULL
);
