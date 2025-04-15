import { Button } from "antd";
import getRequestAndSetNavigateLocation from "../../services/axios.ts";
import { useLocation, useNavigate } from "react-router-dom";
import {  BaseResponseWithoutData } from "../../types/response.ts";
import { useContext } from "react";
import MessageContext, { type MessageContextValue } from "../../context/message.tsx";



export default function CreateProject() {
  const navigate = useNavigate();
  let request = getRequestAndSetNavigateLocation(navigate, useLocation())

  const messageContext = useContext(MessageContext);
  const { middleMessageApi } = messageContext as MessageContextValue;

  function resetRateLimit() {
    request.post<BaseResponseWithoutData>("/admin/reset_rate_limit").then((res) => {
      if (res.data.code == 0) {
        middleMessageApi.success(res.data.message).then();
      } else {
        middleMessageApi.warning(res.data.message).then();
      }
    })
  }
  return(
    <>
      <div style={{ maxWidth: 600, margin: "10vh auto", padding: 24 }}>
        <div style={{ textAlign: "left", marginBottom: 24,  margin: "5vh auto", fontSize: "30px", color: "#1890ff", fontWeight: "bold" }}>
            项目有严格的速率限制，可能会导致请求失败
           <br /> <br />
            由于误操作导致登录失败，需要等待1小时，如果有人急需使用，可以重置速率限制
        </div>

        <div style={{ textAlign: "center" }}>
        <Button
          onClick={resetRateLimit}
          style={{
            fontSize: 40,
            color: "white", // 文字设为白色，确保可见
            fontWeight: "bold",
            padding: "20px 40px",
            border: "2px solid #FF4444", // 红色边框，增强轮廓
            borderRadius: "12px",
            backgroundColor: "#FF4444", // 鲜艳的红色背景（比文字色深）
            boxShadow: "0 4px 8px rgba(255, 68, 68, 0.2)", // 添加阴影，提升层次感
            transition: "all 0.3s ease", // 过渡效果
          }}
          onMouseEnter={(e) => {
            e.currentTarget.style.backgroundColor = "#FF2D2D"; // 悬停时更深的红色
            e.currentTarget.style.boxShadow = "0 6px 12px rgba(255, 45, 45, 0.3)";
          }}
          onMouseLeave={(e) => {
            e.currentTarget.style.backgroundColor = "#FF4444";
            e.currentTarget.style.boxShadow = "0 4px 8px rgba(255, 68, 68, 0.2)";
          }}
        >
          Reset
        </Button>
        </div>
      </div>
    </>
  )
}
