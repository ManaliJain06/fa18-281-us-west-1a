/*
	Client REST API calls to order microservice
*/

const api = process.env.REACT_APP_CONTACTS_API_URL || 'http://54.153.121.217:8000';
const headers = {
    'Accept': 'application/json'
};

export const getOrderItems = (orderId) =>
    fetch(`${api}/order/${orderId}`, {
        method: 'GET',
        headers: headers,
    }).then(res => {
        return res;
    }).catch(error => {
        console.log("[order-api] getOrderItems() Error !!!");
        return error;
    });

export const  addOrderItem= (menuItem) =>
    fetch(`${api}/order`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(menuItem)
    }).then(res => {
        return res;
    }).catch(error => {
        console.log("This is error");
        return error;
    });
