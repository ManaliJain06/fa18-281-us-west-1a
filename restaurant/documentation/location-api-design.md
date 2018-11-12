# Location listing

## 1. Add Restaurant to a Location(zip code)

**- Request**

POST /restaurant

Content-Type: application/json

Parameters

| Parameter      | Type   | Description                                    |
| -------------- | ------ | ---------------------------------------------- |
| restaurantName | string | Restaurant name                                |
| zipcode        | string | Restaurant zip code                            |
| phone          | string | Restaurant phone number                        |
| addressLine1   | string | Restaurant address line 1                      |
| addressLine2   | string | Restaurant address line 2                      |
| city           | string | City of the restaurant                         |
| state          | string | State of the restaurant                        |
| country        | string | Country of the restaurant                      |
| hours          | string | Restaurant open hours                          |
| acceptedCards  | string | Cards accepted by restaurant                   |
| distance       | string | Restaurant distance from that zipcode location |
| email          | string | Restaurant email                               |

**- Response**

Parameters for Success (Status code: 200)

| Parameter      | Type   | Description                                    |
| -------------- | ------ | ---------------------------------------------- |
| restaurantName | string | Restaurant name                                |
| restaurantId   | string | Restuarant ID                                  |
| zipcode        | string | Restaurant zip code                            |
| phone          | string | Restaurant phone number                        |
| addressLine1   | string | Restaurant address line 1                      |
| addressLine2   | string | Restaurant address line 2                      |
| city           | string | City of the restaurant                         |
| state          | string | State of the restaurant                        |
| country        | string | Country of the restaurant                      |
| hours          | string | Restaurant open hours                          |
| acceptedCards  | string | Cards accepted by restaurant                   |
| distance       | string | Restaurant distance from that zipcode location |
| email          | string | Restaurant email                               |

Parameters for Error (Status code: 401)

| Parameter | Type   | Description   |
| --------- | ------ | ------------- |
| message   | string | Error message |



## 2. Get restaurant list for location(zip code)

**- Request**

GET /restaurant/zipcode/:zipcode

Content-Type: application/json

Parameters

| Parameter | Type   | Description                                                  |
| --------- | ------ | ------------------------------------------------------------ |
| zipcode   | string | Area zip code will be send to the go apis for getting the list of all the restaurants in that area |

**- Response**

Response will come in the form of a list of restaurants

Parameters for Success (Status code: 200)

| Parameter      | Type   | Description                                    |
| -------------- | ------ | ---------------------------------------------- |
| restaurantName | string | Restaurant name                                |
| restaurantId   | string | Restuarant ID                                  |
| zipcode        | string | Restaurant zip code                            |
| phone          | string | Restaurant phone number                        |
| addressLine1   | string | Restaurant address line 1                      |
| addressLine2   | string | Restaurant address line 2                      |
| city           | string | City of the restaurant                         |
| state          | string | State of the restaurant                        |
| country        | string | Country of the restaurant                      |
| hours          | string | Restaurant open hours                          |
| acceptedCards  | string | Cards accepted by restaurant                   |
| distance       | string | Restaurant distance from that zipcode location |
| email          | string | Restaurant email                               |

Parameters for Error (Status code: 401)

| Parameter | Type   | Description   |
| --------- | ------ | ------------- |
| message   | string | Error message |



## 3. Get restaurant by restaurant Id

**- Request**

GET /restaurant/:restaurantId

Content-Type: application/json

Parameters

| Parameter    | Type   | Description                                 |
| ------------ | ------ | ------------------------------------------- |
| restaurantId | string | Restaurant Id to get the restaurant details |

**- Response**

Parameters for Success (Status code: 200)

| Parameter      | Type   | Description                                    |
| -------------- | ------ | ---------------------------------------------- |
| restaurantName | string | Restaurant name                                |
| restaurantId   | string | Restuarant ID                                  |
| zipcode        | string | Restaurant zip code                            |
| phone          | string | Restaurant phone number                        |
| addressLine1   | string | Restaurant address line 1                      |
| addressLine2   | string | Restaurant address line 2                      |
| city           | string | City of the restaurant                         |
| state          | string | State of the restaurant                        |
| country        | string | Country of the restaurant                      |
| hours          | string | Restaurant open hours                          |
| acceptedCards  | string | Cards accepted by restaurant                   |
| distance       | string | Restaurant distance from that zipcode location |
| email          | string | Restaurant email                               |

Parameters for Error (Status code: 401)

| Parameter | Type   | Description   |
| --------- | ------ | ------------- |
| message   | string | Error message |



## 4. Update restaurants details

**- Request**

PUT /restuarant/:restaurantId

Content-Type: application/json

Parameters

| Parameter      | Type   | Description                                    |
| -------------- | ------ | ---------------------------------------------- |
| restaurantName | string | Restaurant name                                |
| zipcode        | string | Restaurant zip code                            |
| phone          | string | Restaurant phone number                        |
| addressLine1   | string | Restaurant address line 1                      |
| addressLine2   | string | Restaurant address line 2                      |
| city           | string | City of the restaurant                         |
| state          | string | State of the restaurant                        |
| country        | string | Country of the restaurant                      |
| hours          | string | Restaurant open hours                          |
| acceptedCards  | string | Cards accepted by restaurant                   |
| distance       | string | Restaurant distance from that zipcode location |
| email          | string | Restaurant email                               |

**- Response**

Parameters for Success (Status code: 200)

| Parameter      | Type   | Description                                    |
| -------------- | ------ | ---------------------------------------------- |
| restaurantName | string | Restaurant name                                |
| restaurantId   | string | Restuarant ID                                  |
| zipcode        | string | Restaurant zip code                            |
| phone          | string | Restaurant phone number                        |
| addressLine1   | string | Restaurant address line 1                      |
| addressLine2   | string | Restaurant address line 2                      |
| city           | string | City of the restaurant                         |
| state          | string | State of the restaurant                        |
| country        | string | Country of the restaurant                      |
| hours          | string | Restaurant open hours                          |
| acceptedCards  | string | Cards accepted by restaurant                   |
| distance       | string | Restaurant distance from that zipcode location |
| email          | string | Restaurant email                               |

Parameters for Error (Status code: 401)

| Parameter | Type   | Description   |
| --------- | ------ | ------------- |
| message   | string | Error message |



## 5. Delete Restaurant from a location(zipcode)

**- Request**

DELETE /restaurant/:restaurantId

Content-Type: application/json

Parameters

| Parameter    | Type   | Description                                                  |
| ------------ | ------ | ------------------------------------------------------------ |
| restaurantId | string | Restaurant Unique Id to delete the restaurant from that zip location when the resturant has been shfited from that place or is no more in service |

**- Response**

Parameters for Success (Status code: 200)

| Parameter | Type    | Description                                  |
| --------- | ------- | -------------------------------------------- |
| status    | boolean | Indicate that delete operation is successful |

Parameters for Error (Status code: 401)

| Parameter | Type   | Description   |
| --------- | ------ | ------------- |
| message   | string | Error message |



