//  format date as DD/MM/YYYY
export const formatVietnameseDate = (date: Date | undefined): string => {
  if (date) {
    return [
      padTo2Digits(date.getDate()),
      padTo2Digits(date.getMonth() + 1),
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

  return route.toLowerCase().replaceAll(' ', '+');
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

// source: https://gist.github.com/jarvisluong/f01e108e963092336f04c4b7dd6f7e45
export function toLowerCaseNonAccentVietnamese(str: string) {
  str = str.toLowerCase();
  str = str.replace(/à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ/g, 'a');
  str = str.replace(/è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ/g, 'e');
  str = str.replace(/ì|í|ị|ỉ|ĩ/g, 'i');
  str = str.replace(/ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ/g, 'o');
  str = str.replace(/ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ/g, 'u');
  str = str.replace(/ỳ|ý|ỵ|ỷ|ỹ/g, 'y');
  str = str.replace(/đ/g, 'd');
  str = str.replace(/\u0300|\u0301|\u0303|\u0309|\u0323/g, '');
  str = str.replace(/\u02C6|\u0306|\u031B/g, ''); 
  return str.trim();
}