import { Navigate } from 'react-router-dom';

const RequireAuth = ({ children }) => {
  const isAuth = localStorage.getItem('_a') ? true : false;

  if (!isAuth) {
    console.log('Не авторизован!');
    return <Navigate to="/" />;
  } else {
    return children;
  }
};

export default RequireAuth;
