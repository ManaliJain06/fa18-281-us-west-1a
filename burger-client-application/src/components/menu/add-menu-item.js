/*
	UI Component to add an item in the menu
*/

import React, {Component} from 'react';

class AddItem extends Component{

render(){
  console.log("[AddItem] render ");
  return (
      <div className="menu-home">
          <h1> Welcome to Add Item page !!! </h1>
                          <div className="sign-up-form">
                    <form>
                        <div className="sign-up-container">
                            <input className="signup-input-text" type="text" placeholder="Enter Username" name="email"
                                   onChange={(event) => {
                                     this.setState({
                                         ...this.state,
                                         userdetail: {...this.state.userdetail,email:event.target.value}
                                     })
                                   }}/>
                            <input className="signup-input-text" type="password" placeholder="Enter Password" name="password"
                                   onChange={(event) => {
                                       this.setState({
                                           ...this.state,
                                           userdetail:{...this.state.userdetail,password:event.target.value}
                                       })
                                   }}/>
                            <input className="signup-input-text" type="text" placeholder="Enter First Name" name="firstname"
                                   onChange={(event) => {
                                       this.setState({
                                           ...this.state,
                                           userdetail:{...this.state.userdetail,firstname:event.target.value}
                                       })
                                   }}/>
                            <input className="signup-input-text" type="text" placeholder="Enter Second Name" name="lastname"
                                   onChange={(event) => {
                                       this.setState({
                                           ...this.state,
                                           userdetail:{...this.state.userdetail,lastname:event.target.value}
                                       })
                                   }}/>
                            {<input className="signup-input-text" type="button" id="button"
                                   onClick={()=>this.handleSignUp()} value="Sign Up"/>}

                            
                        </div>
                    </form>
                </div>
      </div>
  )
}
}

export default AddItem;