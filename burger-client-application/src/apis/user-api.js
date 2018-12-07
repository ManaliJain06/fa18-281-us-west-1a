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

export const callLoginApi = (payload) => {
    console.log("payload", JSON.stringify(payload));
    return axios.post(`${userUrl}/users/signin`, payload, headers )
        .then( function(res) {
            console.log('call login api, res:', res);
            return res
        }).catch( function(err) {
            console.log('error calling login api, error:', err);
            return err
        });
};

export const callRegisterAPI = (payload) => {
    console.log("payload", JSON.stringify(payload));
    return axios.post(`${userUrl}/users/signup`, payload, headers )
        .then( function(res){
            console.log('call register api, res:', res);
           return res
        }).catch( function (err) {
            console.log('error calling login api, error:', err);
           return err
        });
};

export const callUserOrdersAPI = (userId) => {
    return axios.get(`${orderUrl}/orders/${userId}`, headers )
        .then( function(res){
            console.log('call register api, res:', res);
            return res
        }).catch( function (err) {
            console.log('error calling login api, error:', err);
            return err
        });
};