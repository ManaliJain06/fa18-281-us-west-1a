import {actionTypes} from "./action-types";

export function userOrdersList(data) {
    return {
        type: actionTypes.USER_LIST,
        data
    }
}