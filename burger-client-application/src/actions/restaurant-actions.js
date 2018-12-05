import {actionTypes} from "./action-types";

export function restaurantList(data) {
    return {
        type: actionTypes.RESTAURANT_LIST,
        data
    }
}