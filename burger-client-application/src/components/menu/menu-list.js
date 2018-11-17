/*
	UI Component to show all menu items
*/

import React, {Component} from 'react';

class Menu extends Component{

render(){
  console.log("[Menu] render ");
  return (
      <div className="menu-home">
          <h1> Welcome to burger menu page !!! </h1>
      </div>
  )
}
}

export default Menu;
