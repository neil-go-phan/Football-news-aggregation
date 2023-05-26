import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { toast } from 'react-toastify';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import BarChartOnHour from './barChartOnHour';

export type ChartOnHourData = {
  time: Date,
  amount_of_jobs: number,
  memory_usage: number,
  cronjob_names: Array<string>
}

export type CronjobBar = {
  name: string;
  start_at: Date;
  end_at: Date;
}


function Dashboard() {
  const [chartOnHour, setChartOnHourDate] = useState<Array<ChartOnHourData>>();
  const [today, setToday] = useState<Date>()
  const router = useRouter();
  const formatTimeHour = (hour: number):string => {
    if (today) {
      let month = today.getUTCMonth() + 1
      return `${today.getUTCFullYear()}-${today.getUTCMonth()+1}-${today.getUTCDate()} ${hour}:${today.getUTCMinutes()}:${today.getUTCSeconds()}`
    }
    return ""
  }

  const requestChartByHour = async (timeString: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('cronjob/cronjob-on-hour', {
        params: {time: timeString}
      });
      setChartOnHourDate(data.charts);
    } catch (error) {
      toast.error('Error occurred while get list cronjob', {
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
  };

  useEffect(() => {
    var date = new Date();
    setToday(date);
  }, [router.asPath]);
  useEffect(() => {
    if (today) {
      requestChartByHour(formatTimeHour(8))
    }
  }, [today]);
  return (
    <div className="adminDashboard">
      <div className="adminDashboard__title">Dashboard</div>
      {chartOnHour ? <BarChartOnHour title="Test" chartData={chartOnHour}/> : <></>}
      
    </div>
  );
}

export default Dashboard;
