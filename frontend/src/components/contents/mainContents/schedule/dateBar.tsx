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
import { Schedules } from '.';
import { formatVietnameseDate, formatISO8601Date } from '@/helpers/format';

const WEEKDAYS = [
  'Chủ nhật',
  'Thứ 2',
  'Thứ 3',
  'Thứ 4',
  'Thứ 5',
  'Thứ 6',
  'Thứ 7',
];

type CustomInputProps = {
  onClick: React.MouseEventHandler<HTMLButtonElement>;
};

type Props = {
  // eslint-disable-next-line no-unused-vars
  handleSchedule: (data: Schedules) => void;
};

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
    requestScheduleDate(dateChose);
  };

  const requestScheduleDate = async (date: Date) => {
    try {
      const { data } = await axiosClient.get('schedules/on-day?', {
        // eslint-disable-next-line camelcase
        params: { date: formatISO8601Date(date) },
      });
      handleSchedule(data.schedules);
    } catch (error) {
      toast.error(`Error occurred while get schedule on ${formatVietnameseDate(date)}`, {
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
  };

  useEffect(() => {
    const today = new Date();
    handleOnClickChooseDay(today);

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [route.asPath]);

  // eslint-disable-next-line react/display-name
  const BtnCustomInput = forwardRef<
    HTMLButtonElement,
    { onClick: React.MouseEventHandler<HTMLButtonElement> }
  >(({ onClick }, ref) => (
    <button className="btnTriggerDate" onClick={onClick} ref={ref}>
      Chọn ngày
    </button>
  ));
  const CustomInput: React.FunctionComponent<CustomInputProps> = ({
    onClick,
  }) => <BtnCustomInput onClick={onClick} />;
  return (
    <div className="schedule__dateBar d-flex px-3">
      <div className="schedule__dateBar--weekdays d-flex col-10">
        <div
          className={
            getFollowingDays(0) === formatVietnameseDate(date)
              ? 'weekdays--item col-2 active'
              : 'weekdays--item col-2'
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
              ? 'weekdays--item col-2 active'
              : 'weekdays--item col-2'
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
              ? 'weekdays--item col-2 active'
              : 'weekdays--item col-2'
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
              ? 'weekdays--item col-2 active'
              : 'weekdays--item col-2'
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
              ? 'weekdays--item col-2 active'
              : 'weekdays--item col-2'
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
              ? 'weekdays--item col-2 active'
              : 'weekdays--item col-2'
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
      <div className="schedule__dateBar--chooseDay col-2">
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
