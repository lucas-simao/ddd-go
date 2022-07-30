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

var sqlUpdateCustomerById = `
	UPDATE customers
	SET 
		first_name = lower(:first_name),
		last_name = lower(:last_name),
		birth_date = :birth_date
	WHERE id = :id
	RETURNING *
`

var sqlDeleteCustomerById = `
	DELETE FROM customers WHERE id = $1
`
