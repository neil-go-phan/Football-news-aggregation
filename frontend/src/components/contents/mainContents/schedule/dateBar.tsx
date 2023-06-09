import axiosClient from '@/helpers/axiosClient';
import { useRouter } from 'next/router';
import React, {
  FunctionComponent,
  forwardRef,
  useEffect,
  useState,
} from 'react';
import { toast } from 'react-toastify';
import DatePicker from 'react-datepicker';
import { formatVietnameseDate, formatISO8601Date } from '@/helpers/format';
import { Schedules } from '.';
import { ERROR_POPUP_USER_TIME } from '@/helpers/constants';

const WEEKDAYS = [
  'Sunday',
  'Monday',
  'Tuesday',
  'Wednesday',
  'Thursday',
  'Friday',
  'Saturday',
];

type CustomInputProps = {
  onClick: React.MouseEventHandler<HTMLButtonElement>;
};

type Props = {
  // eslint-disable-next-line no-unused-vars
  handleSchedule: (data: Schedules) => void;
};

// eslint-disable-next-line react/display-name
const BtnCustomInput = forwardRef<
  HTMLButtonElement,
  { onClick: React.MouseEventHandler<HTMLButtonElement> }
>(({ onClick }, ref) => (
  <button className="btnTriggerDate" onClick={onClick} ref={ref}>
    Choose date
  </button>
));
// eslint-disable-next-line react/display-name
const CustomInput = React.forwardRef<HTMLButtonElement, CustomInputProps>(
  ({ onClick }, ref) => <BtnCustomInput onClick={onClick} ref={ref} />
);

const DateBar: FunctionComponent<Props> = ({ handleSchedule }) => {
  const route = useRouter();
  const [date, setDate] = useState<Date>();

  const getFollowingDays = (day: number): string => {
    let today = new Date();
    today.setDate(today.getDate() + day);
    return formatVietnameseDate(today);
  };

  const getWeeksDay = (day: number): string => {
    let today = new Date();
    today.setDate(today.getDate() + day);
    return WEEKDAYS[today.getDay()];
  };

  const handleOnClickChooseDay = (dateChose: Date) => {
    setDate(dateChose);
    const { league } = route.query;
    if (league === 'Tin tức bóng đá') {
      requestAllScheduleDate(dateChose);
    } else {
      requestScheduleOnLeagueDate(dateChose, league);
    }
  };

  const requestAllScheduleDate = async (date: Date) => {
    try {
      const { data } = await axiosClient.get('schedules/all-league-on-day', {
        // eslint-disable-next-line camelcase
        params: { date: formatISO8601Date(date) },
      });
      handleSchedule(data.schedules);
    } catch (error) {
      toast.error(
        `Error occurred while get schedule on ${formatVietnameseDate(date)}`,
        {
          position: 'top-right',
          autoClose: ERROR_POPUP_USER_TIME,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: 'light',
        }
      );
    }
  };

  const requestScheduleOnLeagueDate = async (
    date: Date,
    league: string | string[] | undefined
  ) => {
    try {
      const { data } = await axiosClient.get('schedules/league-on-day?', {
        // eslint-disable-next-line camelcase
        params: { date: formatISO8601Date(date), league: league },
      });
      handleSchedule(data.schedules);
    } catch (error) {
      toast.error(
        `Error occurred while get schedule on ${formatVietnameseDate(date)}`,
        {
          position: 'top-right',
          autoClose: ERROR_POPUP_USER_TIME,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: 'light',
        }
      );
    }
  };

  useEffect(() => {
    const today = new Date();
    handleOnClickChooseDay(today);

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [route.asPath]);

  return (
    <div className="schedule__dateBar d-flex px-3">
      <div className="schedule__dateBar--weekdays d-flex col-9">
        <div
          className={
            getFollowingDays(0) === formatVietnameseDate(date)
              ? 'weekdays--item col-4 col-sm-3 col-lg-2 col-md-2 active'
              : 'weekdays--item col-4 col-sm-3 col-lg-2 col-md-2'
          }
          onClick={() => {
            let today = new Date();
            today.setDate(today.getDate() + 0);
            handleOnClickChooseDay(today);
          }}
        >
          <p className="followingDay">{getFollowingDays(0)}</p>
          <p>{getWeeksDay(0)}</p>
        </div>
        <div
          className={
            getFollowingDays(1) === formatVietnameseDate(date)
              ? 'weekdays--item col-4 col-sm-3 col-lg-2 col-md-2 active'
              : 'weekdays--item col-4 col-sm-3 col-lg-2 col-md-2'
          }
          onClick={() => {
            let today = new Date();
            today.setDate(today.getDate() + 1);
            handleOnClickChooseDay(today);
          }}
        >
          <p className="followingDay">{getFollowingDays(1)}</p>
          <p>{getWeeksDay(1)}</p>
        </div>
        <div
          className={
            getFollowingDays(2) === formatVietnameseDate(date)
              ? 'weekdays--item col-4 col-sm-3 col-lg-2 col-md-2 active'
              : 'weekdays--item col-4 col-sm-3 col-lg-2 col-md-2'
          }
          onClick={() => {
            let today = new Date();
            today.setDate(today.getDate() + 2);
            handleOnClickChooseDay(today);
          }}
        >
          <p className="followingDay">{getFollowingDays(2)}</p>
          <p>{getWeeksDay(2)}</p>
        </div>
        <div
          className={
            getFollowingDays(3) === formatVietnameseDate(date)
              ? 'weekdays--item col-lg-2 col-sm-3 col-md-2 d-none d-sm-block active'
              : 'weekdays--item col-lg-2 col-sm-3 col-md-2 d-none d-sm-block '
          }
          onClick={() => {
            let today = new Date();
            today.setDate(today.getDate() + 3);
            handleOnClickChooseDay(today);
          }}
        >
          <p className="followingDay">{getFollowingDays(3)}</p>
          <p>{getWeeksDay(3)}</p>
        </div>
        <div
          className={
            getFollowingDays(4) === formatVietnameseDate(date)
              ? 'weekdays--item col-lg-2 col-md-2 d-none d-md-block active'
              : 'weekdays--item col-lg-2 col-md-2 d-none d-md-block '
          }
          onClick={() => {
            let today = new Date();
            today.setDate(today.getDate() + 4);
            handleOnClickChooseDay(today);
          }}
        >
          <p className="followingDay">{getFollowingDays(4)}</p>
          <p>{getWeeksDay(4)}</p>
        </div>
        <div
          className={
            getFollowingDays(5) === formatVietnameseDate(date)
              ? 'weekdays--item col-lg-2 col-md-2 d-none d-md-block active'
              : 'weekdays--item col-lg-2 col-md-2 d-none d-md-block '
          }
          onClick={() => {
            let today = new Date();
            today.setDate(today.getDate() + 5);
            handleOnClickChooseDay(today);
          }}
        >
          <p className="followingDay">{getFollowingDays(5)}</p>
          <p>{getWeeksDay(5)}</p>
        </div>
      </div>
      <div className="schedule__dateBar--chooseDay col-3">
        <DatePicker
          selected={date}
          onChange={(date: Date) => handleOnClickChooseDay(date)}
          customInput={<CustomInput onClick={() => {}} />}
        />
      </div>
    </div>
  );
};

export default DateBar;
