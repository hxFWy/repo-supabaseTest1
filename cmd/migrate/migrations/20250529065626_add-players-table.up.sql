CREATE TABLE IF NOT EXISTS public.players (
  user_id      BIGSERIAL          PRIMARY KEY REFERENCES public.users(id),
  money        DOUBLE PRECISION   NOT NULL DEFAULT 0,
  position     VARCHAR            NOT NULL,
  stamina      INT                NOT NULL DEFAULT 100,
  skill        INT                NOT NULL,
  created_at   TIMESTAMPTZ        NOT NULL DEFAULT now()
);
