CREATE TABLE IF NOT EXISTS public.player_items (
   player_id    INT             REFERENCES public.players(user_id),
   item_id      INT             REFERENCES public.items(id),
   equipped     BOOLEAN         DEFAULT FALSE,
   PRIMARY KEY(player_id, item_id)
);
