import {actionTypes} from "../actions/action-types";

const restaurantList = {

}

const restaurant = (state = restaurantList, action)=>
{
    switch (action.type) {

        case actionTypes.RESTAURANT_LIST:
            console.log("[restaurant list",action.data);
            return Object.assign({},state,{
                restaurantList:action.data
            })
        default :
            return state;
    }
};

export default restaurant;