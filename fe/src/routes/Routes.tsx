import { Route, Routes } from "react-router-dom";
import * as paths from "./paths";
import { AuthLayout } from "../components/layout/auth-layout/AuthLayout";
import { PageLogin } from "../features/login/page-login/PageLogin";

export const RoutesManagement = () => {
  return (
    <Routes>
      <Route element={<AuthLayout />}>
        <Route path={paths.login()} element={<PageLogin />} />
        <Route path={paths.register()} element={<></>} />
      </Route>
    </Routes>
  );
};
