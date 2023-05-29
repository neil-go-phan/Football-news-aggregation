import React, { forwardRef, useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { toast } from 'react-toastify';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import BarChartOnHour from './barChartOnHour';
import DatePicker from 'react-datepicker';
import BarChartOnDay from './barChartOnDay';
import { ThreeDots } from 'react-loader-spinner';
import { Button } from 'react-bootstrap';

export type ChartOnDayData = {
  time: number;
  amount_of_jobs: number;
  cronjobs: Array<CronjobDayBar>;
};

export type CronjobDayBar = {
  name: string;
  times: number;
};

export type ChartOnHourData = {
  minute: number;
  amount_of_jobs: number;
  cronjobs: Array<CronjobBar>;
};

export type CronjobBar = {
  name: string;
  start_at: string;
  end_at: string;
};

type CustomInputProps = {
  onClick: React.MouseEventHandler<HTMLButtonElement>;
  date: Date | undefined;
};

function Dashboard() {
  const [chartOnHourData, setChartOnHourData] = useState<Array<ChartOnHourData>>();
  const [chartOnDayData, setChartOnDayData] = useState<Array<ChartOnDayData>>();
  const [choosenDay, setChoosenDay] = useState<Date>();
  const [choosenHour, setChoosenHour] = useState<Date>()
  const router = useRouter();

  const formatTimeHour = (date: Date | undefined): string => {
    if (date) {
      let month = `${date.getMonth() + 1}`;
      if (date.getMonth() + 1 < 10) {
        month = `0${date.getMonth() + 1}`;
      }
      let day = `${date.getDate()}`;
      if (date.getDate() < 10) {
        day = `0${date.getDate()}`;
      }
      return `${date.getFullYear()}-${month}-${day} ${date.getHours()}`;
    }
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

  const handleChooseHour = (hour: number) => {
    if (choosenDay) {
      const temp = new Date(choosenDay)
      temp.setHours(hour)
      setChoosenHour(temp)
    }
  }

  const requestChartByHour = async (timeString: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('cronjob/cronjob-on-hour', {
        params: { time: timeString },
      });
      setChartOnHourData(data.cronjobs);
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
    const date = new Date();
    setChoosenDay(date);
    setChoosenHour(date)
  }, [router.asPath]);
  useEffect(() => {
    if (choosenHour) {
      requestChartByHour(formatTimeHour(choosenHour))
    }
  }, [choosenHour]);
  useEffect(() => {
    if (choosenDay) {
      requestChartByDay(formatTimeDay(choosenDay));
      requestChartByHour(formatTimeHour(choosenDay))
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
    <Button className="btnTriggerDate bg-success" onClick={onClick} ref={ref}>
      {date ? formatTimeDay(date) : 'Choose date'}
    </Button>
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
        <div className="adminDashboard__chooseDate--title">
          Choose date
        </div>
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
              handleChooseHour={handleChooseHour}
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
      <div className="adminDashboard__hourChart">
        <div className="adminDashboard__hourChart--title">Hour chart</div>
        {chartOnHourData && choosenHour ? (
          <div className="adminDashboard__dayChart--warper">
            <BarChartOnHour
              title={formatTimeHour(choosenHour)}
              hour={choosenHour.getHours()}
              chartData={chartOnHourData}
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
    </div>
  );
}

export default Dashboard;
