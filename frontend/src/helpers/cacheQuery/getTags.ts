import { useQuery } from 'react-query';
import axiosClient from '../axiosClient';
import { toast } from 'react-toastify';
import { ERROR_POPUP_ADMIN_TIME } from '../constants';
// 5 min
const TAGS_STALE_TIME = 5 * (60 * 1000)
// 10 min
const TAGS_CACHE_TIME = 10 * (60 * 1000)

export type Tags = Array<string> | undefined

const callAPIGetTags = async ():Promise<Tags> => {
  try {
    const { data } = await axiosClient.get('tags/list');
    return data.tags.tags
  } catch (error) {
    toast.error('Error occurred while geting tags', {
      position: 'top-right',
      autoClose: ERROR_POPUP_ADMIN_TIME,
      hideProgressBar: false,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
      progress: undefined,
      theme: 'light',
    });
  }
}

const useGetTags = () => {
  const getTagsQuery = useQuery('tags', callAPIGetTags, {
    cacheTime: TAGS_CACHE_TIME,
    refetchOnWindowFocus: false,
    staleTime: TAGS_STALE_TIME,
  });

return getTagsQuery;
}

export default useGetTags;