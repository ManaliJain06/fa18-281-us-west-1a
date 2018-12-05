/*
	Client REST API calls to payment microservice
*/

import axios from 'axios';
import * as actionType from '../actions/payment-actions';
import {paymentUrl, orderUrl} from '../actions/urlConstant';

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


export const axiosOrderUpdateToPaid = (orderId) => () => {
	console.log('axiosOrderUpdateToPaid')

  axios.put(`${orderUrl}/order/${orderId}` )
    .then( res => {
      console.log('after axiosOrderUpdateToPaid, res:', res);

    }).catch( res => {
      console.log('xx  error axiosOrderUpdateToPaid, error:', res);
    })

}

export const axiosCreatePayment = (data, router) => (dispatch) => {
  console.log('axiosCreatePayment');
	console.log('before axios router', router)
  axios.post(`${paymentUrl}/payments`, {
		userId: "0",
		orderId: data.orderId,
		totalAmount: data.totalAmount,
  }).then( res => {
    console.log('after axiosCreatePayment, res:', res);
		console.log('after axios router', router)
    
		dispatch(createPayment(res.data));
		
		// remove orderId in localStorage
		console.log('removing localstorage orderId')
		localStorage.removeItem('orderId');

		// make an API call to order, to update that order has already been paid
		console.log('before axiosOrderUpdateToPaid');
		// axiosOrderUpdateToPaid(data.orderId);
		axios.put(`${orderUrl}/order/${data.orderId}` )
    .then( res => {
      console.log('after axiosOrderUpdateToPaid, res:', res);

    }).catch( res => {
      console.log('xx  error axiosOrderUpdateToPaid, error:', res);
    })
		console.log('after axiosOrderUpdateToPaid');


		router.push('/paymentSuccess');

  }).catch( res => {
		console.log('xx error axiosCreatePayment, error:', res);
		router.push('/paymentError');
  })
  
}

export const createPayment = (data) => {
  return {
    type: actionType.PAYMENT_CREATE,
    data,
  }
}

export const axiosGetOrder = (orderId) => (dispatch) => {
  console.log('axiosGetOrder');

  axios.get(`${orderUrl}/order/${orderId}` )
    .then( res => {
      console.log('after axiosGetOrder, res:', res);

      dispatch(getOrder(res.data));
    }).catch( res => {
      console.log('xx  error axiosGetOrder, error:', res);
    })
}

export const getOrder = (data) => {
  return {
    type: actionType.PAYMENT_GET_ORDER_DETAIL,
    data,
  }
}

