import {combineReducers} from 'redux';
import menu from './menu-reducer';
import order from './order-reducer';
import payment from './payment-reducer';

const allReducers = combineReducers({
    //insert reducer name here to combine
    menu:menu,
    order:order,
    payment,

});

export default allReducers;
