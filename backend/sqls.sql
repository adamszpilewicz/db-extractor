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
        C.TABLE_NAME = 'users_with_id'
  AND (TC.CONSTRAINT_TYPE = 'PRIMARY KEY' OR TC.CONSTRAINT_TYPE IS NULL);


CREATE VIEW users_with_id_first_5 AS
SELECT
    id,
    name,
    age,
    email,
    password
FROM users_with_id
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
