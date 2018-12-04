import {combineReducers} from 'redux';
import menu from './menu-reducer';
import order from './order-reducer';

const allReducers = combineReducers({
    //insert reducer name here to combine
    menu:menu,
    order:order

});

export default allReducers;
