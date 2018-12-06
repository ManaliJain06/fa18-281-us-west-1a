/*
	Client REST API calls to user microservice
*/

import axios from "axios";
import {orderUrl, userUrl} from "../actions/urlConstant";

let headers = {
    headers:{
        "Content-Type": "application/json"
    }
};

export const callLoginApi = (payload) => () => {
    console.log("payload", JSON.stringify(payload))
    return axios.post(`${userUrl}/users/signin`, payload, headers )
        .then( (res) => {
            console.log('call login api, res:', res);
            return res;
        }).catch( (err) => {
            console.log('error calling login api, error:', res);
            return err;
    })
};

export const callRegisterAPI = (payload) => () => {
    console.log("payload", JSON.stringify(payload))
    return axios.post(`${userUrl}/users/signup`, payload, headers )
        .then( (res) => {
            console.log('call register api, res:', res);
            return res;
        }).catch( (err) => {
            console.log('error calling login api, error:', res);
            return err;
        })
};