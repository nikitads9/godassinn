import React from 'react';
import { Route, Routes } from 'react-router-dom';
import { routes } from '../model/routes.jsx';
import RequireAuth from './RequireAuth.jsx';

const AppRouter = () => {
  return (
    <Routes>
      {routes.map((route) => {
        return (
          <Route
            key={route.path}
            path={route.path}
            // element={route.element}
            element={
              route.authOnly ? (
                <RequireAuth routeRole={route.role}>
                  {route.element}
                </RequireAuth>
              ) : (
                route.element
              )
            }
          />
        );
      })}
    </Routes>
  );
};

export default AppRouter;
