SELECT * FROM pg_create_logical_replication_slot('test_logical_slot_6', 'pgoutput');

SELECT * FROM pg_replication_slots;

SELECT pg_drop_replication_slot('test_logical_slot_5');

CREATE PUBLICATION my_publication FOR ALL TABLES;

CREATE PUBLICATION my_publication FOR TABLE public.users;


SELECT pubname AS publication_name,
       puballtables AS replicates_all_tables,
       pubinsert,
       pubupdate,
       pubdelete,
       pubtruncate
FROM pg_publication;

select * from pg_publication_tables;


SELECT slot_name,
       pg_size_pretty(pg_current_wal_lsn() - restart_lsn) AS retained_wal_size
FROM pg_replication_slots;

SELECT
    slot_name,
    slot_type,
    pg_size_pretty(pg_current_wal_lsn() - restart_lsn) AS retained_wal_size,
    active,
    wal_status,
    plugin AS output_plugin,
    database AS active_database,
    restart_lsn,
    confirmed_flush_lsn
FROM
    pg_replication_slots;

select * from pg_replication_slots;

SELECT name, setting, unit, context, short_desc
FROM pg_settings;
-- WHERE name = 'max_replication_slots';

SELECT *
FROM pg_settings;

