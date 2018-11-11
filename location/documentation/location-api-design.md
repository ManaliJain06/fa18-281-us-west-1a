# Location listing

## 1. Add Restaurant to a Location(zip code)

**Request**

POST /location/restaurant

Content-Type: application/json

Parameters

| Parameter      | Type   | Description                                    |
| -------------- | ------ | ---------------------------------------------- |
| restaurantName | string | Restaurant name                                |
| restaurantId   | String | Restuarant ID                                  |
| area           | string | Restaurant area                                |
| phone          | string | Restaurant phone number                        |
| address        | string | Restaurant address                             |
| hours          | string | Restaurant open hours                          |
| acceptedCards  | string | Cards accepted by restaurant                   |
| distance       | string | Restaurant distance from that zipcode location |
| email          | string | Restaurant email                               |

**Response**

Parameters for Success (Status code: 200)

| Parameter      | Type   | Description                                    |
| -------------- | ------ | ---------------------------------------------- |
| restaurantId   | string | Restaurant Id created by mongodb               |
| restaurantName | string | Restaurant name                                |
| restaurantId   | String | Restuarant ID                                  |
| area           | string | Restaurant area                                |
| phone          | string | Restaurant phone number                        |
| address        | string | Restaurant address                             |
| hours          | string | Restaurant open hours                          |
| acceptedCards  | string | Cards accepted by restaurant                   |
| distance       | string | Restaurant distance from that zipcode location |
| email          | string | Restaurant email                               |

Parameters for Error (Status code: 401)

| Parameter | Type   | Description   |
| --------- | ------ | ------------- |
| message   | string | Error message |



## 2. Get burger restaurants for location(zip code)

**Request**

GET /location/:zipcode

Content-Type: application/json

Parameters

| Parameter | Type   | Description                                                  |
| --------- | ------ | ------------------------------------------------------------ |
| zipcode   | string | Area zip code will be send to the go apis for getting the list of all the restaurants in that area |

**Response**

Response will come in the form of a list of restaurants

Parameters for Success (Status code: 200)

| Parameter      | Type   | Description                                    |
| -------------- | ------ | ---------------------------------------------- |
| restaurantName | string | Restaurant name                                |
| restaurantId   | String | Restuarant Id created by mongodb               |
| area           | string | Restaurant area                                |
| phone          | string | Restaurant phone number                        |
| address        | string | Restaurant address                             |
| hours          | string | Restaurant open hours                          |
| acceptedCards  | string | Cards accepted by Restaurants                  |
| distance       | string | Restaurant distance from that zipcode location |
| email          | string | Restaurant email                               |

Parameters for Error (Status code: 401)

| Parameter | Type   | Description   |
| --------- | ------ | ------------- |
| message   | string | Error message |



## 3. Delete Restaurant from the list of a location(zipcode)

**Request**

DELETE /location/restaurant/:restaurantId

Content-Type: application/json

Parameters

| Parameter    | Type   | Description                                                  |
| ------------ | ------ | ------------------------------------------------------------ |
| restaurantId | string | Restaurant Unique Id to delete the restaurant from that zip location when the resturant has been shfited from that place or is no more in service |

**Response**

Parameters for Success (Status code: 200)

| Parameter | Type    | Description                                  |
| --------- | ------- | -------------------------------------------- |
| status    | boolean | Indicate that delete operation is successful |

Parameters for Error (Status code: 401)

| Parameter | Type   | Description   |
| --------- | ------ | ------------- |
| message   | string | Error message |

