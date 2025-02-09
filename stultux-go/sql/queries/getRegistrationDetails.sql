-- name: GetRegistrationDetails :many
SELECT 
    n.name AS first_name,
    l.name AS last_name,
    p.password
FROM 
    available_names n
CROSS JOIN 
    available_last l
CROSS JOIN 
    available_passwords p;

