# Order/cart API

## 1. Add item to new cart
Request
```
Post /order
Content-Type: application/json
```

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| userId   | string  | User ID
| itemId    | string  | Item ID
| quantity     | integer  | Item quantity

Response

Parameters for Success (Status code: 200)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| orderId   | string  | order ID: itemId, quantity
| orderStatus    | string  | placed, paid, removed
| totalAmount     | double  | total price


Parameters for Error (Status code: 400)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| Message   | string  | Error message 
 
## 2. View cart by userId
Request
```
Get /order
Content-Type: application/json
```

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| userId   | string  | User ID

Response
Parameters for Success (Status code: 200)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| userId   | string  | User ID
| orderId   | string  | order ID: itemId, quantity
| orderStatus    | string  | placed, paid, removed
| totalAmount     | double  | total price

Parameters for Error (Status code: 400)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| Message   | string  | Error message 
 

## 3. Edit cart
Request
```
Put /order/
Content-Type: application/json
```

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| userId   | string  | User ID
| itemId    | string  | Item ID
| quantity     | integer  | Item quantity

Response

Parameters for Success (Status code: 200)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| userId   | string  | User ID
| orderId   | string  | order ID: itemId, quantity
| orderStatus    | string  | placed, paid, removed
| totalAmount     | double  | total price

Parameters for Error (Status code: 400)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| Message   | string  | Error message 


## 4. Delete cart
Request
```
Delete /order/:orderId-
Content-Type: application/json
```

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| userId   | string  | User ID

Response

Parameters for Success (Status code: 200)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| orderId   | string  | order ID: itemId, quantity
| orderStatus    | string  | placed, paid, removed

Parameters for Error (Status code: 400)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| Message   | string  | Error message 


## 5. Update cart status
Request
```
Put /order/:orderId
Content-Type: application/json
```

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| userId   | string  | User ID
| orderId   | string  | order ID: itemId, quantity

Response

Parameters for Success (Status code: 200)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| userId   | string  | User ID
| orderId   | string  | order ID: itemId, quantity
| orderStatus    | string  | placed, paid, removed
| totalAmount     | double  | total price

Parameters for Error (Status code: 400)

| Parameter        | Type           | Description  |
| ------------- |-------------| -----|
| Message   | string  | Error message 

