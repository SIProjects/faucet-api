create table payout (
  txid varchar primary key,
  address varchar not null,
  amount varchar not null,
  is_mined boolean not null,
  inserted_at timestamp,
  updated_at timestamp
);

create index on payout (address, inserted_at);
