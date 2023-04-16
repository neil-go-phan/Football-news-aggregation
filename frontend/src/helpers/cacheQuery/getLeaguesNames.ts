import { useQuery } from "react-query";
import axiosClient from "../axiosClient";
import { toast } from "react-toastify";

const LEAGUES_NAME_STALE_TIME = 5 * (60 * 1000)// 5 min
const LEAGUES_NAME_CACHE_TIME = 10 * (60 * 1000)// 10 min

export type LeaguesNames = Array<string> | undefined

const callAPIGetLeaguesNames = async ():Promise<LeaguesNames> => {
  try {
    const { data } = await axiosClient.get('leagues/list-name');
    return data.leagues
  } catch (error) {
    toast.error('Error occurred while geting leagues', {
      position: 'top-right',
      autoClose: 3000,
      hideProgressBar: false,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
      progress: undefined,
      theme: 'light',
    });
  }
}

const useGetLeaguesName = () => {
  const getLeaguesNameQuery = useQuery("leagues-names", callAPIGetLeaguesNames, {
    cacheTime: LEAGUES_NAME_CACHE_TIME,
    refetchOnWindowFocus: false,
    staleTime: LEAGUES_NAME_STALE_TIME,
  });

return getLeaguesNameQuery;
}

export default useGetLeaguesName;