//  format date as DD/MM/YYYY
export const formatVietnameseDate = (date: Date | undefined): string => {
  if (date) {
    return [
      padTo2Digits(date.getDate()),
      padTo2Digits(date.getMonth() + 1),
      date.getFullYear(),
    ].join('/');
  }
  return '';
};

const padTo2Digits = (num: number): string => {
  return num.toString().padStart(2, '0');
};

export const formatRoute = (league: string) => {
  const route = league
    .normalize('NFD')
    .replace(/[\u0300-\u036f]/g, '')
    .replace(/đ/g, 'd')
    .replace(/Đ/g, 'D');

  return route.replaceAll(' ', '+');
};

export const formatISO8601Date = (date: Date): string => {
  let month, year, day;
  year = date.getFullYear();
  if (date.getMonth() + 1 < 10) {
    month = `0${date.getMonth() + 1}`;
  } else {
    month = date.getMonth() + 1;
  }
  if (date.getDate() < 10) {
    day = `0${date.getDate()}`;
  } else {
    day = date.getDate();
  }

  return `${year}-${month}-${day}`;
};