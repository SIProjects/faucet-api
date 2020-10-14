create table payout (
  txid varchar,
  address varchar not null,
  amount numeric(12,8) not null,
  is_mined boolean not null,
  inserted_at timestamp,
  updated_at timestamp,
  primary key (txid, address)
);

create index on payout(txid)

create index on payout (updated_at)
