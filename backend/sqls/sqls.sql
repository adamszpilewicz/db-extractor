create table public.users3 (
    id serial primary key,
    username varchar(255) not null,
    password varchar(255) not null,
    email varchar(255) not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

insert into public.users3 (username, password, email) values ('admin', 'admin', 'admin@localhost');

select *
from public.users3;

SELECT COLUMN_NAME, DATA_TYPE, is_nullable
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_NAME = 'users_with_id';

SELECT *
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_NAME = 'users_with_id';

SELECT
    C.COLUMN_NAME,
    C.DATA_TYPE,
    CASE WHEN KCU.COLUMN_NAME IS NOT NULL THEN 'YES' ELSE 'NO' END AS IS_PRIMARY_KEY
FROM
    INFORMATION_SCHEMA.COLUMNS C
        LEFT JOIN
    INFORMATION_SCHEMA.KEY_COLUMN_USAGE KCU
    ON C.COLUMN_NAME = KCU.COLUMN_NAME AND C.TABLE_NAME = KCU.TABLE_NAME
        LEFT JOIN
    INFORMATION_SCHEMA.TABLE_CONSTRAINTS TC
    ON KCU.CONSTRAINT_NAME = TC.CONSTRAINT_NAME AND KCU.TABLE_NAME = TC.TABLE_NAME
WHERE
        C.TABLE_NAME = 'test'
  AND (TC.CONSTRAINT_TYPE = 'PRIMARY KEY' OR TC.CONSTRAINT_TYPE IS NULL);


CREATE VIEW users_with_id_first_5 AS
SELECT
    id,
    email,
    password
FROM public.users
LIMIT 5;

select * from users_with_id_first_5;


CREATE VIEW users_with_id_first_6 AS
SELECT
    id,
    name,
    age,
    email,
    password
FROM users_with_id
LIMIT 6;

select * from users_with_id_first_6;

SELECT viewname AS view_name,
       viewowner AS owner,
       definition
FROM pg_views
WHERE schemaname NOT IN ('pg_catalog', 'information_schema');


SELECT users_with_id.id,
       users_with_id.name,
       users_with_id.age,
       users_with_id.email,
       users_with_id.password
FROM users_with_id
LIMIT 6;

SELECT
    datname AS database_name,
    pg_encoding_to_char(encoding) AS encoding,
    datcollate AS collate,
    datctype AS ctype,
    datistemplate AS is_template,
    datallowconn AS allow_connections,
    datconnlimit AS connection_limit,
    pg_size_pretty(pg_database_size(datname)) AS size,
    datfrozenxid AS frozen_transaction_id,
    datminmxid AS minimum_multixact_id,
    datlastsysoid AS last_system_oid,
    pg_get_userbyid(datdba) AS owner
FROM pg_database
WHERE datname = current_database();

SELECT
    version() AS postgres_version,
    current_database() AS current_database,
    current_user AS current_user,
    inet_server_addr() AS server_ip,
    inet_server_port() AS server_port,
    pg_postmaster_start_time() AS start_time,
    pg_conf_load_time() AS last_config_load,
    pg_is_in_recovery() AS in_recovery,
    pg_last_wal_receive_lsn() AS last_wal_received,
    pg_last_wal_replay_lsn() AS last_wal_replayed,
    pg_current_wal_lsn() AS current_wal,
    pg_database_size(current_database()) AS current_db_size
FROM
    pg_database
WHERE
        datname = current_database();

WITH database_stats AS (
    SELECT
        datname AS database_name,
        numbackends AS active_connections,
        xact_commit AS total_transactions_committed,
        xact_rollback AS total_transactions_rolled_back,
        blks_read AS blocks_read,
        blks_hit AS cache_hits,
        tup_returned AS rows_returned_by_queries,
        tup_fetched AS rows_fetched_by_queries,
        tup_inserted AS rows_inserted,
        tup_updated AS rows_updated,
        tup_deleted AS rows_deleted,
        conflicts AS conflicts_detected
    FROM pg_stat_database WHERE datname = current_database()
),

     session_stats AS (
         SELECT
                     count(*) FILTER (WHERE state = 'active') AS active_sessions,
                     count(*) FILTER (WHERE state = 'idle') AS idle_sessions,
                     count(*) FILTER (WHERE state = 'idle in transaction') AS idle_in_transaction_sessions,
                     count(*) FILTER (WHERE state = 'fastpath function call') AS fastpath_function_call_sessions
         FROM pg_stat_activity WHERE datname = current_database()
     )

SELECT
    ds.*,
    ss.active_sessions,
    ss.idle_sessions,
    ss.idle_in_transaction_sessions,
    ss.fastpath_function_call_sessions
FROM database_stats ds, session_stats ss;


CREATE OR REPLACE PROCEDURE public.generate_test_data(tablesamount integer DEFAULT 1, rowspertable integer DEFAULT 100)
    LANGUAGE plpgsql
AS $procedure$
DECLARE
    random_table_names text [];
    random_table text;
BEGIN
    SET search_path = "public";
    -- generate tables names
    FOR counter IN 1..tablesamount LOOP
            random_table := FORMAT('"random_table"', counter);
            random_table_names := array_append(random_table_names, random_table);
        END LOOP;
    -- create tables
    FOR counter IN 1..tablesamount LOOP
            EXECUTE format('create table IF NOT EXISTS %s(
    				id serial,
    				time_ff time ,
    				date_ff date,
    				bool_ff bool,
    				small_int_ff smallint default 1,
    				bigint_ff bigint default 2,
    				integer_ff integer default 3,
    				text_ff text default ''text'',
    				byte_ff bytea default ''\000'',
    				test_id text default ''test_id'',
    				boolean_ff smallint default 0,
    				real_ff real default 1.5,
    				uuid_ff uuid default ''A0EEBC99-9C0B-4EF8-BB6D-6BB9BD380A11'',
    				timetz_ff timestamptz default now()
    			);', random_table_names [counter]);
            EXECUTE format('alter table %s replica identity full', random_table_names [counter]);

            -- 		CREATE TABLE partition_first PARTITION OF chunks_partitioned_table FOR VALUES FROM (1) TO (300000);
-- 		CREATE TABLE partition_second PARTITION OF chunks_partitioned_table FOR VALUES FROM (300000) TO (600000);
-- 		CREATE TABLE partition_third PARTITION OF chunks_partitioned_table FOR VALUES FROM (600000) TO (1100000);


        END LOOP;
    -- 	insert elements to tables
    FOR counter_tables IN 1..tablesamount LOOP
            FOR counter_rows IN 1..rowspertable LOOP
                    EXECUTE format('INSERT INTO %s (date_ff,time_ff,test_id,bool_ff)
    							VALUES (''2022-06-08'',''2022-06-08 12:31:03.945145'',''customers_check_8660184831103359040_all_types'', true)', random_table_names [counter_tables]);
                END LOOP;
        END LOOP;
END;
$procedure$


SELECT
    n.nspname AS schema_name,
    p.proname AS procedure_name,
    pg_get_functiondef(p.oid) AS definition,
    r.rolname AS created_by
FROM
    pg_proc p
        JOIN
    pg_namespace n ON p.pronamespace = n.oid
        JOIN
    pg_roles r ON p.proowner = r.oid
WHERE
        n.nspname NOT IN ('pg_catalog', 'information_schema') -- Exclude system schemas
ORDER BY
    schema_name,
    procedure_name;

