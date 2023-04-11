import axiosClient from '@/helpers/axiosClient'
import { useRouter } from 'next/router'
import React, { useEffect } from 'react'
import { toast } from 'react-toastify'

function Detail() {
  const router = useRouter()

  

  useEffect(() => {
    const requestGetMatchDetail =async () => {
      const {date, club_1, club_2} = router.query
      try {
        const { data } = await axiosClient.get('match-detail/get', {
          // eslint-disable-next-line camelcase
          params: { date: date, club_1: club_1, club_2: club_2},
        });
        console.log(data)
        // handleSchedule(data.schedules);
      } catch (error) {
        toast.error(`Error occurred while get match detail`, {
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
    requestGetMatchDetail()
  }, [router.asPath])
  
  return (
    <div>Detail</div>
  )
}

export default Detail


