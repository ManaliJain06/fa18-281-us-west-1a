const burgerUser = {
    user: {},
    isLoggedIn: false,
};

const menu = (state = burgerUser, action)=>
{
    switch (action.type) {

        case "LOGIN":
            console.log("[menu reducer] UPDATE_MENU_LIST data",action.data);
            return Object.assign({},state,{
                restaurantId:action.data.restaurantId,
                restaurantName:action.data.restaurantName,
                items:action.data.items
            });
        default :
            return state;
    }
};

export default menu;