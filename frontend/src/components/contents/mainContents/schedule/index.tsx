import React, { useState } from 'react'
import DateBar from './dateBar'
import ScheduleContent from './scheduleContent';

export type Schedules = {
  date: string;
  date_with_weekday: string;
  schedule_on_leagues: Array<ScheduleOnLeague>;
} | undefined;

export type ScheduleOnLeague = {
  league_name: string;
  matchs: Array<Match>;
};

export type Match = {
  time: string;
  round: string;
  club_1: Club;
  club_2: Club;
  scores: string;
  match_detail_link: string;
};

export type Club = {
  name: string;
  logo: string;
};

export default function Schedule() {
  const [schedule, setSchedule] = useState<Schedules>();
  const handleSchedule = (data: Schedules) => {
    setSchedule(data)
  }
  return (
    <div className='schedule'>
      <DateBar handleSchedule={handleSchedule}/>
      <ScheduleContent schedule={schedule}/>
    </div>
  )
}
