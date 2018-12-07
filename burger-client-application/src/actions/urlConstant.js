// export const paymentUrl = 'http://54.219.241.43:8000';
//export const paymentUrl = 'http://ec2co-ecsel-hgvaibfnr6yn-1190636458.us-west-1.elb.amazonaws.com:8000'; // running on AWS ECS
// export const paymentUrl = 'http://35.247.121.55:8000'; // running on Google GKE
export const paymentUrl = process.env.REACT_APP_API_GATEWAY || 'https://hszrvgkm85.execute-api.us-west-1.amazonaws.com/burger'; // AWS API gateway
export const orderUrl = process.env.REACT_APP_API_GATEWAY || 'http://54.153.121.217:8000';
export const userUrl = process.env.REACT_APP_API_GATEWAY || 'http://EC2Co-EcsEl-XYD9SO7HL5E1-1239896239.us-west-1.elb.amazonaws.com:8000';
