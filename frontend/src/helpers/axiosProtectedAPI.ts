import axios from 'axios';
import { getCookie, deleteCookie } from 'cookies-next';
import { _ROUTES } from './constants';

const unProtectedRoutes = [_ROUTES.NEWS_PAGE, _ROUTES.MATCH_DETAIL_PAGE];

const axiosProtectedAPI = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BACKEND_DOMAIN,
  withCredentials: true,
});

axiosProtectedAPI.interceptors.request.use(
  async (config: any) => {
    const token = getCookie('token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error: any) => {
    return error;
  }
);

axiosProtectedAPI.interceptors.response.use(
  (response) => response,
  async (error: any) => {
    const config = error?.config;
    if (
      error.response.status === 401 &&
      !config?.sent &&
      error.response.data.message === 'Unauthorized access'
    ) {
      deleteCookie('token');
      if (!unProtectedRoutes.includes(window.location.pathname)) {
        window.location.href = _ROUTES.ADMIN_LOGIN;
      }
    }

    return config;
  }
);

export default axiosProtectedAPI;
