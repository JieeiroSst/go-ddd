# Web Service Restaurant
WS for get food and orders.

# Start 🚀
    1. Clone this project -> https://github.com/andresgaviria2020/go-ddd.git
    2. Make sure port 8080 is not busy.
    3. go get all 
    4. go run main.go


# Project Structure 🧱

```
go-ddd
├─ .DS_Store
├─ BD_Restaurant.sql
├─ README.md
└─ WS_Restaurant
   ├─ .DS_Store
   ├─ .env
   ├─ application
   │  └─ RestaurantController.go
   ├─ docs
   │  ├─ docs.go
   │  ├─ swagger.json
   │  └─ swagger.yaml
   ├─ domain
   │  ├─ dto
   │  │  ├─ FoodDto.go
   │  │  ├─ Response.go
   │  │  └─ UserDto.go
   │  ├─ entity
   │  │  ├─ Food.go
   │  │  ├─ Order.go
   │  │  └─ Users.go
   │  └─ service
   │     ├─ RestaurantService.go
   │     └─ RestaurantServiceImpl.go
   ├─ go.mod
   ├─ go.sum
   ├─ infrastructure
   │  ├─ persistence
   │  │  ├─ DbHelper.go
   │  │  └─ FoodRepositoryImpl.go
   │  └─ repository
   │     └─ FoodRepository.go
   ├─ interfaces
   │  └─ middleware
   │     ├─ CORSMiddleware.go
   │     └─ server
   │        ├─ Server.go
   │        └─ ServerImpl.go
   └─ main.go

```

# Endpoints
    - GET /food
