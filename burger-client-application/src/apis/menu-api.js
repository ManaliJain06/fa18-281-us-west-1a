/*
	Client REST API calls to menu microservice
*/

const api = process.env.REACT_APP_CONTACTS_API_URL || 'http://localhost:8080';
const headers = {
    'Accept': 'application/json'
};

export const getAllMenuItems = () =>
    fetch(`${api}/menu/items`, {
        method: 'GET',
        headers: headers,
    }).then(res => {
        return res;
    }).catch(error => {
        console.log("[menu-api] getAllMenuItems() Error !!!");
        return error;
    });

export const createNewMenuItem = () =>
    fetch(`${api}/menu/item`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        credentials:'include'
        // body: JSON.stringify()
    }).then(res => {
        return res;
    }).catch(error => {
        console.log("This is error");
        return error;
    });