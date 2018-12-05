/*
	Client REST API for calls to restaurant microservice
*/

const api = process.env.REACT_APP_CONTACTS_API_URL || 'http://54.183.92.48:8000';
const headers = {
    'Accept': 'application/json'
};

export const getRestaurants = (zipcode) =>
    fetch(`${api}/restaurant/zipcode/${zipcode}`, {
        method: 'GET',
        headers: headers,
    }).then(res => {
        return res;
    }).catch(error => {
        console.log("get restaurant error");
        return error;
    });

// export const addRestaurant = (menuItem) =>
//     fetch(`${api}/menu/item`, {
//         method: 'POST',
//         headers: {
//             ...headers,
//             'Content-Type': 'application/json'
//         },
//         credentials:'include',
//         body: JSON.stringify(menuItem)
//     }).then(res => {
//         return res;
//     }).catch(error => {
//         console.log("This is error");
//         return error;
//     });


