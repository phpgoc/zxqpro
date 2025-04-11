import { useContext, useEffect, useState } from "react";
import getRequestAndSetNavigate from "../../services/axios.ts";
import { useLocation, useNavigate } from "react-router-dom";
import { Input, Button, Modal, Row, Col } from "antd";
import MessageContext, {
  type MessageContextValue,
} from "../../context/message.tsx";
import { BaseResponse, UserInfo } from "../../types/response.ts";
import {useUserContext} from "../../context/userInfo.tsx";
import * as React from "react";


function ImageSelectorPopup ({setAvatar}:{ setAvatar:(n:number) => void}){
  const [isModalVisible, setIsModalVisible] = useState(false);

  const showModal = () => {
    setIsModalVisible(true);
  };

  const handleOk = () => {
    setIsModalVisible(false);
  };

  const handleCancel = () => {
    setIsModalVisible(false);
  };

  return (
    <div>
      <Button type="primary" onClick={showModal}>
        打开图片选择器
      </Button>
      <Modal
        title="选择头像图片"
        open={isModalVisible}
        onOk={handleOk}
        onCancel={handleCancel}
        width="70%"
        height="70vh"
        footer={null}
      >
        <Row gutter={[16, 16]}>
          {Array.from({ length: 20 }, (_, index) => (
            <Col
              key={index + 1}
              xs={24} // 超小屏幕（< 576px），每行 1 个
              sm={12} // 小屏幕（≥ 576px），每行 2 个
              md={8}  // 中等屏幕（≥ 768px），每行 3 个
              lg={6}  // 大屏幕（≥ 992px），每行 4 个
            >
              <img
                src={import.meta.env.VITE_SERVER_URL+`static/avatar/${index + 1}.webp`}
                alt={`Avatar ${index + 1}`}
                style={{ width: '100%', cursor: 'pointer' }}
                onClick={() => setAvatar(index + 1)}
              />
            </Col>
          ))}
        </Row>
      </Modal>
    </div>
  )
}

export default function UpdateUser() {
  const {user, updateUser} = useUserContext()

  const [name, setName] = useState(user.name);
  const [userName, setUserName] = useState(user.user_name);
  const [email, setEmail] = useState(user.email);
  const [avatar, setAvatar] = useState(user.avatar);


  let avatarImgUrl = import.meta.env.VITE_SERVER_URL+`static/avatar/${avatar}.webp`;

  const navigate = useNavigate();
  let request = getRequestAndSetNavigate(navigate, useLocation());

  const messageContext = useContext(MessageContext);
  const { middleApi } = messageContext as MessageContextValue;





  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    request
      .post<BaseResponse>("/user/update", {  user_name: userName, email, avatar})
      .then((response) => {
        if (response.data.code == 0) {
          let newUser = {
            ...user,
            user_name: userName,
            email: email,
            avatar: avatar,
          }
          updateUser(newUser )
          middleApi.success(response.data.message).then((_: any) => {
            middleApi.success(response.data.message).then()
          });
        } else {
          middleApi.warning(response.data.message).then();
        }
      })
      .catch((_) => {
        middleApi.error("Registration failed. Please try again.").then();
      });
  };
  useEffect(()=>{
    request.get<BaseResponse<UserInfo>>("user/info").then((res) => {
      if (res.data.code === 0) {
        setName(res.data.data.name);
        setUserName(res.data.data.user_name);
        setEmail(res.data.data.email);
        setAvatar(res.data.data.avatar)
        localStorage.setItem("userInfo", JSON.stringify(res.data.data))
      }
    })
  },[])

  const divStyle = {
    display: "flex",
    alignItems: "flex-start",
    width: "50%",
    marginBottom: 20,
  };
  const labelStyle = {
    width: "30%",
    // textAlign: "right",
    align : "right",
    marginRight: 10,
    marginTop: 5,
  }

  function handleImgSelect() {

  }

  return (
    <form onSubmit={handleSubmit}>
      <div
        style={{
          textAlign: "center",
          marginTop: 100,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <div style={divStyle}>
          <label style={labelStyle}>
            Name
          </label>
        <Input disabled={true}
          value={name}
               style={{ width: "70%" }}
        />
        </div>
        <div style={divStyle}>
          <label style={labelStyle}>
            userName
          </label>
        <Input
          value={userName}
          onChange={(e) => setUserName(e.target.value)}
          placeholder="userName"
          style={{ width: "70%" }}
        />
        </div>
        <div style={divStyle}>
          <label style={labelStyle}>
            email
          </label>
        <Input
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Email"
          style={{ width: "70%" }}
        />
        </div>
        <div style={divStyle}>
          <label style={labelStyle}>
            avatar
          </label>
          <img src={avatarImgUrl}  alt={userName} onClick={handleImgSelect}/>
          <ImageSelectorPopup setAvatar={setAvatar}/>
        </div>
        <Button
          type="primary"
          htmlType="submit"
          style={{ width: "15%", margin: "20px 5%" }}
        >
          Update
        </Button>
      </div>
    </form>
  );
}
