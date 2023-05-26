import React, { forwardRef, useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { toast } from 'react-toastify';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import BarChartOnHour from './barChartOnHour';
import DatePicker from 'react-datepicker';
import BarChartOnDay from './barChartOnDay';
import { ThreeDots } from 'react-loader-spinner';

export type ChartOnHourData = {
  time: Date;
  amount_of_jobs: number;
  memory_usage: number;
  cronjob_names: Array<string>;
};

export type ChartOnDayData = {
  time: number;
  amount_of_jobs: number;
  cronjobs: Array<CronjobDayBar>;
};

export type CronjobDayBar = {
  name: string;
  times: number;
};

export type CronjobBar = {
  name: string;
  start_at: Date;
  end_at: Date;
};

type CustomInputProps = {
  onClick: React.MouseEventHandler<HTMLButtonElement>;
  date: Date | undefined;
};

function Dashboard() {
  const [chartOnHour, setChartOnHour] = useState<Array<ChartOnHourData>>();
  const [chartOnDayData, setChartOnDayData] = useState<Array<ChartOnDayData>>();
  const [choosenDay, setChoosenDay] = useState<Date>();
  const router = useRouter();
  const formatTimeHour = (hour: number): string => {
    // if (today) {
    //   let month = today.getUTCMonth() + 1;
    //   return `${today.getUTCFullYear()}-${
    //     today.getUTCMonth() + 1
    //   }-${today.getUTCDate()} ${hour}:${today.getUTCMinutes()}:${today.getUTCSeconds()}`;
    // }
    return '';
  };

  const formatTimeDay = (date: Date | undefined): string => {
    if (date) {
      let month = `${date.getMonth() + 1}`;
      if (date.getMonth() + 1 < 10) {
        month = `0${date.getMonth() + 1}`;
      }
      let day = `${date.getDate()}`;
      if (date.getDate() < 10) {
        day = `0${date.getDate()}`;
      }
      return `${date.getFullYear()}-${month}-${day}`;
    }
    return '';
  };

  const requestChartByHour = async (timeString: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('cronjob/cronjob-on-hour', {
        params: { time: timeString },
      });
      setChartOnHour(data.charts);
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

  const requestChartByDay = async (timeString: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('cronjob/cronjob-on-day', {
        params: { time: timeString },
      });
      setChartOnDayData(data.cronjobs);
    } catch (error) {
      toast.error('Error occurred while get chart data', {
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
    setChoosenDay(date);
  }, [router.asPath]);
  useEffect(() => {
    if (choosenDay) {
      requestChartByDay(formatTimeDay(choosenDay));
    }
  }, [choosenDay]);

  const handleOnClickChooseDay = (dateChose: Date) => {
    setChoosenDay(dateChose);
    requestChartByDay(formatTimeDay(dateChose));
  };
  // eslint-disable-next-line react/display-name
  const BtnCustomInput = forwardRef<
    HTMLButtonElement,
    {
      onClick: React.MouseEventHandler<HTMLButtonElement>;
      date: Date | undefined;
    }
  >(({ onClick, date }, ref) => (
    <button className="btnTriggerDate" onClick={onClick} ref={ref}>
      {date ? formatTimeDay(date) : 'Choose date'}
    </button>
  ));

  // eslint-disable-next-line react/display-name
  const CustomInput = React.forwardRef<HTMLButtonElement, CustomInputProps>(
    ({ onClick, date }, ref) => (
      <BtnCustomInput onClick={onClick} ref={ref} date={date} />
    )
  );
  return (
    <div className="adminDashboard">
      <div className="adminDashboard__title">Dashboard</div>
      <div className="adminDashboard__chooseDate">
        <DatePicker
          selected={choosenDay}
          onChange={(date: Date) => handleOnClickChooseDay(date)}
          customInput={<CustomInput onClick={() => {}} date={choosenDay} />}
        />
      </div>
      <div className="adminDashboard__dayChart">
        <div className="adminDashboard__dayChart--title">Day chart</div>
        {chartOnDayData ? (
          <div className="adminDashboard__dayChart--warper">
            <BarChartOnDay
              title={formatTimeDay(choosenDay)}
              chartData={chartOnDayData}
            />
          </div>
        ) : (
          <div className="adminDashboard__dayChart--loading">
            <ThreeDots
              height="50"
              width="50"
              radius="9"
              color="#4fa94d"
              ariaLabel="three-dots-loading"
              visible={true}
            />
          </div>
        )}
      </div>
      {/* <div className="adminDashboard__hourChart">
        <div className="adminDashboard__hourChart--title">Hour chart</div>
        {chartOnHour ? (
          <div className="adminDashboard__dayChart--warper">
            <BarChartOnHour
              title={formatTimeDay(choosenDay)}
              chartData={chartOnHour}
            />
          </div>
        ) : (
          <div className="adminDashboard__dayChart--loading">
            <ThreeDots
              height="50"
              width="50"
              radius="9"
              color="#4fa94d"
              ariaLabel="three-dots-loading"
              visible={true}
            />
          </div>
        )}
      </div> */}
    </div>
  );
}

export default Dashboard;
