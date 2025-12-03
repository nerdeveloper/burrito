import React, { useContext } from 'react';
import { Outlet, Navigate } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';

import { ThemeContext } from '@/contexts/ThemeContext';
import { getUserInfo, UserInfo } from '@/clients/auth/client';

import NavigationBar from '@/components/navigation/NavigationBar';

const isAuthDisabled = import.meta.env.VITE_DISABLE_AUTH === 'true';

const Layout: React.FC = () => {
  const { theme } = useContext(ThemeContext);

  const {
    isLoading,
    isError,
    data: user
  } = useQuery<UserInfo, Error>({
    queryKey: ['userInfo'],
    queryFn: getUserInfo,
    retry: false,
    refetchOnWindowFocus: false,
    enabled: !isAuthDisabled
  });

  if (!isAuthDisabled && isLoading) {
    return (
      <div
        className={`
          flex
          items-center
          justify-center
          h-screen
          w-screen
          ${theme === 'light' ? 'bg-primary-100' : 'bg-nuances-black'}
        `}
      >
        <div
          className={`
            text-lg
            ${theme === 'light' ? 'text-nuances-black' : 'text-nuances-white'}
          `}
        >
          Loading...
        </div>
      </div>
    );
  }

  if (!isAuthDisabled && (isError || !user)) {
    return <Navigate to="/login" replace />;
  }

  return (
    <div
      className={`
        flex
        ${theme === 'light' ? 'bg-primary-100' : 'bg-nuances-black'}
      `}
    >
      <NavigationBar variant={theme} />
      <Outlet />
    </div>
  );
};

export default Layout;
