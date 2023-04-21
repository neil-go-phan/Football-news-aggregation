import ArticleAdmin from '@/components/admin/articles';
import ProgressBar from '@/components/processBar';
import AdminLayout from '@/layouts/AdminLayout';
import { NextPage } from 'next';
import React from 'react';

const AdminArticlesPage: NextPage = () => {
  return (
    <AdminLayout>
      <ProgressBar />
      <ArticleAdmin/>
    </AdminLayout>
  );
}

export default AdminArticlesPage;

