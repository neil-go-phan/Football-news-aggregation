import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import React from 'react';
import { Form } from 'react-bootstrap';
import { toast } from 'react-toastify';

type Props = {
  active: boolean;
  leagueName: string;
  handleSwitch:  () => void
};

const Status: React.FC<Props> = (props: Props) => {
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
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
      if (data.status_active) {
        requestArticleCrawler();
      }
      props.handleSwitch()
    } catch (error) {
      toast.error('Error occurred while request to change league status', {
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
        autoClose: ERROR_POPUP_ADMIN_TIME,
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

  return (
    <Form.Check
      type='switch'
      onChange={onSwitchAction}
      checked={props.active}
      label={props.active ? 'Active' : 'Inactive'}
    />
  );
};

export default Status;
