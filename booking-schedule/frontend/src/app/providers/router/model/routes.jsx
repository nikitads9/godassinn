import Main from '../../../../pages/Main/Main';
import Auth from '../../../../pages/Auth/Auth';
import Profile from '../../../../pages/Profile/Profile';

export const routes = [
  {
    path: '/',
    element: <Main />,
  },
  {
    path: '/auth',
    element: <Auth />,
  },
  {
    path: '/profile',
    element: <Profile />,
    authOnly: true,
  },
];
