# User
1. Sign Up

    Request

    POST /users/signup

    Content-Type: application/json

   
    | Parameter        | Type           | Description  |
    | ------------- |:-------------:| -----:|
    | firstName   | string  | User First Name
    | lastName    | string  | User last name
    | address     | string  | User address
    | city        | string  | User city
    | state       | string  | User state
    | zipcode     | string  | User zipcode
    | phone       | string  | User phone
    | email       | string  | User email
    | password    | string  | User password
    | creditcard  | string  | User credit card number

    Response
    

    Parameters for Success Status code: 200 
    
    
    
    | Parameter        | Type           | Description  |
    | ------------- |:-------------:| -----:|
    | firstName   | string  | User First Name
    | lastName    | string  | User last name
    | email       | string  | User email
    | message     | string  | Successful sign up message  

    Parameters for Error (Status code: 401)
    
    | Parameter        | Type           | Description  |
    | ------------- |:-------------:| -----:|
    | message     | string  | Error Message

2. Sign In

	POST /users/signin
	Content-type: application/json

	| Parameter        | Type           | Description  |
	| ------------- |:-------------:| -----:|
	| email     | string  | User User email
	| password  | string  | User password


	Response

	Parameters for Success (Status code: 200)
	
	| Parameter        | Type           | Description  |
	| ------------- |:-------------:| -----:|
	| firstName   | string  | User First Name
	| lastName    | string  | User last name
	| email       | string  | User email
	| Token       | string  | JSON WEB TOKEN  

	Parameters for Error (Status code: 401)
	
	| Parameter        | Type           | Description  |
	| ------------- |:-------------:| -----:|
	| message     | string  | Error Message

3. Get User by ID

	GET /users/:id
	
	Content-type: application/json

	| Parameter        | Type           | Description  |
	| ------------- |:-------------:| -----:|
	| id     | string  | user id


	Response

	Parameters for Success (Status code: 200)
	
	
	| Parameter        | Type           | Description  |
	| ------------- |:-------------:| -----:|
	| firstName   | string  | User First Name
    	| lastName    | string  | User last name
    	| address     | string  | User address
    	| city        | string  | User city
    	| state       | string  | User state
    	| zipcode     | string  | User zipcode
    	| phone       | string  | User phone
    	| email       | string  | User email
    	| password    | string  | User password
    	| creditcard  | string  | User credit card number 

	Parameters for Error (Status code: 401)
	
	| Parameter        | Type           | Description  |
	| ------------- |:-------------:| -----:|
	| message     | string  | Error Message

4. Delete User

	Delete /users/:id

	Content-type: application/json

	| Parameter        | Type           | Description  |
	| ------------- |:-------------:| -----:|
	| id     | string  | user id


	Response

	Parameters for Success (Status code: 200)
	
	| Parameter        | Type           | Description  |
  	| ------------- |:-------------:| -----:|
   	| message   | string  | Message with successful deletion of the user

	Parameters for Error (Status code: 401)
	
	| Parameter        | Type           | Description  |
	| ------------- |:-------------:| -----:|
	| message     | string  | Error Message

5. Edit user

	put /users/:id
	Content-type: application/json

	| Parameter        | Type           | Description  |
    	| ------------- |:-------------:| -----:|
    	| firstName   | string  | User First Name
    	| lastName    | string  | User last name
    	| address     | string  | User address
    	| city        | string  | User city
    	| state       | string  | User state
    	| zipcode     | string  | User zipcode
    	| phone       | string  | User phone
    	| email       | string  | User email
    	| password    | string  | User password
    	| creditcard  | string  | User credit card number 


	Response

	Parameters for Success (Status code: 200)
	
	| Parameter        | Type           | Description  |
    	| ------------- |:-------------:| -----:|
    	| firstName   | string  | User First Name
    	| lastName    | string  | User last name
    	| address     | string  | User address
    	| city        | string  | User city
    	| state       | string  | User state
    	| zipcode     | string  | User zipcode
    	| phone       | string  | User phone
    	| email       | string  | User email
    	| password    | string  | User password
    	| creditcard  | string  | User credit card number 

	Parameters for Error (Status code: 401)
	
	| Parameter        | Type           | Description  |
	| ------------- |:-------------:| -----:|
	| message     | string  | Error Message

  
