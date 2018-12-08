# Order/cart API

## 1. Add item to cart
Request
```
Post /order
Content-Type: application/json
```

| Parameter     | Type    | Description  |
| ------------- |---------| -------------|
| orderId       | string  | Order ID
| userId        | string  | User ID
| itemId        | string  | Item ID
| itemName      | string  | Item Name
| price         | float32 | Item Price
| description   | string  | Item Description

Response

Parameters for Success (Status code: 200)

| Parameter     | Type      | Description |
| ------------- |-----------| ------------|
| orderId       | string    | order ID
| userId        | string    | User ID
| orderStatus   | string    | placed, paid
| items         | struct    | item ID, itemName, price, description   
| totalAmount   | float32   | total price


Parameters for Error (Status code: 400)

| Parameter     | Type      | Description  |
| ------------- |-----------| -------------|
| Message       | string    | Error message 
 
## 2. View cart by userId
Request
```
Get /orders/{userId}
Content-Type: application/json
```

| Parameter     | Type      | Description |
| ------------- |-----------| ------------|
| userId        | string    | User ID

Response
Parameters for Success (Status code: 200)

| Parameter     | Type      | Description |
| ------------- |-----------| ------------|
| orderId       | string    | order ID
| userId        | string    | User ID
| orderStatus   | string    | placed, paid
| items         | struct    | item ID, itemName, price, description   
| totalAmount   | float32   | total price


Parameters for Error (Status code: 400)

| Parameter     | Type      | Description  |
| ------------- |-----------| -------------|
| Message       | string    | Error message 
 

## 3. Cart Checkout
Request
```
Put /order/{orderId}
Content-Type: application/json
```

| Parameter     | Type      | Description |
| ------------- |-----------| ------------|
| orderId       | string    | order ID
| userId        | string    | User ID

Response

Parameters for Success (Status code: 200)

| Parameter     | Type      | Description |
| ------------- |-----------| ------------|
| orderId       | string    | order ID
| userId        | string    | User ID
| orderStatus   | string    | placed, paid
| items         | struct    | item ID, itemName, price, description   
| totalAmount   | float32   | total price

Parameters for Error (Status code: 400)

Parameters for Error (Status code: 400)

| Parameter     | Type      | Description  |
| ------------- |-----------| -------------|
| Message       | string    | Error message 


## 4. Remove Item from cart

Request
```
Delete /order/{orderId}
Content-Type: application/json
```

| Parameter     | Type    | Description  |
| ------------- |---------| -------------|
| itemId        | string  | Item ID

Response

Parameters for Success (Status code: 200)

| Parameter     | Type      | Description |
| ------------- |-----------| ------------|
| orderId       | string    | order ID
| userId        | string    | User ID
| orderStatus   | string    | placed, paid
| items         | struct    | item ID, itemName, price, description   
| totalAmount   | float32   | total price

Parameters for Error (Status code: 400)

| Parameter     | Type      | Description  |
| ------------- |-----------| -------------|
| Message       | string    | Error message 


## 5. Remove order

Request
```
Delete /order
Content-Type: application/json
```

| Parameter     | Type    | Description  |
| ------------- |---------| -------------|
| orderId       | string  | Order ID

Response

Parameters for Success (Status code: 200)

| Parameter     | Type      | Description |
| ------------- |-----------| ------------|
| message       | string    | Success Message

Parameters for Error (Status code: 400)

| Parameter     | Type      | Description  |
| ------------- |-----------| -------------|
| Message       | string    | Error message 


## 6. Update cart status
Request
```
Put /order/:orderId
Content-Type: application/json
```

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| userId   | string  | User ID
| orderId  | string  | order ID

Response

Parameters for Success (Status code: 200)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| userId        | string  | User ID
| orderId       | string  | order ID
| items         | struct  | item ID, quantity   
| orderStatus   | string  | placed, paid, removed
| totalAmount   | double  | total price

Parameters for Error (Status code: 400)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| Message   | string  | Error message 

