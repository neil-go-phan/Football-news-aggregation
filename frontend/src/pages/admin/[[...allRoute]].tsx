import AdminComponent from '@/components/admin';
import ProgressBar from '@/components/processBar';
import AdminLayout from '@/layouts/AdminLayout';
import React from 'react';

export default function Admin(){
  return (
    <AdminLayout>
      <ProgressBar />
      <AdminComponent />
    </AdminLayout>
  );
}

