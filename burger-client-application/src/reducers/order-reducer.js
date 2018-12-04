import {actionTypes} from "../actions/action-types";

const orderCart = {
    orderCount:0,
    orderId:"",
    items:[]
}

const order = (state = orderCart, action)=>
{
    switch (action.type) {

        case actionTypes.UPDATE_CART:
            console.log("[order reducer] UPDATE_CART data",action.data);
            return Object.assign({},state,{
                orderCount:action.data.items?action.data.items.length:state.orderCount,
                orderId:action.data.orderId,
                items:action.data.items
            })
        default :
            return state;
    }
};

export default order;
