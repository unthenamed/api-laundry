INSERT INTO customers (name, phone, address) VALUES
('John Doe', '081234567890', 'Jl. Melati No. 1'),
('Jane Smith', '082345678901', 'Jl. Mawar No. 2'),
('Alice Johnson', '083456789012', 'Jl. Anggrek No. 3'),
('Bob Brown', '084567890123', 'Jl. Kenanga No. 4'),
('Charlie Davis', '085678901234', 'Jl. Dahlia No. 5'),
('David Lee', '086789012345', 'Jl. Cempaka No. 6'),
('Eva Williams', '087890123456', 'Jl. Kamboja No. 7'),
('Frank Miller', '088901234567', 'Jl. Tulip No. 8'),
('Grace Wilson', '089012345678', 'Jl. Melur No. 9'),
('Hannah Taylor', '080123456789', 'Jl. Soka No. 10');

INSERT INTO employees (name, phone, address) VALUES
('Mark Smith', '081100000001', 'Jl. Mawar No. 1'),
('Lucy Brown', '081100000002', 'Jl. Anggrek No. 2'),
('Tom White', '081100000003', 'Jl. Melati No. 3'),
('Emma Green', '081100000004', 'Jl. Kenanga No. 4'),
('James Harris', '081100000005', 'Jl. Dahlia No. 5');

INSERT INTO products (name, price, unit) VALUES
('Cuci Kering', 5000, '1kg'),
('Cuci Setrika', 7000, '1kg'),
('Cuci Kilat', 10000, '1kg'),
('Cuci Sepatu', 20000, '1pair'),
('Cuci Boneka', 25000, '1piece');

INSERT INTO bills (bill_date, entry_date, finish_date, employee_id, customer_id) VALUES
('2024-12-01', '2024-12-01', '2024-12-03', 1, 1),
('2024-12-02', '2024-12-02', '2024-12-04', 2, 2),
('2024-12-03', '2024-12-03', '2024-12-05', 3, 3),
('2024-12-04', '2024-12-04', '2024-12-06', 4, 4),
('2024-12-05', '2024-12-05', '2024-12-07', 5, 5);

INSERT INTO details (bill_id, product_id, product_price, qty) VALUES
(1, 1, 5000, 5),
(1, 2, 7000, 2),
(2, 2, 7000, 3),
(2, 3, 10000, 1),
(3, 1, 5000, 10),
(3, 4, 20000, 2),
(4, 5, 25000, 1),
(5, 1, 5000, 5),
(5, 3, 10000, 2),
(5, 2, 7000, 1);
