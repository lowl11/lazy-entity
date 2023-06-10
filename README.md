# lazy-entity

> Library like EntityFramework on C# <br>
> Building simple repositories to build SQL queries

### SELECT
Go
```go
builder := queryapi.Select()
	builder.
		Fields("id", "full_name", "phone", "is_resident", "contact.id", "COUNT(id)").
		From("users").
		Alias("user").
		Join("contacts", "contact", builder.And(builder.Equal("phone", "$contact.phone"))).
		Where(
			builder.And(
				builder.Or(
					builder.Equal("phone", "+77474858669"),
					builder.ILike("full_name", "%ussayev%"),
				),
				builder.Like("full_name", "%Ussayev%"),
				builder.Equal("id", ":id"),
			),
			builder.Or(
				builder.Equal("is_resident", 1),
				builder.Equal("is_resident", 0),
			),
			builder.Gte("id", 25),
		).
		OrderBy(order_types.Desc, "phone").
		Having(builder.Count("id", 25, builder.Lte)).
		GroupBy("id").
		Offset(50).
		Limit(10)

	fmt.Println("query:")
	fmt.Println(builder.Build())
```

SQL
```sql
SELECT 
        user.id, 
        user.full_name, 
        user.phone, 
        user.is_resident, 
        contact.id, 
        COUNT(user.id)
FROM users AS user
        INNER JOIN contacts AS contact ON (user.phone = contact.phone)
WHERE 
        ((user.phone = '+77474858669' OR user.full_name ILIKE '%ussayev%') AND user.full_name LIKE '%Ussayev%' AND user.id = :id) AND 
        (user.is_resident = 1 OR user.is_resident = 0) AND 
        user.id >= 25
ORDER BY user.phone DESC
GROUP BY user.id
HAVING COUNT(user.id) <= 25
OFFSET 50
LIMIT 10
```

### INSERT
Go
```go
query := queryapi.
    Insert("users").
    Fields("id", "full_name", "phone", "is_resident").
    Variables(1, "Ussayev Yerik", "+77474858669", true).
    Build()

fmt.Println("query:")
fmt.Println(query)
```

SQL
```sql
INSERT INTO users (id, full_name, phone, is_resident)
VALUES (1, 'Ussayev Yerik', '+77474858669', true)
```

### UPDATE
Go
```go
builder := queryapi.Update("users")
builder.
    Set("phone", "+7788001103").
    Set("is_resident", "false").
    Where(builder.Equal("id", 5))

fmt.Println("query:")
fmt.Println(builder.Build())
```

SQL
```sql
UPDATE users
SET
        phone = '+7788001103',
        is_resident = 'false'
WHERE 
        id = 5
```

### DELETE
Go
```go
builder := queryapi.Delete("users")
builder.
    Where(builder.Or(
        builder.Equal("id", 5),
        builder.Gt("id", 10),
    ))

fmt.Println("query:")
fmt.Println(builder.Build())
```

SQL
```sql
DELETE FROM users
WHERE 
        (id = 5 OR id > 10)
```