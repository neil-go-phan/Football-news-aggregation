import axios from 'axios';

const axiosClient = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BACKEND_DOMAIN,
  withCredentials: true,
});

export default axiosClient;