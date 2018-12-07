import React, { Component } from 'react';
import { connect } from 'react-redux';
import {
  withRouter,

} from 'react-router-dom'

import Moment from 'react-moment';

import '../../index.css';
import '../../stylesheets/payment.css';

import Header from '../header';
import {bindActionCreators} from "redux";
import {removeCartDataAfterExport} from './../../actions/order-actions';

class paymentSuccess extends Component {
	constructor(props){
		super(props);
		this.state={

		}
    this.handleBackButton = this.handleBackButton.bind(this);
    // this.isEmptyObject = this.isEmptyObject.bind(this);
	}

  handleBackButton(event) {
		event.preventDefault();
    console.log(`handleButton go back home`);
    
    this.props.history.push('/');
	}

  isEmptyObject(obj) {
    return !Object.keys(obj).length;
  }

  componentDidMount() {
	this.props.removeCartDataAfterExport();
  }

  render() {
    console.log('Object.keys(this.props.paymentData).length=', Object.keys(this.props.paymentData).length);

    return (
			<div className="menu-home">
				<div className="outerdiv">
					<Header />

						<div className="content payment">
							<div className="card center">
								<h2 id="center">Payment Summary</h2>

								

                {
                  
                  !Object.keys(this.props.paymentData).length ? (
                    <div className="payment-alert">
                      Invalid payment! 😱
                    </div>
                  ) : (

                    <div>
                      <h3>Thank you for your order!</h3>

                      <table>
                        <tbody>
                          <tr>
                            <td>User Id:</td>
                            <td>{this.props.paymentData.userId}</td>
                          </tr>
                          <tr>
                            <td>Order Id:</td>
                            <td>{this.props.paymentData.orderId}</td>
                          </tr>
                          <tr>
                            <td>Payment Id:</td>
                            <td>{this.props.paymentData.paymentId}</td>
                          </tr>

                          <tr>
                            <td>Payment Date:</td>
                            <td>
                              <Moment format="YYYY/MM/DD HH:mm">
                                {this.props.paymentData.paymentDate}
                              </Moment>
                            </td>
                          </tr>
                          <tr>
                            <td>Total Amount:</td>
                            <td>$ {this.props.paymentData.totalAmount}</td>
                          </tr>
                        </tbody>
                      </table>

                    <br />
                    <div className="payment-alert success">
                      Please save your payment details for your own reference. ✌️
                    </div>

                    </div>
                  )
                
                
                }

								<div className="btn-container">
									{/* <button onClick={this.handleButton}>Pay for your order</button> */}
									<input type="button" className="back_button" value="Go back to Home" onClick={this.handleBackButton} />
								</div>
							</div>
						</div>
				</div>
			</div>
    );
  }
}

const mapStateToProps = (state) => {
  return {
    paymentData: state.payment.paymentData,
  }
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators({
        removeCartDataAfterExport: removeCartDataAfterExport
    }, dispatch)
}

const connectedPaymentSuccess = withRouter(connect(mapStateToProps,mapDispatchToProps)(paymentSuccess))
export default connectedPaymentSuccess;
