insert into mst_customer(name, no_hp, address) values ('Ridho', '08131282872', 'Riau');

insert into mst_services(service, price) values ('Cuci', 5000);
insert into mst_services(service, price) values ('Setrika', 5000);
insert into mst_services(service, price) values ('Bedcover', 10000);

insert into trn_laundry(unit, amount, date_in, date_out, id_customer, id_service) values ('Kg', '1', '2023-12-01', '2023-12-02', 1, 1);
insert into trn_laundry(unit, amount, date_in, date_out, id_customer, id_service) values ('Piece', '1', '2023-12-01', '2023-12-02', 1, 2);