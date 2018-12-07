// export const paymentUrl = 'http://54.219.241.43:8000';
//export const paymentUrl = 'http://ec2co-ecsel-hgvaibfnr6yn-1190636458.us-west-1.elb.amazonaws.com:8000'; // running on AWS ECS
// export const paymentUrl = 'http://35.247.121.55:8000'; // running on Google GKE
export const paymentUrl = process.env.API_GATEWAY || 'http://54.153.123.160:8000'; // KONG API gateway
export const orderUrl = process.env.API_GATEWAY || 'http://54.153.123.160:8000';
export const userUrl = process.env.API_GATEWAY || 'http://54.153.123.160:8000';
export const kongAPI = process.env.KONG_API || 'http://54.153.123.160:8000';
