package repo

import (
	"api-laundry/model"
	"database/sql"
	"fmt"
)

// Interface TransactionRepo mendefinisikan metode untuk operasi transaksi
type TransactionRepo interface {
	InsertTransaction(model.Transaction) (model.Response, error)
	GetAllTransaction(model.Transaction) ([]model.Response, error)
	GetTransactionById(id int) (model.Transaction, error)
}

// Struct Transaction mengimplementasikan interface TransactionRepo
type Transaction struct {
	db *sql.DB
}

// Metode InsertTransaction untuk menambahkan transaksi baru ke database
func (t *Transaction) InsertTransaction(mTransaction model.Transaction) (model.Response, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return model.Response{}, err
	}

	// Query untuk memasukkan data transaksi ke tabel transaction
	statment := `
		INSERT INTO bills
		(bill_date, entry_date, finish_date, employee_id, customer_id)
		VALUES($1, $2, $3, $4, $5) RETURNING id ,bill_date, entry_date, finish_date, employee_id, customer_id ;`

	// // Menjalankan query dan mendapatkan ID transaksi yang baru dimasukkan
	// fmt.Println(mTransaction)

	err = tx.QueryRow(
		statment,
		mTransaction.Bills.BillDate,
		mTransaction.Bills.EntryDate,
		mTransaction.Bills.FinishDate,
		mTransaction.Bills.EmployeeId,
		mTransaction.Bills.CustomerId,
	).Scan(
		&mTransaction.Response.Id,
		&mTransaction.Response.BillDate,
		&mTransaction.Response.EntryDate,
		&mTransaction.Response.FinishDate,
		&mTransaction.Response.Employee.Id,
		&mTransaction.Response.Customer.Id,
	)

	if err != nil {
		tx.Rollback()
		return model.Response{}, err
	}

	statement := `SELECT id, name, phone, address FROM employees WHERE id = $1;`
	err = tx.QueryRow(statement, mTransaction.Response.Employee.Id).Scan(
		&mTransaction.Response.Employee.Id,
		&mTransaction.Response.Employee.Name,
		&mTransaction.Response.Employee.PhoneNumber,
		&mTransaction.Response.Employee.Address,
	)
	if err != nil {
		return model.Response{}, err
	}

	statement = `SELECT id, name, phone, address FROM customers WHERE id = $1;`
	err = tx.QueryRow(statement, mTransaction.Response.Customer.Id).Scan(
		&mTransaction.Response.Customer.Id,
		&mTransaction.Response.Customer.Name,
		&mTransaction.Response.Customer.PhoneNumber,
		&mTransaction.Response.Customer.Address,
	)
	if err != nil {
		return model.Response{}, err
	}

	// Memasukkan detail transaksi ke tabel details
	for _, v := range mTransaction.Bills.BillDetails {
		// Mendapatkan harga produk dari tabel product
		var rDetail model.RDetails
		statment = `SELECT id, name, price, unit FROM products WHERE id = $1;`
		err = tx.QueryRow(statment, v.ProductId).Scan(
			&rDetail.Product.Id,
			&rDetail.Product.Name,
			&rDetail.Product.Price,
			&rDetail.Product.Unit,
		)

		if err != nil {
			tx.Rollback()
			return model.Response{}, err
		}

		// Query untuk memasukkan data detail transaksi ke tabel details
		statment = `
				INSERT INTO details
				(bill_id, product_id, product_price, qty)
				VALUES($1, $2, $3, $4) RETURNING id;`

		err = tx.QueryRow(
			statment,
			mTransaction.Response.Id,
			v.ProductId,
			rDetail.Product.Price,
			v.Qty,
		).Scan(&rDetail.Id)

		rDetail.BillId = mTransaction.Response.Id
		rDetail.Qty = v.Qty
		rDetail.ProductPrice = rDetail.Product.Price * v.Qty

		if err != nil {
			tx.Rollback()
			return model.Response{}, err
		}

		mTransaction.Response.BillDetails = append(mTransaction.Response.BillDetails, rDetail)
	}

	// Commit transaksi database
	err = tx.Commit()
	if err != nil {
		return model.Response{}, err
	}
	return mTransaction.Response, nil
}

// Metode GetAllTransaction untuk mendapatkan semua data transaksi dari database
func (t *Transaction) GetAllTransaction(mTransaction model.Transaction) ([]model.Response, error) {
	var nTransaction []model.Response

	var args []interface{}
	statment := `
		SELECT DISTINCT b.id, b.bill_date, b.entry_date, b.finish_date,
						e.id, e.name, e.address, e.phone,
						c.id, c.name, c.address, c.phone
		FROM bills b
		JOIN employees e ON b.employee_id = e.id
		JOIN customers c ON b.customer_id = c.id
`

	query := ""
	counter := 1

	if mTransaction.Query.StartDate != "" {
		if counter != 1 {
			query += " AND"
		}
		query += fmt.Sprintf(" b.entry_date >= $%d", counter)
		args = append(args, mTransaction.Query.StartDate)
		counter++
	}
	if mTransaction.Query.EndDate != "" {
		if counter != 1 {
			query += " AND"
		}
		query += fmt.Sprintf(" b.finish_date <= $%d", counter)
		args = append(args, mTransaction.Query.EndDate)
		counter++
	}
	if mTransaction.Query.ProductName != "" {
		if counter != 1 {
			query += " AND"
		}
		query += fmt.Sprintf(" p.name ILIKE $%d", counter)
		args = append(args, "%"+mTransaction.Query.ProductName+"%")
		counter++
	}
	if counter != 1 {
		statment += `
			JOIN details d on b.id = d.bill_id
			JOIN products p on d.product_id = p.id
		`
		statment += "WHERE" + query + " ORDER BY b.id ASC ;"
	} else {
		statment += " ORDER BY b.id ASC ;"

	}

	// Mengambil data dari database
	rows, err := t.db.Query(statment, args...)
	fmt.Printf("Mengambil Data Transaksi\n Query ----\n%s\n Argument ---- %v\n\n", statment, args)
	if err != nil {
		return []model.Response{}, err
	}
	defer rows.Close()

	// Mengambil data dari setiap baris hasil query
	for rows.Next() {
		dump := model.Response{}
		err := rows.Scan(
			&dump.Id,
			&dump.BillDate,
			&dump.EntryDate,
			&dump.FinishDate,
			&dump.Employee.Id,
			&dump.Employee.Name,
			&dump.Employee.Address,
			&dump.Employee.PhoneNumber,
			&dump.Customer.Id,
			&dump.Customer.Name,
			&dump.Customer.Address,
			&dump.Customer.PhoneNumber,
		)

		if err != nil {
			return []model.Response{}, err
		}

		// Mengambil data detail transaksi
		statment = `
		SELECT	d.id, d.bill_id, d.product_id, d.product_price, d.qty,
				p.id, p.name, p.price, p.unit
		FROM details d
		JOIN bills b ON d.bill_id = b.id
		JOIN products p ON d.product_id = p.id
		`
		argsDetail := append(args, dump.Id)
		if query != "" {
			statment += ` WHERE ` + query + fmt.Sprintf(" AND d.bill_id = $%d", counter) + ";"
			//args = append(args, dump.Response.Id)
		} else {
			statment += `WHERE d.bill_id = $1;`
		}

		rowsDetail, err := t.db.Query(statment, argsDetail...)
		fmt.Printf("Mengambil Detail Transaksi\n Query ----\n%s\n Argument ---- %v\n\n", statment, argsDetail)
		if err != nil {
			return []model.Response{}, err
		}
		defer rowsDetail.Close()

		// Mengambil data dari setiap baris hasil query
		for rowsDetail.Next() {
			detail := model.RDetails{}
			err := rowsDetail.Scan(
				&detail.Id,
				&detail.BillId,
				&detail.Product.Id,
				&detail.ProductPrice,
				&detail.Qty,
				&detail.Product.Id,
				&detail.Product.Name,
				&detail.Product.Price,
				&detail.Product.Unit,
			)

			if err != nil {
				return []model.Response{}, err
			}

			dump.BillDetails = append(dump.BillDetails, detail)
		}

		nTransaction = append(nTransaction, dump)

	}
	return nTransaction, nil
}

// Metode GetTransactionById untuk mendapatkan data transaksi berdasarkan ID
func (t *Transaction) GetTransactionById(id int) (model.Transaction, error) {
	var nTransaction model.Transaction

	statment := `
		SELECT b.id, b.bill_date, b.entry_date, b.finish_date,
			e.id, e.name, e.address, e.phone,
			c.id, c.name, c.address, c.phone
		FROM bills b
		JOIN employees e ON b.employee_id = e.id
		JOIN customers c ON b.customer_id = c.id
		WHERE b.id = $1;`

	// Mengambil data dari database
	err := t.db.QueryRow(statment, id).Scan(
		&nTransaction.Response.Id,
		&nTransaction.Response.BillDate,
		&nTransaction.Response.EntryDate,
		&nTransaction.Response.FinishDate,
		&nTransaction.Response.Employee.Id,
		&nTransaction.Response.Employee.Name,
		&nTransaction.Response.Employee.Address,
		&nTransaction.Response.Employee.PhoneNumber,
		&nTransaction.Response.Customer.Id,
		&nTransaction.Response.Customer.Name,
		&nTransaction.Response.Customer.Address,
		&nTransaction.Response.Customer.PhoneNumber,
	)

	if err != nil {
		return model.Transaction{}, err
	}

	// Mengambil data detail transaksi
	statment = `
		SELECT d.id, d.bill_id, d.product_id, d.product_price, d.qty, 
			p.id, p.name, p.price, p.unit
		FROM details d
		JOIN products p ON d.product_id = p.id
		WHERE d.bill_id = $1;`

	rowsDetail, err := t.db.Query(statment, nTransaction.Response.Id)
	if err != nil {
		return model.Transaction{}, err
	}
	defer rowsDetail.Close()

	// Mengambil data dari setiap baris hasil query
	for rowsDetail.Next() {
		detail := model.RDetails{}
		err := rowsDetail.Scan(
			&detail.Id,
			&detail.BillId,
			&detail.Product.Id,
			&detail.ProductPrice,
			&detail.Qty,
			&detail.Product.Id,
			&detail.Product.Name,
			&detail.Product.Price,
			&detail.Product.Unit,
		)

		if err != nil {
			return model.Transaction{}, err
		}

		nTransaction.Response.BillDetails = append(nTransaction.Response.BillDetails, detail)
	}

	return nTransaction, nil

}

// Fungsi ObjTransactionRepo untuk membuat instance TransactionRepo
func ObjTransactionRepo(db *sql.DB) TransactionRepo {
	return &Transaction{db}
}
