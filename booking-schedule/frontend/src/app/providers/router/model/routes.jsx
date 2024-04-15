import Main from '../../../../pages/Main/Main';
import Auth from '../../../../pages/Auth/Auth';
import Profile from '../../../../pages/Profile/Profile';
import Bookings from '../../../../pages/Bookings/Bookings';
import Booking from '../../../../pages/Booking/Booking';

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
  {
    path: '/bookings',
    element: <Bookings />,
    authOnly: true,
  },
  {
    path: '/booking',
    element: <Booking />,
    authOnly: true,
  },
];
