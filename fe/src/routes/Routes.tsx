import { Route, Routes } from "react-router-dom";
import * as paths from "./paths";
import { AuthLayout } from "../components/layout/auth-layout/AuthLayout";

export const RoutesManagement = () => {
  return (
    <Routes>
      <Route element={<AuthLayout />}>
        <Route path={paths.login()} element={<></>} />
        <Route path={paths.register()} element={<></>} />
      </Route>
    </Routes>
  );
};
