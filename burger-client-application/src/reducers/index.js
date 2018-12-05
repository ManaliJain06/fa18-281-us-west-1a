import {combineReducers} from 'redux';
import menu from './menu-reducer';
import restaurant from './restaurant-reducer';

const allReducers = combineReducers({
    //insert reducer name here to combine
    menu:menu,
    restaurant:restaurant

});

export default allReducers;
