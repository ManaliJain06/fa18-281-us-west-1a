import * as actionType from '../actions/payment-actions';

const initialState = {
  data: [],
  paymentData: {},
  orderDetail: {
    orderId: '',
    totalAmount: 0,
  },
}

const PaymentReducer = (state = initialState, action) => {
  switch(action.type) {
    case actionType.PAYMENT_GET_ALL: {
      console.log('Payment Get all');
      console.log('action.data =', action.data);
      
      return {
        ...state,
        data: action.data,

      }
    }
    case actionType.PAYMENT_CREATE: {
      console.log('Payment Create');
      console.log('action.data =', action.data);

      return {
        ...state,
        paymentData: action.data,

      }
    }
    case actionType.PAYMENT_GET_ORDER_DETAIL: {
      console.log('Payment get order detail');
      console.log('action.data=', action.data);

      return {
        ...state,
        orderDetail: action.data,
      }
    }


    default:
      return state;
  }
}

export default PaymentReducer;
