// export const paymentUrl = 'http://54.219.241.43:8000';
<<<<<<< HEAD
//export const paymentUrl = 'http://EC2Co-EcsEl-1CNG22VMRNP2G-342756940.us-west-1.elb.amazonaws.com:8000'; // running on AWS ECS
export const paymentUrl = 'http://35.247.100.252:8000'; // running on Google GKE
export const orderUrl = 'http://54.153.121.217:8000';
export const userUrl = 'http://EC2Co-EcsEl-TXPKT9YTJTZN-1941324801.us-west-1.elb.amazonaws.com:8000'
=======
//export const paymentUrl = 'http://ec2co-ecsel-hgvaibfnr6yn-1190636458.us-west-1.elb.amazonaws.com:8000'; // running on AWS ECS
// export const paymentUrl = 'http://35.247.121.55:8000'; // running on Google GKE
export const paymentUrl = process.env.API_GATEWAY || 'http://54.153.123.160:8000'; // KONG API gateway
export const orderUrl = process.env.API_GATEWAY || 'http://54.153.123.160:8000';
export const userUrl = process.env.API_GATEWAY || 'http://54.153.123.160:8000';
export const kongAPI = process.env.KONG_API || 'http://54.153.123.160:8000';
>>>>>>> 82533e02b5d5ae3a4502f7ae8f0bb1f211c73c90
