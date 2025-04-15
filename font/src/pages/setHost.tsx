import { useContext, useEffect, useState } from "react";
import { Input, Button, Space } from 'antd';
import { HOST_KEY } from "../types/const.ts";
import { useNavigate } from "react-router-dom";
import MessageContext from "../context/message.tsx";
import type { MessageContextValue } from "../context/message.tsx";

export  default function SetHost(){
  const navigate = useNavigate()
  const [hostValue, setHostValue] = useState("");

  const messageContext = useContext(MessageContext);
  const { middleMessageApi } = messageContext as MessageContextValue;


  const onSubmit = () => {
    // 处理确定按钮点击事件，这里可以添加具体的逻辑
      localStorage.setItem(HOST_KEY, hostValue)
    // 这里使用window的跳转而不是navigate，是为了让react刷新整个页面，request可以重新加载
      middleMessageApi.success({content:"设置成功", duration: 2}).then(()=>
        window.location.href = "/set_host"
      )
  };

  const onBack = () => {
    navigate("/");
  };

  function onReset() {
    localStorage.removeItem(HOST_KEY)

    // 这里使用window的跳转而不是navigate，是为了让react刷新整个页面，request可以重新加载
    middleMessageApi.success({content:"设置成功", duration: 1}).then(()=>
      window.location.href = "/set_host"
    )
  }

  useEffect(()=>{
    let curHost = localStorage.getItem(HOST_KEY);
    if (curHost == null || curHost == "") {
      const { protocol, host } = window.location;
      curHost= `${protocol}//${host}`;
    }
    setHostValue(curHost)
  },[])

  useEffect(()=> console.log(hostValue), [hostValue])


  return (
    <div style={{ maxWidth: 600, margin: "15vh auto", padding: 24 }}>
        <div style={{ textAlign: "left", marginBottom: 24,  margin: "5vh auto", fontSize: "30px", color: "#1890ff", fontWeight: "bold" }}>

          这个界面更适合桌面和手机端使用，动态的修改请求的host地址
          <br /> <br />

          浏览器里修改成不是默认的值会有跨域问题
          <br /> <br />
          一定要加上协议和端口号
          <br /> <br />
          Example: http://192.168.1.123:8080

        </div>
      <div style={{ width: '100%' ,fontSize: '30px' ,display: "flex" ,marginBottom : 40 }} >
      <label  htmlFor={"host"} style={{marginRight : 30}} >
        Host:
      </label>
        <Input
          id={"host"}
          value={hostValue}
          onChange={(e) => setHostValue(e.target.value)}
          style={{ fontSize: '30px' }}
        />
      </div>

          <Space>
            <Button onClick={onSubmit} style={{ fontSize: '18px' , margin: 40 }}>
              设置
            </Button>
            <Button onClick={onReset} style={{ fontSize: '18px', margin: 40 }}>
                    重置
            </Button>
            <Button onClick={onBack} style={{ fontSize: '18px', margin: 40 }}>返回</Button>
          </Space>
    </div>
  );
}
