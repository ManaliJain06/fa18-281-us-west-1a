import {actionTypes} from "./action-types";

export function updateCart(data) {
    return {
        type: actionTypes.UPDATE_CART,
        data
    }
}
