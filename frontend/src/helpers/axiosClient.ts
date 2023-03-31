import axios from 'axios';
import { getCookie } from 'cookies-next';
// import { _ROUTES } from './constants';

// const unProtectedRoutes = [
//   _ROUTES.NEWS_PAGE
// ]


const axiosClient = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BACKEND_DOMAIN,
  withCredentials: true,
});

axiosClient.interceptors.request.use(
  async (config: any) => {
    
    const token = getCookie('token');
    if (token) {
      config.headers = {
        'token': token,
      };
    }
    return config;
  },
  (error: any) => {
    return error;
  }
);

// axiosClient.interceptors.response.use(
//   (response) => response,
//   async (error: any) => {
//     const config = error?.config;
//     try {
//       if (
//         error.response.status === 401 &&
//         !config?.sent &&
//         error.response.data.message === 'Unauthorized access'
//       ) {
//         config.sent = true;
//         const res = await newToken();

//         if (res?.data) {
//           setCookie('access_token', res.data.accessToken);
//           config.headers = {
//             'x-access-token': res.data.accessToken,
//           };
//         }
//         return config;
//       }
//     } catch (err) {
//       deleteCookie('refresh_token');
//       deleteCookie('access_token');
//     }
//     deleteCookie('refresh_token');
//     deleteCookie('access_token');
//     if (!unProtectedRoutes.includes(window.location.pathname)) {
//       window.location.href = _ROUTES.HOME_PAGE;
//     }
//     return error;
//   }
// );


export default axiosClient;