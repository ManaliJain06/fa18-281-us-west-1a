# Payment API 💵

**Attributes**
|Attribute Name	| Type	| Description|
|---------------|-------|------------|
|paymentId |String|Payment| Id         |
|userId |	String  |	User  |	 Id |
|orderId |	String | Unique | Order Number |
|totalAmount |Double |Total | Amount Paid |
|status	| Boolean |	Payment Status (True = Paid, False = Cancelled) |
|paymentDate |DateTime	|Paid date|


1. Pay for an order

    **Request**
    ```
    POST /payment
    Content-Type: application/json
    ```
    **Parameters**

    |Parameter	|Type |	Description|
    |-----|-----|------|
    |userId	|String|	User Id|
    |orderId|	String|	Unique order number|
    |totalAmount|	Double|	Total amount of the order|

    **Response**

    Parameters for Success (Status code: 200)

    |Parameter	|Type	|Description  |
    |----|----|----|
    |paymentId	|String	|Unique order number|
    |userId	|String	|User Id
    |orderId	|String	|Unique order number
    |totalAmount	|Double	|Total amount paid
    |status	| Boolean	|True (True = Paid, False = Cancelled)
    |paymentDate	|String	|Date of the payment

    **Parameters for Error (Status code: 400)**
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |message	|String|	Error Message|


2. Get Payment List

    **Request**
    ```
    GET /payment
    Content-Type: application/json
    ```
    **Parameters**

    None

    **Response**

    Parameters for Success (Status code: 200)

    |Parameter	|Type	|Description  |
    |----|----|----|
    |paymentId	|String	|Unique order number|
    |userId	|String	|User Id
    |orderId	|String	|Unique order number
    |totalAmount	|Double	|Total amount paid
    |status	|Boolean  |Payment Status (True = Paid, False = Cancelled)
    |paymentDate	|String	|Date of the payment

    **Parameters for Error (Status code: 400)**
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |message	|String|	Error Message|


3. View Payment Details by orderId

    **Request**

    ```
    GET /payment/:orderId
    Content-Type: application/json
    ```

    **Parameters**
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |orderId	|String|	Unique order number|

    **Response**

    Parameters for Success (Status code: 200)

    |Parameter	|Type	|Description  |
    |----|----|----|
    |paymentId	|String	|Unique order number|
    |userId	|String	|User Id
    |orderId	|String	|Unique order number
    |totalAmount	|Double	|Total amount paid
    |status	|Boolean  |Payment Status (True = Paid, False = Cancelled)
    |paymentDate	|String	|Date of the payment

    **Parameters for Error (Status code: 400)**
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |message	|String|	Error Message|


4. View Payment Details by userId

    **Request**

    ```
    GET /payment/:userId
    Content-Type: application/json
    ```

    **Parameters**
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |userId	|String|	User Id|

    **Response**

    Parameters for Success (Status code: 200)

    |Parameter	|Type	|Description  |
    |----|----|----|
    |paymentId	|String	|Unique order number|
    |userId	|String	|User Id
    |orderId	|String	|Unique order number
    |totalAmount	|Double	|Total amount paid
    |status	|Boolean  |Payment Status (True = Paid, False = Cancelled)
    |paymentDate	|String	|Date of the payment

    **Parameters for Error (Status code: 400)**
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |message	|String|	Error Message|

5. Edit Payment

    **Request**

    ```
    PUT /payment/:paymentId
    Content-Type: application/json
    ```

    **Parameters**
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |paymentId	|String|	User Id|
    |totalAmount|	Double|	Total amount of the order |
    |status|	Boolean|	(True = Paid, False = Cancelled)|

    **Response**

    Parameters for Success (Status code: 200)

    |Parameter	|Type	|Description  |
    |----|----|----|
    |paymentId	|String	|Unique order number|
    |userId	|String	|User Id
    |orderId	|String	|Unique order number
    |totalAmount	|Double	|Total amount paid
    |status	|Boolean  |Payment Status (True = Paid, False = Cancelled)
    |paymentDate	|String	|Date of the payment

    **Parameters for Error (Status code: 400)**
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |message	|String|	Error Message|


6. Cancel Payment

    **Request**

    ```
    DELETE /payment/:paymentId
    Content-Type: application/json
    ```

    **Cancel payment will only set the payment status to False**

    **Parameters**
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |paymentId	|String|	User Id|

    **Response**

    Parameters for Success (Status code: 200)

    |Parameter	|Type	|Description  |
    |----|----|----|
    |paymentId	|String	|Unique order number|
    |userId	|String	|User Id
    |orderId	|String	|Unique order number
    |totalAmount	|Double	|Total amount paid
    |status	|Boolean  |Payment Status (True = Paid, False = Cancelled)
    |paymentDate	|String	|Date of the payment

    **Parameters for Error (Status code: 400)**
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |message	|String|	Error Message|
