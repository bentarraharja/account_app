create database AccountServiceApp;
use AccountServiceApp;

-- Buat tabel User
create table User (
    UserID INT primary key auto_increment,
    Fullname VARCHAR(255),
    Email VARCHAR(255),
    Phone VARCHAR(15),
    Password VARCHAR(255),
    Balance INT,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp
    
);
-- Buat Tabel Transfer_Histories
create table Transfer_Histories (
    TransferID INT primary key auto_increment,
    UserID_Sender INT,
    UserID_Recipient INT,
    Amount INT,
    created_at timestamp default current_timestamp,
    foreign key (UserID_Sender) references User(UserID),
    foreign key (UserID_Recipient) references User(UserID)
);

-- Buat Tabel Top_Up_Histories
create  table Top_Up_Histories (
    UserID INT,
    Amount INT,
    created_at timestamp default current_timestamp,
    foreign key (UserID) references User(UserID)
);

desc user;
desc Top_Up_Histories;
desc Transfer_Histories;


-- Isi data pada tabel User
-- INSERT INTO User (Fullname, Email, Phone, Password, Balance)
-- VALUES 
--     ('firly', 'firly@example.com', '123456789', 'password123', 100.00),
--     ('dani', 'dani@example.com', '987654321', 'password123', 50.00);
-- 
--    
--    Isi data pada tabel Transfer_Histories
-- INSERT INTO Transfer_Histories (UserID_Sender, UserID_Recipient, Amount)
-- VALUES 
--     (1, 2, 3000000),
--     (2, 1, 2000000);
--   
-- 
-- Isi data pada tabel Top_Up_Histories
-- INSERT INTO Top_Up_Histories (UserID, Amount)
-- VALUES 
--     (1, 500000),
--     (2, 250000);
    
select * from `user` u ;
select * from top_up_histories tuh ;
select * from transfer_histories th ;