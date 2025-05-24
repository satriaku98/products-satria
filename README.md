# products-satria
This is for Erajaya.

Aplikasi ini mengimplementasikan Redis Cache dan Docker

## Getting started

1.  Create .env file in root

    Gunakan template .env.example, anda bisa menduplikasi file tersebut, dan isikan sesuai configurasi yang ada di local anda 
2.  Jalan kan aplikasi
```
go run cmd/main.go
```
3. Akses Dokumentasi
```
localhost:8080/swagger/index.html
```
4. Opsional. gunakan docker pull untuk mendapatkan image
```
docker pull satriaku/erajaya-satria
```

### Architecture
Aplikasi ini menggunakan Clean Architecture (oleh Uncle Bob) yang bertujuan untuk memudahkan testing, scalability, dan perubahan teknologi (misalnya ganti dari GORM ke SQLx tanpa ubah logika bisnis).
Clean Architecture ini terwujud dengan pemisahan antara :
 - handler
 - service
 - repository
 - model

### Erajaya API Documentation
This is an API documentation for Erajaya. 

(dokumentasi ini dibuat dengan AI berdasarkan docs.json di swagger) 

Base path: `/api/v1`  
Version: `1.0`

---

## 📦 Products

### GET `/products`

**Description:** Get a list of products  
**Query Parameters:**

| Name | Type | Description |
|------|------|-------------|
| `sort` | `string` | Sort parameters in the format `'column:direction[,column2:direction]'` <br> Example: `created_at:desc,price:asc` |

**Responses:**

- `200 OK`: List of products  
- `400 Bad Request`: Invalid query or input  
- `500 Internal Server Error`: Server error

**Response Example (`200 OK`):**
```json
[
  {
    "id": 1,
    "name": "Product Name",
    "description": "Product Description",
    "price": 100000,
    "quantity": 50,
    "created_at": "2025-05-24T12:00:00Z"
  }
]
```

---

### POST `/products`

**Description:** Add a new product  
**Request Body:**

```json
{
  "name": "Product Name",
  "description": "Product Description",
  "price": 100000,
  "quantity": 50
}
```

**Responses:**

- `201 Created`: Product successfully added  
- `400 Bad Request`: Invalid input  
- `500 Internal Server Error`: Server error

**Response Example (`201 Created`):**
```json
{
  "id": 2,
  "name": "Product Name",
  "description": "Product Description",
  "price": 100000,
  "quantity": 50,
  "created_at": "2025-05-24T12:00:00Z"
}
```

---

## 📘 Definitions

### model.Product

```json
{
  "id": 1,
  "name": "string",
  "description": "string",
  "price": 100000,
  "quantity": 50,
  "created_at": "2025-05-24T12:00:00Z"
}
```

### request.ProductRequest

```json
{
  "name": "string",
  "description": "string",
  "price": 100000,
  "quantity": 50
}
```

