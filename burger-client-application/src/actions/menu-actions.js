import {actionTypes} from "./action-types";

export function updateMenuList(data) {
    return {
        type: actionTypes.UPDATE_MENU_LIST,
        data
    }
}
