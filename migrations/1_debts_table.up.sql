CREATE TABLE IF NOT EXISTS "debts" (
    "id" serial not null PRIMARY key,
    "first_name" varchar(50) not null,
    "last_name" varchar(50) not null,
    "phone_number" varchar(30) not null UNIQUE,
    "additional_phone_number" varchar(30) UNIQUE,
    "addres_work" varchar(100) not null,
    "seller_fullname" varchar(50) check ("seller_fullname" in ('zohid saidov', 'ali ismoilov')),
    "created_at" TIMESTAMP DEFAULT current_timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp
);