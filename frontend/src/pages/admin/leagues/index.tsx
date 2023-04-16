import AdminLeagues from '@/components/admin/leagues';
import ProgressBar from '@/components/processBar';
import AdminLayout from '@/layouts/AdminLayout';
import { NextPage } from 'next';
import React from 'react';

const AdminLeaguesPage: NextPage = () => {
  return (
    <AdminLayout>
      <ProgressBar />
      <AdminLeagues />
    </AdminLayout>
  );
}

export default AdminLeaguesPage;

