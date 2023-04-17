import AdminTags from '@/components/admin/tags';
import ProgressBar from '@/components/processBar';
import AdminLayout from '@/layouts/AdminLayout';
import { NextPage } from 'next';
import React from 'react';

const AdminTagsPage: NextPage = () => {
  return (
    <AdminLayout>
      <ProgressBar />
      <AdminTags />
    </AdminLayout>
  );
}

export default AdminTagsPage;

