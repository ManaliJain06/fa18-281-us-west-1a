import {combineReducers} from 'redux';
import menu from './menu-reducer';
import payment from './payment-reducer';

const allReducers = combineReducers({
    //insert reducer name here to combine
    menu:menu,
    payment,
});

export default allReducers;
