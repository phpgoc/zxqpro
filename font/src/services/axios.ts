import axios from "axios";

type NavigateFunction = (path: string) => void;

// 声明一个 navigate 变量
let navigate: NavigateFunction | null = null;

// 定义一个设置 navigate 函数的方法
const setNavigate = (nav: NavigateFunction) => {
  navigate = nav;
};

const request = axios.create({
  baseURL: `${import.meta.env.VITE_SERVER_URL}api/`,
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
      const fullUrl =
        window.location.pathname +
        window.location.search +
        window.location.hash;
      sessionStorage.setItem("redirectUrl", fullUrl); // 存储完整 URL
      if (navigate) {
        navigate("/"); // 使用 navigate 进行页面跳转
      }
    }
    // 业务成功，返回响应数据
    return response;
  },
  (error) => {
    // 处理 HTTP 级错误（如网络问题、500 等）
    if (error.response?.status === 401) {
      // 若同时有 HTTP 401（可选，根据实际情况）
      const fullUrl =
        window.location.pathname +
        window.location.search +
        window.location.hash;
      sessionStorage.setItem("redirectUrl", fullUrl); // 存储完整 URL
      if (navigate) {
        navigate("/"); // 使用 navigate 进行页面跳转
      }
    }
    return Promise.reject(error);
  },
);
export default function getRequestAndSetNavigate(nav: NavigateFunction) {
  setNavigate(nav);
  return request;
}
