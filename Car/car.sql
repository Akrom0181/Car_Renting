CREATE TABLE IF NOT EXISTS cars (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name Varchar(50) NOT NULL,
    year INTEGER NOT NULL DEFAULT 2020,
    brand Varchar(20) NOT NULL,
    model Varchar(30) NOT NULL,
    hourse_power INTEGER DEFAULT 0,
    colour VARCHAR(20) NOT NULL DEFAULT 'black',
    engine_cap DECIMAL(10,2) NOT NULL DEFAULT 1.0,
    year INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS customers(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    gmail VARCHAR(50) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS orders (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  car_id uuid REFERENCES cars(id),
  customer_id uuid REFERENCES customers(id),
  from_date timestamp NOT NULL,
  to_date timestamp NOT NULL,
  status varchar(60) NOT NULL CHECK(status IN ('success', 'in process', 'failed')) DEFAULT 'in process',
  payment_status BOOLEAN DEFAULT TRUE, 
  amount DECIMAL(10,2),
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp DEFAULT NOW()
);


