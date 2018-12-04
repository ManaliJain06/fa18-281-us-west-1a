import React, {Component} from 'react';
import '../../index.css';

class ListRestaurant extends Component{


  render(){
    console.log("[Header] render: ";
    return (
        <div className="menu-home">
            {this.displayMenu()}
        </div>
    )
  }
}
