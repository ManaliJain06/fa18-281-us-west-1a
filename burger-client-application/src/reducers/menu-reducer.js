import {actionTypes} from "../actions/action-types";

const restaurantMenu = {
    restaurantid: "",
    restaurantname: "",
    items:[]
}

const menu = (state = restaurantMenu, action)=>
{
    switch (action.type) {

        case actionTypes.UPDATE_MENU_LIST:
            //console.log("[menu reducer] UPDATE_MENU_LIST data",action.data);
            return Object.assign({},state,{
                restaurantId:action.data.restaurantid,
                restaurantName:action.data.restaurantname,
                items:action.data.items
            })
        default :
            return state;
    }
};

export default menu;
