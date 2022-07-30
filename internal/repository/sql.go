package repository

var sqlCreateCustomer = `
	INSERT INTO customers(first_name, last_name, birth_date) 
	VALUES(lower(:first_name), lower(:last_name), :birth_date) RETURNING *;
`
var sqlGetCustomerById = `
	SELECT 
		id,
		first_name,
		last_name,
		birth_date,
		updated_at,
		created_at
	FROM customers
	WHERE id = $1 AND deleted_at IS NULL
`
