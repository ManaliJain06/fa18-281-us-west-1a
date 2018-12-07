/*
	Client REST API calls to order microservice
*/
import {kongAPI, orderUrl} from '../actions/urlConstant';

const api = kongAPI;
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

export const deleteOrderItem = (orderId,itemId) =>
    fetch(`${api}/order/${orderId}`, {
        method: 'DELETE',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({itemId:itemId})
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
