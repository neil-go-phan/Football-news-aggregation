import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import React, { useState } from 'react';
import { Button } from 'react-bootstrap';
import { toast } from 'react-toastify';
import Popup from 'reactjs-popup';
import EditTimeModal from './editTimeModal';

type Props = {
  url: string;
  name: string;
  runEveryMinOld: number;
  handleChangeTime: () => void;
  // handleDeactive: () => void
};

const CronjobActions: React.FC<Props> = (props: Props) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const handleIsModalClose = (newTime: number) => {
    setIsModalOpen(false);
    requestChangeTime(newTime);
  };
  console.log("url", props.url)
  const requestChangeTime = async (newTime: number) => {
    try {
      const { data } = await axiosProtectedAPI.post('crawler/change-time', {
        name: props.name,
        url: props.url,
        run_every_min_old: props.runEveryMinOld,
        run_every_min_new: newTime,
      });
      if (!data.success) {
        throw 'Change time cronjob fail';
      }
      toast.success('Change time cronjob success', {
        position: 'top-right',
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
      props.handleChangeTime();
    } catch (error) {
      toast.error('Error occurred while Change time cronjob', {
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
    <>
      <Button variant="primary" onClick={() => setIsModalOpen(!isModalOpen)}>
        Edit time
      </Button>
      <Popup modal open={isModalOpen} onClose={() => setIsModalOpen(false)}>
        <EditTimeModal handleIsModalClose={handleIsModalClose} />
      </Popup>
    </>
  );
};

export default CronjobActions;
