import {actionTypes} from "./action-types";

export function updateCart(data) {
    return {
        type: actionTypes.UPDATE_CART,
        data
    }
}

export function removeCartDataAfterExport() {
    return {
        type: actionTypes.REMOVE_CART_DATA_AFTER_PAYMENT
    }
}