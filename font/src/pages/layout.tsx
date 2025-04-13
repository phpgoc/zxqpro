import { Button, Layout, Menu } from "antd";
import {
  HomeOutlined,
  SettingOutlined,
  SettingFilled,
  UsergroupAddOutlined,
} from "@ant-design/icons";
import { Outlet, useLocation, useNavigate } from "react-router-dom";
import { BaseResponse, BaseResponseWithoutData } from "../types/response.ts";
import getRequestAndSetNavigate from "../services/axios.ts";
import {useUserContext} from "../context/userInfo.tsx";
import { avatarUrl, isAdmin } from "../services/utils.ts";
import UserListSelect from "../components/userList.tsx";
import { useContext, useState } from "react";
import MessageContext, { type MessageContextValue } from "../context/message.tsx";

const { Header, Content } = Layout;

//根据 user.id是不是1来判断是不是需要admin权限
const items = [
  {
    key: "project",
    label: "Project",
    icon: <HomeOutlined style={{ fontSize: "5vh", lineHeight: "6vh" }} />,
  },
  {
    key: "task",
    label: "Task",
    icon: (
      <UsergroupAddOutlined style={{ fontSize: "5vh", lineHeight: "6vh" }} />
    ),
  },
  {
    key: "admin",
    label: "Admin",
    icon: <SettingFilled style={{ fontSize: "5vh", lineHeight: "6vh" }} />,
  },
  {
    key: "setting",
    label: "Setting",
    icon: <SettingOutlined style={{ fontSize: "5vh", lineHeight: "6vh" }} />,
  },
];

export default function ZxqLayout() {
  const [isSelectVisible, setIsSelectVisible] = useState(false);

  const navigate = useNavigate();
  let request = getRequestAndSetNavigate(navigate, useLocation());
  const currentPath = useLocation().pathname;
  const messageContext = useContext(MessageContext);
  const { middleMessageApi } = messageContext as MessageContextValue;
  const {user} = useUserContext()
  if (!isAdmin(user.id)) {
    delete items[2];
  }
  const [sharedUserId, setSharedUserId] = useState(0);
  const avatarSrc =  avatarUrl(user.avatar)

  function logout() {
    request.post<BaseResponse>("user/logout").then((res) => {
      if (res.data.code == 0) {
        navigate("/");
      }
    });
  }

  function handleButtonClick() {
    setIsSelectVisible(true)
    setTimeout(() => {
      setIsSelectVisible(false)
    }, 10000)
  }

  function handleSelectChange(newUserId : number) {
    setSharedUserId(newUserId)
    request.post<BaseResponseWithoutData>("message/share_link", {
      to_user_id: newUserId,
      link: window.location.href,
    }).then((res) => {
      if (res.data.code == 0) {
        middleMessageApi.success("分享成功").then();
      } else {
        middleMessageApi.error("分享失败").then();
      }
    });
    setIsSelectVisible(false)
  }


  return (
    <Layout>
      <Header
        style={{
          // 导航栏背景：渐变蓝色（从 #1890FF 到 #40a9ff）
          background: "linear-gradient(90deg, #1890FF 0%, #40a9ff 100%)",
          // 增加阴影提升层次感
          boxShadow: "0 2px 4px rgba(0, 0, 0, 0.08)",
          // height: '56px', // 适当增加高度让导航更舒展
          height: "7vh",
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          padding: "0 1vw", // 左右内边距，让内容不贴边
        }}
      >
        <div
          style={{
            height: "100%",
            width: "80%", // 左边元素占据50%宽度
            background: "linear-gradient(90deg, #1890FF 0%, #40a9ff 100%)",
          }}
        >
          <Menu
            style={{ height: "100%" }}
            mode="horizontal"
            items={items}
            selectedKeys={
              currentPath.includes("/project")
                ? ["project"]
                : currentPath.includes("/task")
                  ? ["task"]
                  : currentPath.includes("/admin")
                    ? ["admin"]
                    : currentPath.includes("/setting")
                      ? ["setting"]
                      : []
            }
            onClick={(e) => {
              if (e.key === "project") navigate("/project");
              else if (e.key === "task") navigate("/task");
              else if (e.key === "admin") navigate("/admin");
              else if (e.key === "setting") navigate("/setting");
            }}
          />
        </div>
        <div
          style={{
            display: "flex",
            alignItems: "center",
            height: "100%",
          }}
        >
          <div style={{ position: 'relative' }}>
              <Button
                style={{
                  position: 'relative',
                  zIndex: 1,
                }}
                onClick={handleButtonClick}
              >
                Share Link
              </Button>

            {isSelectVisible && (
              <div
                style={{
                  position: 'absolute',
                  top: 0,
                  left: 0,
                  zIndex: 2
                }}
              >
                <UserListSelect
                  userId={sharedUserId}
                  onChange={handleSelectChange}
                  includeAdmin={true}
                  filterSelf={true}
                />
              </div>
            )}
          </div>

          <img
            src={avatarSrc}
            alt="User Avatar"
            width={40}
            height={40}
            style={{ borderRadius: 10 }}
          />
          <span style={{ marginLeft: 10 }}>{user.user_name}</span>

          <Button
            type="primary"
            size="middle"
            style={{
              backgroundColor: "#fff", // 按钮背景为白色
              color: "#1890FF", // 按钮文字为蓝色主色
              border: "none",
              fontWeight: 600,
              padding: "8px 16px",
              borderRadius: 24, // 圆角更柔和
            }}
            onClick={logout}
          >
            Logout
          </Button>

        </div>
      </Header>
      <Content style={{ padding: "20px 24px 0" }}>
        {" "}
        {/* 顶部 20px 内边距 */}
        <Outlet />
      </Content>
      {/* 其他页面内容 */}
    </Layout>
  );
}
