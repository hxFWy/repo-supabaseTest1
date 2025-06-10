BEGIN;
    DELETE FROM public.items;
    INSERT INTO
        public.items (name, slot, cost, skill_bonus)
        VALUES ('Short Sword', 'weapon', 100, 5)
        ON CONFLICT (name) DO NOTHING;
COMMIT;
