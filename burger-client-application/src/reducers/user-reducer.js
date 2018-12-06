import {actionTypes} from "../actions/action-types";

const ordersList = {

};

const user = (state = ordersList, action)=>
{
    switch (action.type) {

        case actionTypes.USER_LIST:
            console.log("[User] list",action.data);
            return Object.assign({},state,{
                ordersList:action.data
            });
        default :
            return state;
    }
};

export default user;