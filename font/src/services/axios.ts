import axios from "axios";
import type { Location } from "react-router-dom";
import { REDIRECT_KEY } from "../types/const.ts";
import { serverUrl } from "./utils.ts";

type NavigateFunction = (path: string) => void;

// 声明一个 navigate 变量
let navigate: NavigateFunction | null = null;

// 定义一个设置 navigate 函数的方法
const setNavigate = (nav: NavigateFunction) => {
  navigate = nav;
};
let fullPath: string = "/";

const request = axios.create({
  baseURL: `${serverUrl()}api/`,
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

request.interceptors.response.use(
  (response) => {
    // 检查业务状态码（如 response.data.code）
    //如果没有code
    if (response.data.code === undefined) {
      return response;
    }
    if (response.data.code === 401) {
      sessionStorage.setItem(REDIRECT_KEY, fullPath); // 存储完整 URL
      if (navigate) {
        navigate("/"); // 使用 navigate 进行页面跳转
      }
    }
    // 业务成功，返回响应数据
    return response;
  },
  (error) => {
    // 处理 HTTP 级错误（如网络问题、500 等）
    const status = error.response?.status;
    const message = `server err: ${error.message}`;
    // 返回正常响应
    return {
      data: {
        code: status,
        message: message,
      },
    };
  },
);
export default function getRequestAndSetNavigateLocation(
  nav: NavigateFunction,
  location: Location<any>,
) {
  setNavigate(nav);
  fullPath = setFullPath(location);
  return request;
}
function setFullPath(lct: Location<any>): string {
  return lct.pathname + lct.search + lct.hash;
}
