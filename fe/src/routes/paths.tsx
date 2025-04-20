import { customGeneratePath } from "../utils/helper";

export const loginPath = "/login";
export const registerPath = "/register";

export const login = customGeneratePath(`/${loginPath}`);
export const register = customGeneratePath(`/${registerPath}`);
