import React, { Component } from 'react';
import {
  withRouter,

} from 'react-router-dom'


import Header from '../header';

class paymentError extends Component {
	constructor(props){
		super(props);
		this.state={
  
		}
		this.handleButton = this.handleButton.bind(this);
  }
  
  handleButton(event) {
		event.preventDefault();
    console.log(`handleButton go back home`);
    
    this.props.history.push('/');
  }
  
  render() {
    return (
			<div className="menu-home">
				<div className="outerdiv">
					<Header />

						<div className="content payment">
							<div className="card center">
								<h2 id="center">Payment Error</h2>

                <div class="payment-alert">
                  Please try again! 👻
                </div>

								<div className="btn-container">
									{/* <button onClick={this.handleButton}>Pay for your order</button> */}
									<input type="button" className="back_button" value="Go back to Home" onClick={this.handleButton} />
								</div>
							</div>
						</div>
				</div>
			</div>
    );
  }
}

export default withRouter(paymentError);
