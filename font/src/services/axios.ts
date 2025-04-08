import axios from 'axios';


const request = axios.create({
  baseURL: import.meta.env.MODE === 'development' ?"http://localhost:8080/api/" : "/api/",
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json'
  },
  withCredentials: true
});

// 响应拦截器：先处理成功响应中的业务错误，再处理 HTTP 错误
// request.interceptors.response.use(
//     (response) => {
//         // 检查业务状态码（如 response.data.code）
//         //如果没有code
//         if (response.data.code === undefined) {
//             return response.data;
//         }
//         if (response.data.code === 401) {
//             const fullUrl = window.location.pathname + window.location.search + window.location.hash;
//             sessionStorage.setItem('redirectUrl', fullUrl); // 存储完整 URL
//             window.location.href = '/'; // 使用 window.location.href 进行页面跳转 需提前在组件中初始化 navigate（见下方上下文方案）
//         }
//         // 业务成功，返回响应数据
//         return response.data;
//     },
//     (error) => {
//         // 处理 HTTP 级错误（如网络问题、500 等）
//         if (error.response?.status === 401) {
//             // 若同时有 HTTP 401（可选，根据实际情况）
//             const fullUrl = window.location.pathname + window.location.search + window.location.hash;
//             sessionStorage.setItem('redirectUrl', fullUrl); // 存储完整 URL
//             window.location.href = '/'; // 使用 window.location.href 进行页面跳转 需提前在组件中初始化 navigate（见下方上下文方案）
//         }
//         return Promise.reject(error);
//     }
// );
export default request;