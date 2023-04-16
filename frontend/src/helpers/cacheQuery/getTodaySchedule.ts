// import { useQuery } from "react-query";
// import axiosClient from "../axiosClient";
// import { toast } from "react-toastify";
// import { formatISO8601Date, formatVietnameseDate } from "../format";

// const SCHEDULE_STALE_TIME = 5 * (60 * 1000)// 5 min
// const SCHEDULE_CACHE_TIME = 10 * (60 * 1000)// 10 min

// export type Schedules = {
//   date: string;
//   date_with_weekday: string;
//   schedule_on_leagues: Array<ScheduleOnLeague>;
// } | undefined;

// export type ScheduleOnLeague = {
//   league_name: string;
//   matchs: Array<Match>;
// };

// export type Match = {
//   time: string;
//   round: string;
//   club_1: Club;
//   club_2: Club;
//   scores: string;
//   match_detail_link: string;
// };

// export type Club = {
//   name: string;
//   logo: string;
// };

// const callAPIGetTodayShedule = async ():Promise<Schedules> => {
//   const today = new Date();
//   try {
//     const { data } = await axiosClient.get('schedules/on-day?', {
//       // eslint-disable-next-line camelcase
//       params: { date: formatISO8601Date(today) },
//     });
//     return data.schedules
//   } catch (error) {
//     toast.error(
//       `Error occurred while get schedule on ${formatVietnameseDate(today)}`,
//       {
//         position: 'top-right',
//         autoClose: 3000,
//         hideProgressBar: false,
//         closeOnClick: true,
//         pauseOnHover: true,
//         draggable: true,
//         progress: undefined,
//         theme: 'light',
//       }
//     );
//   }
// }

// const useGetTodaySchedule = () => {
//   const getTodayScheduleQuery = useQuery("today-schedule", callAPIGetTodayShedule, {
//     cacheTime: SCHEDULE_CACHE_TIME,
//     refetchOnWindowFocus: false,
//     staleTime: SCHEDULE_STALE_TIME,
//   });

// return getTodayScheduleQuery;
// }

// export default useGetTodaySchedule;
