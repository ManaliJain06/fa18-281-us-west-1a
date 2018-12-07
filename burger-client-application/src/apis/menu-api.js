/*
	Client REST API calls to menu microservice
*/
import {kongAPI} from '../actions/urlConstant';

const api = kongAPI;
const headers = {
    'Accept': 'application/json'
};

export const getRestaurantMenuItems = (restaurantId) =>
    fetch(`${api}/menu/${restaurantId}`, {
        method: 'GET',
        headers: headers,
    }).then(res => {
        return res;
    }).catch(error => {
        console.log("[menu-api] getRestaurantMenuItems() Error !!!");
        return error;
    });

export const AddMenuItem = (menuItem) =>
    fetch(`${api}/menu/`, {
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
