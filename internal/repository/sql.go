package repository

var sqlCreateCustomer = `
	INSERT INTO customers(first_name, last_name, birth_date) VALUES(lower(:first_name), lower(:last_name), :birth_date) RETURNING *;
`
