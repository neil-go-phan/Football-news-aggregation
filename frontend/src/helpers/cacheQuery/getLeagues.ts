import { useQuery } from "react-query";
import axiosClient from "../axiosClient";
import { toast } from "react-toastify";

const LEAGUES_STALE_TIME = 5 * (60 * 1000)// 5 min
const LEAGUES_CACHE_TIME = 10 * (60 * 1000)// 10 min

export type Leagues = Array<string> | undefined

const callAPIGetLeagues = async ():Promise<Leagues> => {
  try {
    const { data } = await axiosClient.get('leagues/list');
    return data.leagues.leagues
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

const useGetLeagues = () => {
  const getLeaguesQuery = useQuery("leagues", callAPIGetLeagues, {
    cacheTime: LEAGUES_CACHE_TIME,
    refetchOnWindowFocus: false,
    staleTime: LEAGUES_STALE_TIME,
  });

return getLeaguesQuery;
}

export default useGetLeagues;