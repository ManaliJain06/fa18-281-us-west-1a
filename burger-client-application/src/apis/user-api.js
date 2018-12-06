/*
	Client REST API calls to user microservice
*/

import axios from "axios";
import {orderUrl} from "../actions/urlConstant";

export const callLoginApi = (payload) => () => {
    console.log("payload", JSON.stringify(payload))
    axios.put(`${orderUrl}/order/${orderId}` )
        .then( res => {
            console.log('after axiosOrderUpdateToPaid, res:', res);

        }).catch( res => {
        console.log('xx  error axiosOrderUpdateToPaid, error:', res);
    })

};