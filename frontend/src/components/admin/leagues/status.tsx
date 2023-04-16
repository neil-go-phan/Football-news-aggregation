import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useState } from 'react';
import { Form } from 'react-bootstrap';
import { toast } from 'react-toastify';

type Props = {
  active: boolean;
  leagueName: string;
};

const Status: React.FC<Props> = (props: Props) => {
  const [isChecked, setIsChecked] = useState<boolean>(props.active);

  const onSwitchAction = () => {
    requestChangeLeagueActive();
  };

  const requestChangeLeagueActive = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('leagues/change-status', {
        params: { league: props.leagueName },
      });
      if (!data.success) {
        throw 'change fail';
      }
      toast.success('Change league status success', {
        position: 'top-right',
        autoClose: 3000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
      setIsChecked(!isChecked);
      if (data.status_active) {
        requestArticleCrawler();
      }
    } catch (error) {
      toast.error('Error occurred while request to change league status', {
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

  const requestArticleCrawler = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('article/signal-crawler', {
        params: { league: props.leagueName },
      });
      if (!data.success) {
        throw 'change fail';
      }
      toast.success('Request to crawl articles success', {
        position: 'top-right',
        autoClose: 3000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
    } catch (error) {
      toast.error('Error occurred while request to crawl articles', {
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

  return (
    <Form.Check
      type='switch'
      onChange={onSwitchAction}
      checked={isChecked}
      label={isChecked ? 'Active' : 'Disable'}
    />
  );
};

export default Status;
