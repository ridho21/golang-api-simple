CREATE TABLE MST_CUSTOMER(
	id SERIAL PRIMARY KEY,
	name varchar(50) not null,
	no_hp varchar(15),
	address varchar(30)
);

CREATE TABLE MST_SERVICES(
	id SERIAL PRIMARY KEY,
	service varchar(30) not null,
	price int
);

CREATE TABLE TRN_LAUNDRY(
	id SERIAL PRIMARY KEY,
	unit varchar(10) not null,
	amount int not null,
	date_in date,
	date_out date,
	id_customer int,
	id_service int,
	constraint fk_customer
	foreign key (id_customer)
	references MST_CUSTOMER(id),
	constraint fk_service
	foreign key (id_service)
	references MST_SERVICES(id)
);