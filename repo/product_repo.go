package repo

import (
	"api-laundry/model"
	"database/sql"
)

type ProductRepo interface {
	InsertProduct(model.Products) (model.Products, error)
	GetProductById(int) (model.Products, error)
	GetAllProduct(string) ([]model.Products, error)
	UpdateProductById(int, model.Products) (model.Products, error)
	DeleteProductById(int) error
}

type productRepo struct {
	db *sql.DB
}

// Insert product to database
func (p *productRepo) InsertProduct(mProduct model.Products) (model.Products, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return model.Products{}, err
	}

	statment := `INSERT INTO products (name, price, unit) VALUES($1, $2, $3) RETURNING id;`
	err = tx.QueryRow(statment, mProduct.Name, mProduct.Price, mProduct.Unit).Scan(&mProduct.Id)

	if err != nil {

		return model.Products{}, err
	}

	return mProduct, tx.Commit()
}

// Get product by id from database
func (p *productRepo) GetProductById(id int) (model.Products, error) {
	var mProduct model.Products
	statment := `SELECT id, name, price, unit FROM products WHERE id = $1;`

	err := p.db.QueryRow(statment, id).Scan(&mProduct.Id, &mProduct.Name, &mProduct.Price, &mProduct.Unit)
	if err != nil {
		return model.Products{}, err
	}

	return mProduct, nil
}

func (p *productRepo) GetAllProduct(productName string) ([]model.Products, error) {
	// Get all product from database
	var mProducts []model.Products

	statment := `SELECT id, name, price, unit FROM products`
	args := []interface{}{}
	if productName != "" {
		args = append(args, "%"+productName+"%")
		statment += ` WHERE name ILIKE $1;`
	}

	rows, err := p.db.Query(statment, args...)
	if err != nil {
		return []model.Products{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var dump model.Products
		err := rows.Scan(&dump.Id, &dump.Name, &dump.Price, &dump.Unit)
		if err != nil {
			return []model.Products{}, err

		}

		mProducts = append(mProducts, dump)

	}
	return mProducts, nil
}

// Update product by id from database
func (p *productRepo) UpdateProductById(id int, mProduct model.Products) (model.Products, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return model.Products{}, err
	}

	statment := `UPDATE products SET name = $1, price = $2, unit = $3 WHERE id = $4 RETURNING id;`
	err = tx.QueryRow(statment, mProduct.Name, mProduct.Price, mProduct.Unit, id).Scan(&mProduct.Id)
	if err != nil {
		return model.Products{}, err
	}

	return mProduct, tx.Commit()
}

func (p *productRepo) DeleteProductById(id int) error {
	// Delete product by id from database
	statment := `DELETE FROM products WHERE id = $1;`

	_, err := p.db.Exec(statment, id)
	if err != nil {
		return err
	}

	return nil
}

func ObjProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{db}
}
