create database account_service_app;
use account_service_app;

-- Buat tabel accounts
create table accounts (
    id INT primary key auto_increment,
    full_name VARCHAR(255),
    address VARCHAR(255),
    phone VARCHAR(15),
    email VARCHAR(255),
    password VARCHAR(255),
    balance INT default 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp
);
-- Buat tabel transfers
create table transfers (
    id INT primary key auto_increment,
    account_id_sender INT,
    account_id_receiver INT,
    amount INT,
    created_at timestamp default current_timestamp,
    constraint fk_transfers_accounts_s foreign key (account_id_sender) references accounts(id),
    constraint fk_transfers_accounts_r foreign key (account_id_receiver) references accounts(id)
);

-- Buat Tabel top_ups
create  table top_ups (
    id INT primary key auto_increment,
    account_id INT,
    amount INT,
    created_at timestamp default current_timestamp,
    constraint fk_top_ups_accounts_ foreign key (account_id) references accounts(id)
);

-- Isi data pada tabel accounts
-- INSERT INTO accounts (full_name, address, phone, email, password, balance)
-- VALUES 
--     ('firly', 'jakarta', '123456789', 'firly@example.com', 'password123', 100000),
--     ('dani', 'bandung', '987654321', 'dani@example.com', 'password123', 50000);
-- 
--    
-- Isi data pada tabel transfers
-- INSERT INTO transfers (account_id_sender, account_id_receiver, amount)
-- VALUES 
--     (1, 2, 3000000),
--     (2, 1, 2000000);
--   
-- 
-- Isi data pada tabel top_ups
-- INSERT INTO top_ups (account_id, amount)
-- VALUES 
--     (1, 500000),
--     (2, 250000);
    
select * from accounts;
select * from transfers;
select * from top_ups;