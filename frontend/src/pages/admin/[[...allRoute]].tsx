import AdminComponent from '@/components/admin';
import ArticleAdmin from '@/components/admin/articles';
import ProgressBar from '@/components/processBar';
import { _ROUTES } from '@/helpers/constants';
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

