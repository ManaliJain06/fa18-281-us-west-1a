/*
	Client REST API calls to payment microservice
*/

import axios from 'axios';
import * as actionType from '../actions/payment-actions';
import {paymentUrl} from '../actions/urlConstant';

export const axiosGetAll = () => (dispatch) => {
  console.log('axiosGetAll');

  axios.get(`${paymentUrl}/payments` )
    .then( res => {
      console.log('after axiosGetAll, res:', res);

      dispatch(getAll(res.data));
    }).catch( res => {
      console.log('xx  error axiosGetAll, error:', res);
    })
}

export const getAll = (data) => {
  return {
    type: actionType.PAYMENT_GET_ALL,
    data,
  }
}


export const axiosCreatePayment = (data) => (dispatch) => {
  console.log('axiosCreatePayment');

  axios.post(`${paymentUrl}/payments`, {
		userId: 0,
		orderId: data.orderId,
		totalAmount: data.totalAmount,
  }).then( res => {
    console.log('after axiosCreatePayment, res:', res);
    
    dispatch(createPayment(res.data));

  }).catch( res => {
    console.log('xx error axiosCreatePayment, error:', res);
  })
  
}

export const createPayment = (data) => {
  return {
    type: actionType.PAYMENT_CREATE,
    data,
  }
}

